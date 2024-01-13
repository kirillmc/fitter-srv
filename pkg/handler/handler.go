package handler

import (
	"fitter-srv/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	// DI
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

// gin - фреймворк для разработки рест апи
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		programs := api.Group("/programs")
		{
			programs.POST("/", h.createProgram)
			programs.GET("/", h.getAllPrograms)
			programs.GET("/:id", h.getProgramByID)
			programs.PUT("/:id", h.updateProgram)
			programs.DELETE("/:id", h.deleteProgram)

			days := programs.Group(":id/days")
			{
				days.POST("/", h.createDay)
				days.GET("/", h.getAllDays)
				days.GET("/:idD", h.getDayByID)
				days.PUT("/:idD", h.updateDay)
				days.DELETE("/:idD", h.deleteDay)

				exercises := days.Group(":idD/exercises")
				{
					exercises.POST("/", h.createExercise)
					exercises.GET("/", h.getAllExercises)
					exercises.GET("/:idE", h.getExerciseByID)
					exercises.PUT("/:idE", h.updateExercise)
					exercises.DELETE("/:idE", h.deleteExercise)

					sets := exercises.Group(":idE/sets")
					{
						sets.POST("/", h.createSet)
						sets.GET("/", h.getAllSets)
						sets.GET("/:idS", h.getSetByID)
						sets.PUT("/:idS", h.updateSet)
						sets.DELETE("/:idS", h.deleteSet)

						statistics := sets.Group(":idS/statistics")
						{
							statistics.POST("/", h.createStatistic)
							statistics.GET("/", h.getAllStatistics)
							statistics.GET("/:idSs", h.getStatisticByID)
							statistics.PUT("/:idSs", h.updateStatistic)
							statistics.DELETE("/:idSs", h.deleteStatistic)
						}
					}
				}
			}
		}

	}
	//days := api.Group("days")
	//{
	//
	//}
	/*wasds
		items := api.Group(":id/items")
		{
			items.POST("/", h.createItem)
			items.GET("/", h.getAllItems)
		}
	}
	items := api.Group("items")
	{
		items.GET("/:id", h.getItemById)
		items.PUT("/:id", h.updateItem)
		items.DELETE("/:id", h.deleteItem)

	}
	*/
	return router
}
