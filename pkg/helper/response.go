package helper

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrEmailEmpty          = errors.New("email required")
	ErrInvalidEmail        = errors.New("invalid email")
	ErrPasswordEmpty       = errors.New("password required")
	ErrInvalidPassword     = errors.New("invalid password")
	ErrDuplicateEmail      = errors.New("email already used")
	ErrUserAlreadyMerchant = errors.New("user already as a merchant")

	ErrCategoriesNotFound = errors.New("categories not found")

	ErrUnauthorized = errors.New("unauthorized")

	ErrEmptyName       = errors.New("name is required")
	ErrEmptyImageURL   = errors.New("image_url is required")
	ErrEmptyStock      = errors.New("stock is required")
	ErrEmptyPrice      = errors.New("price is required")
	ErrEmptyCategoryId = errors.New("category_id is required")
	ErrNotFound        = errors.New("product not found")

	ErrGenderEmpty        = errors.New("gender is required")
	ErrInvalidGender      = errors.New("gender is invalid")
	ErrEmptyPhoneNumber   = errors.New("phone_number is empty")
	ErrInvalidPhoneNumber = errors.New("phone_number length must be greater than equal 10")
	ErrEmptyNameUser      = errors.New("name is required")
	ErrEmptyAddress       = errors.New("address is required")
	ErrEmptyDateOfBirth   = errors.New("date_of_birth is required")
	ErrInvalidDateOfBirth = errors.New("date_of_birth is invalid")
	ErrEmptyImageURLUser  = errors.New("image_url is required")
	ErrUnauthorizedUser   = errors.New("please provide jwt token")
	ErrInvalidRole        = errors.New("invalid role")
	ErrDuplicateAuthId    = errors.New("user can't create more profile")
	ErrUserNotFound       = errors.New("user not found")

	ErrRepository     = errors.New("error repository")
	ErrInternalServer = errors.New("unknown error")
)

