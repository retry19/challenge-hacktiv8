package mygram

import (
	"github.com/gofiber/fiber/v2"
	"github.com/retry19/challenge-hacktiv8/12-final-project/api/my-gram/middleware"
	"github.com/retry19/challenge-hacktiv8/12-final-project/internal/auth"
	"github.com/retry19/challenge-hacktiv8/12-final-project/internal/comment"
	"github.com/retry19/challenge-hacktiv8/12-final-project/internal/photo"
	"github.com/retry19/challenge-hacktiv8/12-final-project/internal/socialmedia"
	"gorm.io/gorm"
)

func buildV1Routes(app *fiber.App, db *gorm.DB) {
	authService := auth.NewAuthService(db)
	authHandler := auth.NewAuthHandler(authService)
	app.Post("/register", authHandler.Register)
	app.Post("/login", authHandler.Login)

	app.Use("/*", middleware.NewAuthJwt())

	photoService := photo.NewPhotoService(db)
	photoRouter := app.Group("/photos")
	{
		photoHander := photo.NewPhotoHandler(photoService, authService)
		photoRouter.Get("/", photoHander.GetAll)
		photoRouter.Post("/", photoHander.CreatePhoto)
		photoRouter.Get("/:id", photoHander.GetOne)
		photoRouter.Delete("/:id", photoHander.DeletePhoto)
		photoRouter.Put("/:id", photoHander.UpdatePhoto)
	}

	socialMediaRouter := app.Group("/social-media")
	{
		socialMediaHander := socialmedia.NewSocialMediaHandler(socialmedia.NewSocialMediaService(db), authService)
		socialMediaRouter.Get("/", socialMediaHander.GetAll)
		socialMediaRouter.Post("/", socialMediaHander.CreateSocialMedia)
		socialMediaRouter.Get("/:id", socialMediaHander.GetOne)
		socialMediaRouter.Delete("/:id", socialMediaHander.DeleteSocialMedia)
		socialMediaRouter.Put("/:id", socialMediaHander.UpdateSocialMedia)
	}

	commentRouter := app.Group("/comments")
	{
		commentHander := comment.NewCommentHandler(comment.NewCommentService(db), authService, photoService)
		commentRouter.Get("/", commentHander.GetAll)
		commentRouter.Post("/", commentHander.CreateComment)
		commentRouter.Get("/:id", commentHander.GetOne)
		commentRouter.Delete("/:id", commentHander.DeleteComment)
		commentRouter.Put("/:id", commentHander.UpdateComment)
	}
}
