package main

import (
	"context"
	fitter "fitter-srv"
	"fitter-srv/pkg/handler"
	"fitter-srv/pkg/repository"
	"fitter-srv/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Задаем формат JSON для логгера
	logrus.SetFormatter(new(logrus.JSONFormatter))
	// Инициализируем конфиг файл
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	// Загружаем значение пароля для переменной окружения
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	// Инициализируем базу данных
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	// Объявлем зависимсти в нужном порядке
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// Инициализируем и запускаем сервер
	srv := new(fitter.Server)

	// Для реализации Graceful shutdown
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occurred while running http server: %s", err.Error())
		}
	}()

	logrus.Print("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit // чтение из канала блокирует выполнение главной горутины main

	// Остнаовка сервера
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on seerver shutting down: %s", err.Error())
	}
	// Закрытие соединений с БД
	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close", err.Error())
	}
}

// Функция для инициализации конфигурационных файлов
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
