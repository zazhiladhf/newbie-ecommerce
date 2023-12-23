package product

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/helper"
)

type ProductHandler struct {
	svc Service
}

func NewHandler(svc Service) ProductHandler {
	return ProductHandler{
		svc: svc,
	}
}

func (h ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var req = CreateProductRequest{}
	email := c.Locals("email").(string)
	id := c.Locals("id").(string)
	idInt, _ := strconv.Atoi(id)

	err := c.BodyParser(&req)
	if err != nil {
		log.Println("error when try to parsing body request with error", err)
		return helper.ResponseError(c, err)
	}

	product, err := NewProduct().newFromRequest(req)
	if err != nil {
		log.Println("error when try to validate request body with error", err)
		return helper.ResponseError(c, err)
	}
	product.MerchantId = idInt

	err = h.svc.CreateProductByMerchant(c.UserContext(), product, email)
	if err != nil {
		log.Println("error when try to create product by token merchant with error", err)
		pqErr, ok := err.(*pq.Error)
		if ok {
			switch pqErr.Code {
			default:
				return helper.ResponseError(c, helper.ErrRepository)
			}
		} else {
			log.Println("unknown error with error:", helper.ErrInternalServer)
		}

		return helper.ResponseError(c, err)
	}

	return helper.ResponseSuccess(c, true, "create product success", http.StatusCreated, nil, nil)
}

func (h ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	resp, err := h.svc.GetAllProducts(c.UserContext())
	if err != nil {
		log.Println("error when try to get all product with error", err)
		return helper.ResponseError(c, err)
	}

	return helper.ResponseSuccess(c, true, "get products success", http.StatusOK, resp, nil)
}

func (h ProductHandler) GetListProducts(c *fiber.Ctx) error {
	queryParam := c.Query("query")
	email := c.Locals("email").(string)

	limit := c.Query("limit", "10")
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		log.Println("error when try convert limit to int with error", err)
		return helper.ResponseError(c, err)
	}

	page := c.Query("page", "1")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		log.Println("error when try to convert page to int with error", err)
		return helper.ResponseError(c, err)
	}

	resp, totalData, err := h.svc.GetListProductsMerchant(c.UserContext(), queryParam, email, limitInt, pageInt)
	if err != nil {
		log.Println("error when try to get products by token with error", err)
		pqErr, ok := err.(*pq.Error)
		if ok {
			switch pqErr.Code {
			default:
				return helper.ResponseError(c, helper.ErrRepository)
			}
		} else {
			log.Println("unknown error with error:", helper.ErrInternalServer)
		}
		return helper.ResponseError(c, err)
	}
	log.Println("total data:", totalData)
	log.Println("list products:", resp)

	if totalData != 0 {
		pagination := helper.NewPaginationResponse(queryParam, limitInt, pageInt, totalData)
		return helper.ResponseSuccess(c, true, "get products success", http.StatusOK, resp, pagination)
	}

	return helper.ResponseSuccess(c, true, "get products success", http.StatusOK, resp, nil)

}

func (h ProductHandler) GetDetailProduct(c *fiber.Ctx) error {
	email := c.Locals("email").(string)
	// id := c.Locals("id").(string)
	productId := c.Params("product_id")

	productIdInt, err := strconv.Atoi(productId)
	if err != nil {
		log.Println("error when try convert product_id to int with error", err)
		return helper.ResponseError(c, err)
	}

	response, err := h.svc.GetDetailProductById(c.UserContext(), productIdInt, email)
	if err != nil {
		log.Println("error when try to get detail product with error", err)
		return helper.ResponseError(c, err)
	}

	return helper.ResponseSuccess(c, true, "get product success", http.StatusOK, response, nil)
}

func (h ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	var req CreateProductRequest
	email := c.Locals("email").(string)
	productId := c.Params("product_id")

	productIdInt, err := strconv.Atoi(productId)
	if err != nil {
		log.Println("error when try convert product_id to int with error", err)
		return helper.ResponseError(c, err)
	}

	err = c.BodyParser(&req)
	if err != nil {
		log.Println("error when try to parsing body request with error", err)
		return err
	}

	product, err := NewProduct().newFromRequest(req)
	if err != nil {
		log.Println("error when try to validate request body with error", err)
		return helper.ResponseError(c, err)
	}
	product.Id = productIdInt

	err = h.svc.UpdateProduct(c.UserContext(), product, email)
	if err != nil {
		log.Println("error when try to update product with error", err)
		pqErr, ok := err.(*pq.Error)
		if ok {
			switch pqErr.Code {
			default:
				return helper.ResponseError(c, helper.ErrRepository)
			}
		} else {
			log.Println("unknown error with error:", helper.ErrInternalServer)
		}

		return helper.ResponseError(c, err)
	}

	return helper.ResponseSuccess(c, true, "update product success", http.StatusOK, nil, nil)
}

func (h ProductHandler) GetDetailProductUserPerspective(c *fiber.Ctx) error {
	sku := c.Params("sku")

	resp, err := h.svc.GetDetailProductUserPerspective(c.UserContext(), sku)
	if err != nil {
		log.Println("error when try to get detail product user perspective with error", err)
		return helper.ResponseError(c, err)
	}

	return helper.ResponseSuccess(c, true, "get products success", http.StatusOK, resp, nil)
}
