package merchant

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/helper"
)

type MerchantHandler struct {
	svc MerchantService
}

func NewHandler(svc MerchantService) MerchantHandler {
	return MerchantHandler{
		svc: svc,
	}
}

func (h MerchantHandler) CreateProfile(c *fiber.Ctx) (err error) {
	var req RequestBodyCreateMerchant
	email := c.Locals("email").(string)
	id := c.Locals("id").(string)
	idInt, _ := strconv.Atoi(id)

	err = c.BodyParser(&req)
	if err != nil {
		log.Println("error when try to parsing body request with error", err)
		return helper.ResponseError(c, err)
	}

	merchant, err := NewMerchant().newFromRequest(req)
	if err != nil {
		log.Println("error when try to validate request body with error", err)
		return helper.ResponseError(c, err)
	}
	merchant.AuthId = idInt

	_, err = h.svc.CreateProfileMerchant(c.UserContext(), merchant, email)
	if err != nil {
		log.Println("error when try to create profile merchant with error", err)
		pqErr, ok := err.(*pq.Error)
		if ok {
			switch pqErr.Code {
			case "23505":
				return helper.ResponseError(c, helper.ErrDuplicateAuthId)
			default:
				return helper.ResponseError(c, helper.ErrRepository)
			}
		} else {
			log.Println("unknown error with error:", helper.ErrInternalServer)
		}

		return helper.ResponseError(c, err)
	}

	return helper.ResponseSuccess(c, true, "create merchant success", http.StatusCreated, nil, nil)
}

func (h MerchantHandler) GetProfile(c *fiber.Ctx) (err error) {
	id := c.Locals("id").(string)
	idInt, _ := strconv.Atoi(id)

	merchant, err := h.svc.GetMerchantById(c.UserContext(), idInt)
	if err != nil {
		log.Println("error when try to get merchant by token (id) with error", err)
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

	return helper.ResponseSuccess(c, true, "get merchant success", http.StatusOK, merchant, nil)
}

func (h MerchantHandler) UpdateProfile(c *fiber.Ctx) (err error) {
	var req RequestBodyCreateMerchant
	email := c.Locals("email").(string)
	id := c.Locals("id").(string)
	idInt, _ := strconv.Atoi(id)

	err = c.BodyParser(&req)
	if err != nil {
		log.Println("error when try to parsing body request with error", err)
		return helper.ResponseError(c, err)
	}

	merchant, err := NewMerchant().newFromRequest(req)
	if err != nil {
		log.Println("error when try to validate request body with error", err)
		return helper.ResponseError(c, err)
	}
	merchant.AuthId = idInt

	err = h.svc.UpdateProfileMerchant(c.UserContext(), merchant, email)
	if err != nil {
		log.Println("error when try to update profile merchant with error", err)
		pqErr, ok := err.(*pq.Error)
		if ok {
			switch pqErr.Code {
			default:
				return helper.ResponseError(c, helper.ErrRepository)
			}
		} else {
			log.Println("unknown error with error:", err)
		}
		return helper.ResponseError(c, err)
	}

	return helper.ResponseSuccess(c, true, "update merchant success", http.StatusOK, nil, nil)
}
