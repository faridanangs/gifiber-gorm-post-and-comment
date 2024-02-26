package apps

import (
	"post_comment_api/internals/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitializeRouter(app *fiber.App, db *gorm.DB) {
	InitializeRouterUser(app, db)
	InitializeRouterPost(app, db)
	InitializeRouterComment(app, db)
}

func InitializeRouterUser(app *fiber.App, db *gorm.DB) {
	user := services.NewUserService(db)
	api := app.Group("/api/user")
	api.Get("/", func(c *fiber.Ctx) error {
		return user.GetAllUser(c)
	})
	api.Get("/:id", func(c *fiber.Ctx) error {
		return user.GetUser(c)
	})
	api.Post("/create", func(c *fiber.Ctx) error {
		return user.CreateUser(c)
	})
	api.Delete("/:id", func(c *fiber.Ctx) error {
		return user.DeleteUser(c)
	})
	api.Put("/:id", func(c *fiber.Ctx) error {
		return user.UpdateUser(c)
	})
}
func InitializeRouterPost(app *fiber.App, db *gorm.DB) {
	user := services.NewPostService(db)
	api := app.Group("/api/post")
	api.Get("/", func(c *fiber.Ctx) error {
		return user.GetAllPost(c)
	})
	api.Get("/:id", func(c *fiber.Ctx) error {
		return user.GetPost(c)
	})
	api.Post("/create", func(c *fiber.Ctx) error {
		return user.CreatePost(c)
	})
	api.Delete("/:id", func(c *fiber.Ctx) error {
		return user.DeletePost(c)
	})
	api.Put("/:id", func(c *fiber.Ctx) error {
		return user.UpdatePost(c)
	})
}
func InitializeRouterComment(app *fiber.App, db *gorm.DB) {
	comment := services.NewCommentService(db)
	api := app.Group("/api/comment")
	api.Post("/create", func(c *fiber.Ctx) error {
		return comment.CreateComment(c)
	})
	api.Put("/:id", func(c *fiber.Ctx) error {
		return comment.UpdateComment(c)
	})
	api.Delete("/:id", func(c *fiber.Ctx) error {
		return comment.DeleteComment(c)
	})
	api.Get("/", func(c *fiber.Ctx) error {
		return comment.GetAllComment(c)
	})
}
