package search

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meilisearch/meilisearch-go"
)

func RegisterRoutesMeili(router fiber.Router, client *meilisearch.Client) {
	// meiliRepo := NewMeiliRepository(client)
	// authRepo := auth.NewPostgreSqlxRepository(dbSqlx)
	// redis := auth.NewRedisRepository(&redis.Client{})
	// authRepo := auth.NewPostgreSqlxRepository(dbSqlx)

	svc := NewMeiliService(client)
	handler := NewMeiliHandler(svc)

	searchRouter := router.Group("/v1/search/products")
	{
		searchRouter.Get("/", handler.SearchProduct)
		// searchRouter.Get("/sync", handler.SyncProduct)
	}

}
