package auth

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/middleware"
)

func RegisterRoutesAuth(router fiber.Router, db *sqlx.DB, client *redis.Client) {
	repo := NewPostgreSqlxRepository(db)
	redis := NewRedisRepository(client)
	svc := NewService(repo, redis)
	handler := NewHandler(svc)

	authRouter := router.Group("/v1/auth")
	{
		authRouter.Post("/register", handler.Register)
		authRouter.Post("/login", handler.Login)
		authRouter.Patch("/role", middleware.AuthMiddleware(), handler.UpdateRole)
	}

	router.Get("", func(c *fiber.Ctx) error {
		return c.SendString("I'm a GET request!")
	})
}
