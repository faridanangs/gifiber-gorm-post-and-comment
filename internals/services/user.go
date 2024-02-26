package services

import (
	"log"
	"post_comment_api/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		DB: db,
	}
}

func (service *UserService) CreateUser(c *fiber.Ctx) error {
	user := models.UserRequest{}
	if err := c.BodyParser(&user); err != nil {
		log.Fatal("Error parsing user: ", err)
	}

	generatePassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		log.Fatal("Error generating password: ", err)
	}

	userModel := models.User{
		Id:       uuid.NewString(),
		Username: user.Username,
		Password: string(generatePassword),
	}
	if err := service.DB.Debug().Create(&userModel).Error; err != nil {
		log.Fatal("Error creating user: ", err)
	}

	return c.JSON(models.Response{
		Code:   200,
		Status: "OK",
		Data:   models.UserResponse{Id: userModel.Id, Username: userModel.Username},
	})
}
func (service *UserService) UpdateUser(c *fiber.Ctx) error {
	user := models.UserUpdate{}
	if err := c.BodyParser(&user); err != nil {
		log.Fatal("Error parsing user: ", err)
	}

	userModel := models.User{}
	if err := service.DB.Debug().Model(models.User{}).Where("id = ?", c.Params("id")).Take(&userModel).Error; err != nil {
		log.Fatal("Error get user: ", err)
	}
	userModel.Username = user.Username
	if err := service.DB.Debug().Save(&userModel).Error; err != nil {
		log.Fatal("Error updating user: ", err)
	}
	return c.JSON(models.Response{
		Code:   200,
		Status: "OK",
		Data:   models.UserResponse{Id: userModel.Id, Username: userModel.Username},
	})

}
func (service *UserService) GetUser(c *fiber.Ctx) error {
	user := models.User{}
	if err := service.DB.Debug().Model(models.User{}).Where("id = ?", c.Params("id")).Take(&user).Error; err != nil {
		log.Fatal("Error getting user: ", err)
	}

	return c.JSON(fiber.Map{
		"status":   "ok",
		"messaege": "Get user successfully",
		"data":     user,
	})
}
func (service *UserService) DeleteUser(c *fiber.Ctx) error {
	user := models.User{}
	if err := service.DB.Debug().Model(models.User{}).Where("id = ?", c.Params("id")).Delete(&user).Error; err != nil {
		log.Fatal("Error delete user: ", err)
	}

	return c.JSON(fiber.Map{
		"status":   "ok",
		"messaege": "Delete user successfully",
	})
}

func (service *UserService) GetAllUser(c *fiber.Ctx) error {
	user := []models.User{}
	if err := service.DB.Debug().Preload("Posts").Preload("Posts.Comments").Preload("Posts.Comments.User").Model(models.User{}).Find(&user).Error; err != nil {
		log.Fatal("Error getting all users: ", err)
	}

	return c.JSON(fiber.Map{
		"status":   "ok",
		"messaege": "Get all user successfully",
		"data":     user,
	})
}
