package models

type Post struct {
	Id        string    `gorm:"column:id"`
	Title     string    `gorm:"column:title"`
	Deskripsi string    `gorm:"column:deskripsi"`
	IdUser    string    `gorm:"column:id_user"`
	Comments  []Comment `gorm:"foreignKey:id_post;reference:id"`
}

type PostRequest struct {
	Title     string `json:"title"`
	Deskripsi string `json:"deskripsi"`
	IdUser    string `json:"id_user"`
}
type PostUpdate struct {
	Title     string `json:"title"`
	Deskripsi string `json:"deskripsi"`
}
type PostResponse struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Deskripsi string `json:"deskripsi"`
}
