package product

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/meilisearch/meilisearch-go"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/middleware"
)

func RegisterRoutesProduct(router fiber.Router, dbSqlx *sqlx.DB, client *meilisearch.Client) {
	meiliRepo := NewMeiliRepository(client)
	// authRepo := auth.NewPostgreSqlxRepository(dbSqlx)
	// redis := auth.NewRedisRepository(&redis.Client{})
	// authService := auth.NewService(authRepo, redis)
	repo := NewPostgresSQLXRepository(dbSqlx)
	svc := NewService(repo, meiliRepo)
	handler := NewHandler(svc)

	productRouter := router.Group("/v1/products")
	{
		productRouter.Post("/", middleware.AuthMiddleware(), handler.CreateProduct)
		// productRouter.Get("/", middleware.AuthMiddleware(), handler.GetProducts)
		productRouter.Get("/", middleware.AuthMiddleware(), handler.GetProductsByEmail)
	}

}
