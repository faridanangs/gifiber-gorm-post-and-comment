package models

type User struct {
	Id       string `gorm:"column:id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Posts    []Post `gorm:"foreignKey:id_user;references:id"`
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserUpdate struct {
	Username string `json:"username"`
}

type UserResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}
