package models

type Comment struct {
	Id      int64  `gorm:"column:id;autoIncrement"`
	Comment string `gorm:"column:comment"`
	IdPost  string `gorm:"column:id_post"`
	IdUser  string `gorm:"column:id_user"`
	User    User   `gorm:"foreignKey:id_user;references:id"`
}

type CommentRequest struct {
	Comment string `json:"comment"`
	IdPost  string `json:"id_post"`
	IdUser  string `json:"id_user"`
}

type CommentResUpd struct {
	Id      int64  `json:"id"`
	Comment string `json:"comment"`
}
