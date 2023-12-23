package order

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/helper"
)

type Handler struct {
	service service
}

func NewHandler(service service) Handler {
	return Handler{
		service: service,
	}
}

func (h Handler) Checkout(c *fiber.Ctx) (err error) {
	var req CheckoutRequest

	err = c.BodyParser(&req)
	if err != nil {
		log.Println("error when try to parsing body request with error", err)
		return helper.ResponseError(c, err)
	}

	id := c.Locals("id").(string)
	idInt, _ := strconv.Atoi(id)

	req.AuthId = idInt

	resp, err := h.service.Checkout(c.UserContext(), req)
	if err != nil {
		log.Println("error when try to parsing body request with error", err)
		return helper.ResponseError(c, err)
	}

	return helper.ResponseSuccess(c, true, "success", http.StatusOK, resp, nil)
}

// func (h Handler) OrderHistories(c *fiber.Ctx) (err error)
// func (h Handler) ListOrders(c *fiber.Ctx) (err error)
// func (h Handler) Webhook(c *fiber.Ctx) (err error)