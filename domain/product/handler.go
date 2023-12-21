package product

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/helper"
)

type productHandler struct {
	svc Service
}

func NewHandler(svc Service) productHandler {
	return productHandler{
		svc: svc,
	}
}

func (h productHandler) CreateProduct(c *fiber.Ctx) error {
	var req = CreateProductRequest{}
	email := c.Locals("email").(string)

	err := c.BodyParser(&req)
	if err != nil {
		log.Println("error when try to parsing body request with error", err)
		return helper.ResponseError(c, err)
	}

	product, err := NewProduct().newFromRequest(req)
	if err != nil {
		log.Println("error when try to validate from register with error", err)
		return helper.ResponseError(c, err)
	}
	product.AuthEmail = email

	err = h.svc.createProduct(c.UserContext(), product)
	if err != nil {
		log.Println("error when try to create product by token with error", err)
		pqErr, ok := err.(*pq.Error)
		if ok {
			switch pqErr.Code {
			// case "23505":
			// 	return helper.ResponseError(c, ErrDuplicateEmail)
			default:
				return helper.ResponseError(c, helper.ErrRepository)
			}
		} else {
			log.Println("unknown error with error:", ErrInternalServer)
		}

		return helper.ResponseError(c, err)
	}

	return helper.ResponseSuccess(c, true, "create product success", http.StatusCreated, nil, nil)

	// if resp.Email != req.Auth.Email {
	// 	log.Println("token tidak valid")
	// 	return err
	// }

	// currentAuth := c.GetReqHeaders().(auth.Auth)
	// req.Auth = currentAuth

	// model := Product{
	// 	Name:  req.Name,
	// 	Price: req.Price,
	// 	Stock: req.Stock,
	// }

	// client, err := database.ConnectRedis(config.Cfg.Redis)
	// if err != nil {
	// 	log.Println("error when to try migration db with error :", err.Error())
	// 	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
	// 		"success": false,
	// 		"message": "ERR BAD REQUEST",
	// 		"error":   err.Error(),
	// 	})
	// }

	// if client == nil {
	// 	log.Println("db not connected with unknown error")
	// 	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
	// 		"success": false,
	// 		"message": "ERR BAD REQUEST",
	// 		"error":   err.Error(),
	// 	})
	// }

	// var itemAuth auth.Auth
	// token := client.Get(c.UserContext(), itemAuth.Email)
	// if err != nil {
	// 	log.Println("error when try to get data to redis with message :", err.Error())
	// 	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
	// 		"success": false,
	// 		"message": "ERR BAD REQUEST",
	// 		"error":   err.Error(),
	// 	})
	// }

	// claim, err := jwt.ValidateToken(token.String())
	// if err != nil {
	// 	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
	// 		"error": err.Error(),
	// 	})
	// }

}

// func (h productHandler) GetProducts(c *fiber.Ctx) error {
// 	// var models []Product
// 	listProducts, err := h.svc.GetProducts(c.UserContext())

// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"success":    false,
// 			"message":    "internal server error",
// 			"error":      err.Error(),
// 			"error_code": "50001",
// 			// "payload": listCategories,
// 		})
// 	}

// 	return c.Status(http.StatusOK).JSON(fiber.Map{
// 		"success": true,
// 		"message": "GET ALL SUCCESS",
// 		"payload": listProducts,
// 	})
// }

func (h productHandler) GetProductsByEmail(c *fiber.Ctx) error {
	queryParam := c.Query("query")
	email := c.Locals("email").(string)

	limit := c.Query("limit", "10")
	limitValue, err := strconv.Atoi(limit)
	if err != nil {
		log.Println("error when try convert limit to int with error", err)
		return helper.ResponseError(c, err)
	}

	page := c.Query("page", "1")
	pageValue, err := strconv.Atoi(page)
	if err != nil {
		log.Println("error when try to convert page to int with error", err)
		return helper.ResponseError(c, err)
	}

	listProducts, totalData, err := h.svc.GetProductsByEmail(c.UserContext(), queryParam, email, limitValue, pageValue)
	if err != nil {
		log.Println("error when try to get products by token with error", err)
		pqErr, ok := err.(*pq.Error)
		if ok {
			switch pqErr.Code {
			// case "HV":
			// 	return helper.ResponseError(c, ErrRepository)
			default:
				return helper.ResponseError(c, helper.ErrRepository)
			}
		} else {
			log.Println("unknown error with error:", ErrInternalServer)
		}
		return helper.ResponseError(c, err)
	}
	log.Println("total data:", totalData)
	log.Println("list products:", listProducts)

	if totalData != 0 {
		pagination := helper.NewPaginationResponse(queryParam, limitValue, pageValue, totalData)
		return helper.ResponseSuccess(c, true, "get products success", http.StatusOK, listProducts, pagination)
	}

	return helper.ResponseSuccess(c, true, "get products success", http.StatusOK, listProducts, nil)

}

// func (h Handler) GetProductById(c *fiber.Ctx) error {
// 	model := Product{}
// 	id, err := strconv.Atoi(c.Params("id"))
// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
// 			"success": false,
// 			"message": "ERR BAD REQUEST",
// 			"error":   "invalid id",
// 		})
// 	}

