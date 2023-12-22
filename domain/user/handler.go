package user

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/helper"
)

type UserHandler struct {
	svc UserService
}

func NewHandler(svc UserService) UserHandler {
	return UserHandler{
		svc: svc,
	}
}

func (h UserHandler) CreateProfile(c *fiber.Ctx) (err error) {
	var req RequestBodyCreateProfileUser
	email := c.Locals("email").(string)
	id := c.Locals("id").(string)
	idInt, _ := strconv.Atoi(id)

	err = c.BodyParser(&req)
	if err != nil {
		log.Println("error when try to parsing body request with error", err)
		return helper.ResponseError(c, err)
	}

	user, err := NewUser().newFromRequest(req)
	if err != nil {
		log.Println("error when try to validate form register with error", err)
		return helper.ResponseError(c, err)
	}
	user.AuthId = idInt

	_, err = h.svc.CreateProfileUser(c.UserContext(), user, email)
	if err != nil {
		log.Println("error when try to create profile with error", err)
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

	return helper.ResponseSuccess(c, true, "registration success", http.StatusCreated, nil, nil)
}

// func (h UserHandler) Login(c *fiber.Ctx) error {
// 	var req loginRequest

// 	err := c.BodyParser(&req)
// 	if err != nil {
// 		log.Println("error when try to parsing body request with error", err)
// 		return helper.ResponseError(c, err)
// 	}

// 	model, err := NewAuth().ValidateFormLogin(req)
// 	if err != nil {
// 		log.Println("error when try to validate form login with error", err)
// 		return helper.ResponseError(c, err)
// 	}

// 	itemAuth, token, err := h.svc.Login(c.UserContext(), model)
// 	if err != nil {
// 		log.Println("error when try to login with error", err)
// 		if err == ErrInvalidEmail {
// 			return helper.ResponseError(c, helper.ErrInvalidEmail)
// 		}
// 		pqErr, ok := err.(*pq.Error)
// 		if ok {
// 			switch pqErr.Code {
// 			case "23505":
// 				return helper.ResponseError(c, helper.ErrDuplicateEmail)
// 			default:
// 				return helper.ResponseError(c, helper.ErrRepository)
// 			}
// 		} else {
// 			log.Println("unknown error with error:", err)
// 		}
// 		return helper.ResponseError(c, err)
// 	}

// 	return helper.ResponseSuccess(c, true, "login success", http.StatusOK, Payload{
// 		AccessToken: token,
// 		Role:        itemAuth.Role,
// 	}, nil)
// }
