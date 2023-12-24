package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/zazhiladhf/newbie-ecommerce/config"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/helper"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/jwt"
)

// // JWTProtected func for specify routes group with JWT authentication.
// // See: https://github.com/gofiber/jwt
// func JWTProtected() func(*fiber.Ctx) error {
// 	// Create config for JWT authentication middleware.
// 	config := jwtMiddleware.Config{
// 		SigningKey:   []byte("secret"),
// 		ContextKey:   "jwt", // used in private routes
// 		ErrorHandler: jwtError,
// 	}

// 	return jwtMiddleware.New(config)
// }

// func jwtError(c *fiber.Ctx, err error) error {
// 	// Return status 401 and failed authentication error.
// 	if err.Error() == "Missing or malformed JWT" {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   err.Error(),
// 		})
// 	}

// 	// Return status 401 and failed authentication error.
// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 		"error": true,
// 		"msg":   err.Error(),
// 	})
// }

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		// tokenHeader := c.Get("Authorization")
		// if tokenHeader == "" {
		// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		// 		"error": true,
		// 		"msg":   err.Error(),
		// 	})
		// }

		claims, err := jwt.ExtractTokenMetadata(c)
		if err != nil {
			log.Println("error when try extract token metadata with error", err)
			return helper.ResponseError(c, helper.ErrUnauthorized)

		}

		// log.Println(&claims)

		id := claims.Id
		email := claims.Email
		role := claims.Role

		c.Locals("id", id)
		c.Locals("email", email)
		c.Locals("role", role)

		return c.Next()
	}
}

func XenditTokenValidation() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		token := c.Get("X-Callback-Token")

		if token != config.Cfg.Payment.WebhookToken {
			log.Println("error when when try to webhook error", err)
			return helper.ResponseError(c, helper.ErrUnauthorized)
		}

		return c.Next()
	}
}
