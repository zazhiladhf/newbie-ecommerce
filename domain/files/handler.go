package files

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/helper"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/images"
)

type CloudSvc interface {
	Upload(ctx context.Context, file interface{}, pathDestination string, quality string) (uri string, err error)
}

type cloudHandler struct {
	svc images.Cloudinary
}

func NewHandler(svc images.Cloudinary) cloudHandler {
	return cloudHandler{
		svc: svc,
	}
}

func (h cloudHandler) Upload(c *fiber.Ctx) error {
	typeFile := c.FormValue("type", "")
	quality := c.FormValue("quality", "10 ")
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("error when try to parse FormFile with detail :", err.Error())
		return helper.ResponseError(c, err)
	}

	if file.Size > 1*1024*1024 {
		errBadrequest := fiber.ErrBadRequest
		errBadrequest.Message = "file to big. expected 1MB"
		log.Println("error with detail :", errBadrequest.Error(), "file size :", (file.Size / 1024 / 1024), "MB")
		return helper.ResponseError(c, errBadrequest)
	}

	if typeFile != "products" {
		if typeFile != "users" {
			if typeFile != "merchants" {
				return helper.ErrInvalidEmail
			}
		}

	}

	source, err := file.Open()
	if err != nil {
		errBadrequest := fiber.ErrBadRequest
		errBadrequest.Message = err.Error()
		log.Println("error when try to open file with detail :", err.Error())
		return helper.ResponseError(c, errBadrequest)
	}

	defer source.Close()

	// siapin buffer kosong
	buffer := bytes.NewBuffer(nil)

	// lalu copy file ke object buffer
	_, err = io.Copy(buffer, source)
	if err != nil {
		log.Println("error when try to Copy file to buffer with detail :", err.Error())
		errInternal := fiber.ErrInternalServerError
		errInternal.Message = err.Error()
		return errInternal
	}

	// lalu upload file
	// disini kita letakin image nya pada folder `nbid-intermediate`
	uri, err := h.svc.Upload(context.Background(), buffer, "nbid-intermediate/zazhil/"+typeFile, "q_"+quality)
	if err != nil {
		log.Println("error when try to Upload with detail :", err.Error())
		errInternal := fiber.ErrInternalServerError
		errInternal.Message = err.Error()
		return errInternal
	}

	log.Println("upload success with URL :", uri)

	return helper.ResponseSuccess(c, true, "upload file success", http.StatusOK, helper.Payload{
		Url: uri,
	}, nil)
}
