package search

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/helper"
)

type MeiliHandler struct {
	svc MeiliService
}

func NewMeiliHandler(svc MeiliService) MeiliHandler {
	return MeiliHandler{
		svc: svc,
	}
}

func (h MeiliHandler) SearchProduct(c *fiber.Ctx) error {
	query := c.Query("query")
	// categoryIDs := c.Query("category")
	// categoryIDList := strings.Split(categoryIDs, ",")
	limit := c.Query("limit", "5")
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

	myProducts, err := h.svc.Search(c.UserContext(), SearchProductModel{
		Query:  query,                // kita buat typo sedikit
		Facets: []string{"category"}, // lalu tampilkan total data berdasarkan category
		Pagination: Pagination{
			Limit: limitInt,
			Page:  pageInt,
		},
	})

	if err != nil {
		log.Println("error when try to search with error", err)
	}

	// for _, p := range myProducts {
	// 	// hanya untuk pembatas
	// 	div := strings.Repeat("=", 11)
	// 	log.Println(div, "[PRODUCT]", div)
	// 	log.Printf("%+v\n", p)
	// 	log.Println(div + div + div)
	// }

	// resp, totalData, err := h.svc.SearchProduct(c.UserContext(), query, limitInt, pageInt)
	// if err != nil {
	// 	return helper.ResponseError(c, err)
	// }
	return helper.ResponseSuccess(c, true, "search products success", http.StatusOK, myProducts, nil)
}
