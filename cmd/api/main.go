package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/zazhiladhf/newbie-ecommerce/config"
	"github.com/zazhiladhf/newbie-ecommerce/domain/auth"
	"github.com/zazhiladhf/newbie-ecommerce/domain/category"
	"github.com/zazhiladhf/newbie-ecommerce/domain/files"
	"github.com/zazhiladhf/newbie-ecommerce/domain/merchant"
	"github.com/zazhiladhf/newbie-ecommerce/domain/product"
	"github.com/zazhiladhf/newbie-ecommerce/domain/search"
	"github.com/zazhiladhf/newbie-ecommerce/domain/user"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/database"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/images"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/meili"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/middleware"
)

// type CloudSvc interface {
// 	Upload(ctx context.Context, file interface{}, pathDestination string, quality string) (uri string, err error)
// }

// // setup services
// type service struct {
// 	// disini kita akan menggunakan kontrak ke cloud providernya
// 	cloud CloudSvc
// }

// var path = "public/uploads"
// var svc = service{}

// func init() {
// 	err := config.LoadConfig("./config/config.yaml")
// 	if err != nil {
// 		log.Println("error when try to LoadConfig with error :", err.Error())
// 	}

// 	cloudName := config.Cfg.Cloudinary.Name
// 	apiKey := config.Cfg.Cloudinary.ApiKey
// 	apiSecret := config.Cfg.Cloudinary.ApiSecret

// 	cloudClient, err := images.NewCloudinary(cloudName, apiKey, apiSecret)
// 	if err != nil {
// 		panic(err)
// 	}

// 	svc = service{
// 		cloud: cloudClient,
// 	}
// }

func main() {
	// setup config
	err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		log.Println("error when try to LoadConfig with error :", err.Error())
	}

	// setup fiber
	router := fiber.New(fiber.Config{
		AppName: config.Cfg.App.Name,
		// Prefork: true,
	})

	middleware.FiberMiddleware(router)

	// setup database PostgreSQL
	dbSqlx, err := database.ConnectPostgresSqlx(config.Cfg.DB)
	if err != nil {
		log.Println("error connect postgre with error :", err.Error())
		// panic(err)
	}

	// setup redis
	dbRedis, err := database.ConnectRedis(context.Background(), config.Cfg.Redis)
	if err != nil {
		log.Println("error connect redis", err)
	}

	// setup meilisearch
	client, err := meili.ConnectMeilisearch(config.Cfg.Meili.Host, config.Cfg.Meili.ApiKey)
	if err != nil {
		log.Println("error connect meili", err)
	}

	// setup cloudinary
	cloudName := config.Cfg.Cloudinary.Name
	apiKey := config.Cfg.Cloudinary.ApiKey
	apiSecret := config.Cfg.Cloudinary.ApiSecret

	cloudClient, err := images.NewCloudinary(cloudName, apiKey, apiSecret)
	if err != nil {
		panic(err)
	}

	// migration db
	log.Println("running db migration")
	err = database.Migrate(dbSqlx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("migration done")

	// regoster routes
	auth.RegisterRoutesAuth(router, dbSqlx, dbRedis)
	category.RegisterRoutesCategory(router, dbSqlx)
	product.RegisterRoutesProduct(router, dbSqlx, client)
	files.RegisterRoutesFile(router, cloudClient, cloudName, apiKey, apiSecret)
	user.RegisterRoutesUser(router, dbSqlx)
	merchant.RegisterRoutesMerchant(router, dbSqlx)
	search.RegisterRoutesMeili(router, client)

	// listen app
	router.Listen(config.Cfg.App.Port)
}
