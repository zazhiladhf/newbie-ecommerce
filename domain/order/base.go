package order

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/zazhiladhf/newbie-ecommerce/domain/product"
	"github.com/zazhiladhf/newbie-ecommerce/domain/user"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/middleware"
	paymentgateway "github.com/zazhiladhf/newbie-ecommerce/pkg/payment-gateway"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRouteOrder(router fiber.Router, sqlx *sqlx.DB, mongo *mongo.Database, xenditClient paymentgateway.Xendit) {
	productRepo := product.NewPostgresSQLXRepository(sqlx)
	userRepo := user.NewPostgreSqlxRepository(sqlx)
	orderRepo := NewMongoRepository(mongo)
	paymentRepo := newPaymentGatewayRepository(xenditClient)
	service := NewService(productRepo, userRepo, orderRepo, paymentRepo)
	handler := NewHandler(service)

	orderRoute := router.Group("/v1/orders")
	{
		orderRoute.Post("", middleware.AuthMiddleware(), handler.Checkout)
	}
}
