package model

type User struct {
	Id          int    `json:"-" db:"id"`
	Name        string `json:"name" binding:"required"`
	Surname     string `json:"surname" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Avatar      string `json:"avatar" binding:"required"`
	Description string `json:"description" binding:"required"`
	Login       string `json:"login" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Locked      bool   `json:"locked"`
	IsTrainer   bool   `json:"is_trainer"`
	IsAdmin     bool   `json:"is_admin"`
	IsModer     bool   `json:"is_moder"`
}

/*type User struct {
	Id          int    `json:"-" db:"id"`
	Name        string `json:"name" binding:"required"`
	Surname     string `json:"surname" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Avatar      string `json:"avatar" binding:"required"`
	Description string `json:"description" binding:"required"`
	Login       string `json:"login" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Locked      bool   `json:"locked" binding:"required"`
	IsTrainer   bool   `json:"is_trainer" binding:"required"`
	IsAdmin     bool   `json:"is_admin" binding:"required"`
	IsModer     bool   `json:"is_moder" binding:"required"`
}
*/
