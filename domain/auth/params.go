package auth

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type registerRequest struct {
	Email    string `json:"email" binding:"email"`
	Password string `json:"password"`
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type response struct {
	HttpCode  int         `json:"-"`
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Error     string      `json:"error,omitempty"`
	ErrorCode string      `json:"error_code,omitempty"`
	Payload   interface{} `json:"payload,omitempty"`
}

type Payload struct {
	AccessToken string `json:"access_token,omitempty"`
	Role        string `json:"role,omitempty"`
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

func ResponseSuccess(c *fiber.Ctx, success bool, message string, httpCode int, payload interface{}) error {
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
	case err == ErrRepository:
		return ApiResponse(c, http.StatusInternalServerError, false, "error repository", err.Error(), "50001", nil)
	default:
		return ApiResponse(c, http.StatusInternalServerError, false, "internal server error", err.Error(), "99999", nil)
	}
}
