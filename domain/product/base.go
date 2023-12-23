package product

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/meilisearch/meilisearch-go"
	"github.com/zazhiladhf/newbie-ecommerce/domain/auth"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/middleware"
)

func RegisterRoutesProduct(router fiber.Router, dbSqlx *sqlx.DB, client *meilisearch.Client) {
	meiliRepo := NewMeiliRepository(client)
	// authRepo := auth.NewPostgreSqlxRepository(dbSqlx)
	// redis := auth.NewRedisRepository(&redis.Client{})
	productRepo := NewPostgresSQLXRepository(dbSqlx)
	authRepo := auth.NewPostgreSqlxRepository(dbSqlx)

	svc := NewService(productRepo, authRepo, meiliRepo)
	handler := NewHandler(svc)

	productRouter := router.Group("/v1/products")
	{
		productRouter.Post("/", middleware.AuthMiddleware(), handler.CreateProduct)
		productRouter.Get("/", middleware.AuthMiddleware(), handler.GetListProducts)
		productRouter.Get("/id/:product_id", middleware.AuthMiddleware(), handler.GetDetailProduct)
		productRouter.Put("/id/:product_id", middleware.AuthMiddleware(), handler.UpdateProduct)
		productRouter.Get("/detail/:sku", middleware.AuthMiddleware(), handler.GetDetailProductUserPerspective)
	}

}
