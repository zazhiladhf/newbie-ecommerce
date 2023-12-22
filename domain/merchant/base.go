package merchant

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/zazhiladhf/newbie-ecommerce/domain/auth"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/middleware"
)

func RegisterRoutesMerchant(router fiber.Router, db *sqlx.DB) {
	aRepo := auth.NewPostgreSqlxRepository(db)
	mRepo := NewPostgreSqlxRepository(db)
	svc := NewService(mRepo, aRepo)
	handler := NewHandler(svc)

	merchantRouter := router.Group("/v1/merchants")
	{
		merchantRouter.Post("/profile", middleware.AuthMiddleware(), handler.CreateProfile)
		merchantRouter.Get("/profile", middleware.AuthMiddleware(), handler.GetProfile)
		merchantRouter.Put("/profile", middleware.AuthMiddleware(), handler.UpdateProfile)
	}
}
