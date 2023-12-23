package product

import (
	"github.com/google/uuid"
	"github.com/zazhiladhf/newbie-ecommerce/domain/merchant"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/helper"
)

type Product struct {
	Id           int    `db:"id"`
	Name         string `db:"name"`
	Description  string `db:"description"`
	Stock        int    `db:"stock"`
	Price        int    `db:"price"`
	Category     string `db:"category"`
	CategoryId   int    `db:"category_id"`
	ImageURL     string `db:"image_url"`
	MerchantId   int    `db:"merchant_id"`
	Sku          string `db:"sku"`
	CreatedAt    string `db:"created_at"`
	UpdatedAt    string `db:"updated_at"`
	MerchantName string `db:"merchant_name"`
	MerchantCity string `db:"merchant_city"`
	TotalData    int    `db:"total_data"`
}

func NewProduct() Product {
	return Product{}
}

func (p Product) newFromRequest(req CreateProductRequest) (product Product, err error) {
	product = Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		CategoryId:  req.CategoryId,
		ImageURL:    req.ImageURL,
		Sku:         uuid.New().String(),
	}

	err = product.validateRequestProduct()
	return
}

func (p Product) validateRequestProduct() (err error) {
	if p.Price == 0 {
		return helper.ErrEmptyPrice
	}

	if p.Price < 0 {
		return helper.ErrInvalidPrice
	}

	if p.Stock == 0 {
		return helper.ErrEmptyStock
	}

	if p.Stock < 0 {
		return helper.ErrInvalidStock
	}

	if p.Name == "" {
		return helper.ErrEmptyName
	}

	if p.Description == "" {
		return helper.ErrEmptyDescription
	}

	if p.ImageURL == "" {
		return helper.ErrEmptyImageURL
	}

	if p.CategoryId == 0 {
		return helper.ErrEmptyCategoryId
	}

	if err != nil {
		return helper.ErrNotFound
	}

	return
}

func (p Product) ProductResponse(products []Product) []GetListProductResponse {
	resp := []GetListProductResponse{}

	for _, product := range products {
		response := GetListProductResponse{
			Id:          product.Id,
			Sku:         product.Sku,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			Category:    product.Category,
			ImageURL:    product.ImageURL,
		}

		resp = append(resp, response)
	}

	return resp
}

func (p Product) ProductDetailResponse(product Product) GetDetailProductResponse {
	response := GetDetailProductResponse{
		Id:          product.Id,
		Sku:         product.Sku,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		Category:    product.Category,
		CategoryId:  product.CategoryId,
		ImageURL:    product.ImageURL,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}

	return response
}

func (p Product) ProductDetailUserPerspectiveResponse(product Product) GetDetailProductUserPerspectiveResponse {
	response := GetDetailProductUserPerspectiveResponse{
		Id:          product.Id,
		Sku:         product.Sku,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		Category:    product.Category,
		CategoryId:  product.CategoryId,
		Merchant: merchant.Merchant{
			Id:   product.MerchantId,
			Name: product.MerchantName,
			City: product.MerchantCity,
		},
		ImageURL:  product.ImageURL,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}

	return response
}