type response struct {
	HttpCode   int         `json:"-"`
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Error      string      `json:"error,omitempty"`
	ErrorCode  string      `json:"error_code,omitempty"`
	Payload    interface{} `json:"payload,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
}

type Payload struct {
	AccessToken string `json:"access_token,omitempty"`
	Role        string `json:"role,omitempty"`
	Url         string `json:"url,omitempty"`
}

func ApiResponse(c *fiber.Ctx, httpCode int, success bool, message string, err string, errorCode string, payload interface{}) error {
	c = c.Status(httpCode)
	isSuccess := httpCode >= 200 && httpCode < 300

	if isSuccess {
		return c.JSON(response{
			Success: true,
			Message: message,
			Payload: payload,
		})
	}

	return c.JSON(response{
		Success:   false,
		Message:   message,
		Error:     err,
		ErrorCode: errorCode,
	})
}

func ResponseSuccess(c *fiber.Ctx, success bool, message string, httpCode int, payload interface{}, pagination interface{}) error {
	resp := response{
		Success:   success,
		Message:   message,
		Error:     "",
		ErrorCode: "",
		Payload:   payload,
	}
	c = c.Status(httpCode)
	return c.JSON(resp)
}

func ResponseError(c *fiber.Ctx, err error) error {
	switch {
	case err == ErrEmailEmpty:
		return ApiResponse(c, http.StatusBadRequest, false, "bad request", err.Error(), "40001", nil)
	case err == ErrInvalidEmail:
		return ApiResponse(c, http.StatusBadRequest, false, "bad request", err.Error(), "40002", nil)
	case err == ErrPasswordEmpty:
		return ApiResponse(c, http.StatusBadRequest, false, "bad request", err.Error(), "40003", nil)
	case err == ErrInvalidPassword:
		return ApiResponse(c, http.StatusBadRequest, false, "bad request", err.Error(), "40004", nil)
	case err == ErrDuplicateEmail:
		return ApiResponse(c, http.StatusConflict, false, "duplicate entry", err.Error(), "40901", nil)
	case err == ErrUserAlreadyMerchant:
		return ApiResponse(c, http.StatusBadRequest, false, "bad request", err.Error(), "40001", nil)

	case err == ErrCategoriesNotFound:
		return ApiResponse(c, http.StatusNotFound, false, "category not found", err.Error(), "40401", nil)

	case err == ErrUnauthorized:
		return ApiResponse(c, http.StatusUnauthorized, false, "unauthorized", err.Error(), "40101", nil)

	case err == ErrEmptyName:
		return ApiResponse(c, http.StatusBadRequest, false, "bad request", err.Error(), "40401", nil)
	case err == ErrEmptyImageURL:
		return ApiResponse(c, http.StatusBadRequest, false, "bad request", err.Error(), "40402", nil)
	case err == ErrEmptyStock:
		return ApiResponse(c, http.StatusBadRequest, false, "bad request", err.Error(), "40403", nil)
	case err == ErrEmptyPrice:
		return ApiResponse(c, http.StatusBadRequest, false, "bad request", err.Error(), "40404", nil)
	case err == ErrEmptyCategoryId:
		return ApiResponse(c, http.StatusBadRequest, false, "bad request", err.Error(), "40405", nil)
	case err == ErrNotFound:
		return ApiResponse(c, http.StatusNotFound, false, "category not found", err.Error(), "40401", nil)

	case err == ErrDuplicateAuthId:
		return ApiResponse(c, http.StatusConflict, false, "duplicate entry", err.Error(), "40901", nil)

	case err == ErrGenderEmpty:
		return ApiResponse(c, http.StatusBadRequest, false, "bad request", err.Error(), "40001", nil)
	case err == ErrInvalidGender:
		return ApiResponse(c, http.StatusBadRequest, false, "bad request", err.Error(), "40002", nil)
	case err == ErrEmptyPhoneNumber:
		return ApiResponse(c, http.StatusBadRequest, false, "bad request", err.Error(), "40003", nil)
	case err == ErrInvalidPhoneNumber:
		return ApiResponse(c, http.StatusBadRequest, false, "bad request", err.Error(), "40004", nil)
	case err == ErrEmptyNameUser:
		return ApiResponse(c, http.StatusBadRequest, false, "bad request", err.Error(), "40005", nil)
	case err == ErrEmptyAddress:
		return ApiResponse(c, http.StatusBadRequest, false, "bad request", err.Error(), "40006", nil)
	case err == ErrEmptyDateOfBirth:
		return ApiResponse(c, http.StatusBadRequest, false, "bad request", err.Error(), "40007", nil)
	case err == ErrInvalidDateOfBirth:
		return ApiResponse(c, http.StatusBadRequest, false, "bad request", err.Error(), "40008", nil)
	case err == ErrEmptyImageURLUser:
		return ApiResponse(c, http.StatusBadRequest, false, "bad request", err.Error(), "40009", nil)
	case err == ErrUnauthorizedUser:
		return ApiResponse(c, http.StatusUnauthorized, false, "unauthorized", err.Error(), "40101", nil)
	case err == ErrInvalidRole:
		return ApiResponse(c, http.StatusUnauthorized, false, "unauthorized", err.Error(), "40102", nil)
	case err == ErrDuplicateAuthId:
		return ApiResponse(c, http.StatusConflict, false, "duplicate", err.Error(), "40901", nil)
	case err == ErrUserNotFound:
		return ApiResponse(c, http.StatusNotFound, false, "not found", err.Error(), "40000", nil)

	case err == ErrRepository:
		return ApiResponse(c, http.StatusInternalServerError, false, "error repository", err.Error(), "50001", nil)
	default:
		return ApiResponse(c, http.StatusInternalServerError, false, "internal server error", err.Error(), "99999", nil)
	}
}

type Pagination struct {
	Query string `json:"query"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Total int    `json:"total"`
}

func NewPaginationResponse(queryParam string, limit, page, totalData int) Pagination {
	return Pagination{
		Query: queryParam,
		// Limit:     limit,
		// Page:      page,
		Total: CountTotalPage(totalData, limit),
	}
}

func CountTotalPage(total, limit int) int {
	if limit == 0 {
		return 0
	}
	return (total + limit - 1) / limit
}
