package services

import (
	"log"
	"post_comment_api/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CommentService struct {
	DB *gorm.DB
}

func NewCommentService(db *gorm.DB) *CommentService {
	return &CommentService{
		DB: db,
	}
}

func (service *CommentService) CreateComment(c *fiber.Ctx) error {
	comment := models.CommentRequest{}
	if err := c.BodyParser(&comment); err != nil {
		log.Fatal("Error parsing comment: ", err)
	}

	commentModel := models.Comment{
		Comment: comment.Comment,
		IdPost:  comment.IdPost,
		IdUser:  comment.IdUser,
	}
	if err := service.DB.Debug().Create(&commentModel).Error; err != nil {
		log.Fatal("Error creating comment: ", err)
	}

	return c.JSON(models.Response{
		Code:   200,
		Status: "OK",
		Data: models.CommentResUpd{
			Comment: commentModel.Comment,
			Id:      commentModel.Id,
		},
	})
}

func (service *CommentService) UpdateComment(c *fiber.Ctx) error {
	comment := models.CommentResUpd{}
	if err := c.BodyParser(&comment); err != nil {
		log.Fatal("Error parsing Comment: ", err)
	}

	CommentModel := models.Comment{}
	if err := service.DB.Debug().Where("id = ?", c.Params(("id"))).Take(&CommentModel).Error; err != nil {
		log.Fatal("Error get Comment: ", err)
	}

	CommentModel.Comment = comment.Comment

	if err := service.DB.Debug().Save(&CommentModel).Error; err != nil {
		log.Fatal("Error updating Comment: ", err)
	}
	return c.JSON(models.Response{
		Code:   200,
		Status: "OK",
		Data:   models.CommentResUpd{Id: CommentModel.Id, Comment: CommentModel.Comment},
	})

}
func (service *CommentService) DeleteComment(c *fiber.Ctx) error {
	comment := models.Comment{}
	if err := service.DB.Debug().Model(models.Comment{}).Where("id = ?", c.Params("id")).Delete(&comment).Error; err != nil {
		log.Fatal("Error delete Comment: ", err)
	}

	return c.JSON(fiber.Map{
		"status":   "ok",
		"messaege": "Delete Comment successfully",
	})
}

func (service *CommentService) GetAllComment(c *fiber.Ctx) error {
	comment := []models.Comment{}
	if err := service.DB.Debug().Preload("User").Model(models.Comment{}).Find(&comment).Error; err != nil {
		log.Fatal("Error getting all comments: ", err)
	}

	return c.JSON(fiber.Map{
		"status":   "ok",
		"messaege": "Get all comment successfully",
		"data":     comment,
	})
}
