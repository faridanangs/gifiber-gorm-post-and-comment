package services

import (
	"log"
	"post_comment_api/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostService struct {
	DB *gorm.DB
}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{
		DB: db,
	}
}

func (service *PostService) CreatePost(c *fiber.Ctx) error {
	post := models.PostRequest{}
	if err := c.BodyParser(&post); err != nil {
		log.Fatal("Error parsing Post: ", err)
	}

	PostModel := models.Post{
		Id:        uuid.NewString(),
		Title:     post.Title,
		Deskripsi: post.Deskripsi,
		IdUser:    post.IdUser,
	}
	if err := service.DB.Debug().Create(&PostModel).Error; err != nil {
		log.Fatal("Error creating Post: ", err)
	}

	return c.JSON(models.Response{
		Code:   200,
		Status: "OK",
		Data: models.PostResponse{
			Id:        PostModel.Id,
			Title:     post.Title,
			Deskripsi: post.Deskripsi,
		},
	})
}
func (service *PostService) UpdatePost(c *fiber.Ctx) error {
	post := models.PostUpdate{}
	if err := c.BodyParser(&post); err != nil {
		log.Fatal("Error parsing Post: ", err)
	}

	postModel := models.Post{}
	if err := service.DB.Debug().Model(models.Post{}).Where("id = ?", c.Params("id")).Take(&postModel).Error; err != nil {
		log.Fatal("Error get Post: ", err)
	}

	postModel.Title = post.Title
	postModel.Deskripsi = post.Deskripsi

	if err := service.DB.Debug().Save(&postModel).Error; err != nil {
		log.Fatal("Error updating Post: ", err)
	}
	return c.JSON(models.Response{
		Code:   200,
		Status: "OK",
		Data: models.PostResponse{
			Id:        postModel.Id,
			Title:     postModel.Title,
			Deskripsi: postModel.Deskripsi,
		},
	})

}
func (service *PostService) GetPost(c *fiber.Ctx) error {
	post := models.Post{}
	if err := service.DB.Debug().Model(models.Post{}).Where("id = ?", c.Params("id")).Take(&post).Error; err != nil {
		log.Fatal("Error getting Post: ", err)
	}

	return c.JSON(fiber.Map{
		"status":   "ok",
		"messaege": "Get Post successfully",
		"data":     post,
	})
}
func (service *PostService) DeletePost(c *fiber.Ctx) error {
	post := models.Post{}
	if err := service.DB.Debug().Model(models.Post{}).Where("id = ?", c.Params("id")).Delete(&post).Error; err != nil {
		log.Fatal("Error delete post: ", err)
	}

	return c.JSON(fiber.Map{
		"status":   "ok",
		"messaege": "Delete post successfully",
	})
}

func (service *PostService) GetAllPost(c *fiber.Ctx) error {
	post := []models.Post{}
	if err := service.DB.Debug().Preload("Comments").Model(models.Post{}).Find(&post).Error; err != nil {
		log.Fatal("Error getting all posts: ", err)
	}

	return c.JSON(fiber.Map{
		"status":   "ok",
		"messaege": "Get all post successfully",
		"data":     post,
	})
}
