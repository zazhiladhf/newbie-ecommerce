package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/zazhiladhf/newbie-ecommerce/domain/auth"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/middleware"
)

func RegisterRoutesUser(router fiber.Router, db *sqlx.DB) {
	authRepo := auth.NewPostgreSqlxRepository(db)
	repo := NewPostgreSqlxRepository(db)
	// redis := NewRedisRepository(client)
	svc := NewService(repo, authRepo)
	handler := NewHandler(svc)

	userRouter := router.Group("/v1/users")
	{
		userRouter.Post("/profile", middleware.AuthMiddleware(), handler.CreateProfile)
		userRouter.Get("/profile", middleware.AuthMiddleware(), handler.GetProfile)
		userRouter.Put("/profile", middleware.AuthMiddleware(), handler.UpdateProfile)
	}
}
