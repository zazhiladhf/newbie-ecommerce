package category

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutesCategory(router fiber.Router, db *sqlx.DB) {
	repo := NewPostgreSqlxRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	categoryRouter := router.Group("/v1/categories")
	{
		categoryRouter.Get("/", handler.GetListCategories)
	}
}
