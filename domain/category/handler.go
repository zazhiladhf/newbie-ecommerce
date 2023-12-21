package category

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/helper"
)

type categoryHandler struct {
	svc CategoryService
}

func newHandler(svc CategoryService) categoryHandler {
	return categoryHandler{
		svc: svc,
	}
}

func (h categoryHandler) GetListCategories(c *fiber.Ctx) error {
	listCategories, err := h.svc.getListCategories(c.UserContext())
	if err != nil {
		log.Println("error when try to get list categories with error", err)
		pqErr, ok := err.(*pq.Error)
		if ok {
			switch pqErr.Code {
			// case "23505":
			// 	return helper.ResponseError(c, helper.ErrDuplicateEmail)
			default:
				return helper.ResponseError(c, helper.ErrRepository)
			}
		} else {
			log.Println("unknown error with error:", ErrInternalServer)
		}

		return helper.ResponseError(c, err)
	}
	return helper.ResponseSuccess(c, true, "get categories success", http.StatusOK, listCategories, nil)
}