// 	product, err := h.svc.GetProductById(c.UserContext(), model, id)
// 	if err != nil {
// 		var payload fiber.Map
// 		httpCode := 400

// 		if model.Id != id {
// 			payload = fiber.Map{
// 				"success": false,
// 				"message": "ERR BAD REQUEST",
// 				"error":   err.Error(),
// 			}
// 			httpCode = http.StatusNotFound
// 		} else if err == ErrEmptyName || err == ErrEmptyImageURL || err == ErrEmptyPrice || err == ErrEmptyStock {
// 			payload = fiber.Map{
// 				"success": false,
// 				"message": "ERR BAD REQUEST",
// 				"error":   "invaid id",
// 			}
// 			httpCode = http.StatusBadRequest
// 		} else {
// 			payload = fiber.Map{
// 				"success": false,
// 				"message": "ERR INTERNAL",
// 				"error":   "ada masalah pada server",
// 			}
// 			httpCode = http.StatusInternalServerError
// 		}

// 		return c.Status(httpCode).JSON(payload)
// 	}
// 	return c.Status(http.StatusOK).JSON(fiber.Map{
// 		"success": true,
// 		"message": "GET DATA SUCCESS",
// 		"payload": product,
// 	})

// }

// func (h Handler) UpdateProduct(c *fiber.Ctx) error {
// 	var model Product
// 	var req = CreateProductRequest{}
// 	err := c.BodyParser(&req)

// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
// 			"success": false,
// 			"message": "ERR BAD REQUEST",
// 			"error":   err.Error(),
// 		})
// 	}

// 	id, err := strconv.Atoi(c.Params("id"))
// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
// 			"success": false,
// 			"message": "ERR BAD REQUEST",
// 			"error":   "invalid id",
// 		})
// 	}

// 	product, err := h.svc.GetProductById(c.UserContext(), model, id)
// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
// 			"success": false,
// 			"message": "ERR BAD REQUEST",
// 			"error":   err.Error(),
// 		})
// 	}

// 	product.Name = req.Name
// 	product.Price = req.Price
// 	product.Stock = req.Stock

// 	err = h.svc.UpdateProduct(c.UserContext(), product, id)
// 	if err != nil {
// 		var payload fiber.Map
// 		httpCode := 400

// 		if model.Id != id {
// 			payload = fiber.Map{
// 				"success": false,
// 				"message": "ERR BAD REQUEST",
// 				"error":   err.Error(),
// 			}
// 			httpCode = http.StatusNotFound
// 		} else if err == ErrEmptyName || err == ErrEmptyImageURL || err == ErrEmptyPrice || err == ErrEmptyStock {
// 			payload = fiber.Map{
// 				"success": false,
// 				"message": "ERR BAD REQUEST",
// 				"error":   "invaid id",
// 			}
// 			httpCode = http.StatusBadRequest
// 		} else {
// 			payload = fiber.Map{
// 				"success": false,
// 				"message": "ERR INTERNAL",
// 				"error":   "ada masalah pada server",
// 			}
// 			httpCode = http.StatusInternalServerError
// 		}

// 		return c.Status(httpCode).JSON(payload)
// 	}
// 	return c.Status(http.StatusCreated).JSON(fiber.Map{
// 		"success": true,
// 		"message": "UPDATE SUCCESS",
// 	})
// }

// func (h Handler) DeleteProduct(c *fiber.Ctx) error {
// 	model := Product{}
// 	id, err := strconv.Atoi(c.Params("id"))
// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
// 			"success": false,
// 			"message": "ERR BAD REQUEST",
// 			"error":   "invalid id",
// 		})
// 	}

// 	product, err := h.svc.GetProductById(c.UserContext(), model, id)
// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
// 			"success": false,
// 			"message": "ERR BAD REQUEST",
// 			"error":   err.Error(),
// 		})
// 	}

// 	err = h.svc.DeleteProduct(c.UserContext(), product, id)
// 	if err != nil {
// 		var payload fiber.Map
// 		httpCode := 400

// 		if model.Id != id {
// 			payload = fiber.Map{
// 				"success": false,
// 				"message": "ERR BAD REQUEST",
// 				"error":   err.Error(),
// 			}
// 			httpCode = http.StatusNotFound
// 		} else if err == ErrEmptyName || err == ErrEmptyImageURL || err == ErrEmptyPrice || err == ErrEmptyStock {
// 			payload = fiber.Map{
// 				"success": false,
// 				"message": "ERR BAD REQUEST",
// 				"error":   "invaid id",
// 			}
// 			httpCode = http.StatusBadRequest
// 		} else {
// 			payload = fiber.Map{
// 				"success": false,
// 				"message": "ERR INTERNAL",
// 				"error":   "ada masalah pada server",
// 			}
// 			httpCode = http.StatusInternalServerError
// 		}

// 		return c.Status(httpCode).JSON(payload)
// 	}
// 	return c.Status(http.StatusOK).JSON(fiber.Map{
// 		"success": true,
// 		"message": "DELETE DATA SUCCESS",
// 	})

// }
