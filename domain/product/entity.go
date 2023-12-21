package product

import (
	"errors"
)

var (
	ErrEmptyName       = errors.New("name is required")
	ErrEmptyImageURL   = errors.New("image_url is required")
	ErrEmptyStock      = errors.New("stock is required")
	ErrEmptyPrice      = errors.New("price is required")
	ErrEmptyCategoryId = errors.New("category_id is required")
	ErrNotFound        = errors.New("product not found")
	ErrRepository      = errors.New("error repository")
	ErrInternalServer  = errors.New("unknown error")
)

type Product struct {
	Id         int    `db:"id"`
	Name       string `db:"name"`
	Stock      int    `db:"stock"`
	Price      int    `db:"price"`
	Category   string `db:"category"`
	CategoryId int    `db:"category_id"`
	ImageURL   string `db:"image_url"`
	AuthEmail  string `db:"email_auth"`
}

func NewProduct() Product {
	return Product{}
}

func (p Product) newFromRequest(req CreateProductRequest) (product Product, err error) {
	product = Product{
		Name:       req.Name,
		ImageURL:   req.ImageURL,
		Stock:      req.Stock,
		Price:      req.Price,
		CategoryId: req.CategoryId,
	}

	err = product.validateRequestProduct()
	return
}

func (p Product) validateRequestProduct() (err error) {
	if p.Name == "" {
		return ErrEmptyName
	}

	if p.ImageURL == "" {
		return ErrEmptyImageURL
	}

	if p.Stock == 0 {
		return ErrEmptyStock
	}

	if p.Price == 0 {
		return ErrEmptyPrice
	}

	if p.CategoryId == 0 {
		return ErrEmptyCategoryId
	}

	if err != nil {
		return ErrNotFound
	}

	return
}

func (p Product) ProductResponse(products []Product) []GetListProductResponse {
	resp := []GetListProductResponse{}

	for _, product := range products {
		response := GetListProductResponse{
			ID:       product.Id,
			Name:     product.Name,
			Price:    product.Price,
			Stock:    product.Stock,
			Category: product.Category,
			ImageUrl: product.ImageURL,
		}

		resp = append(resp, response)
	}

	return resp
}
