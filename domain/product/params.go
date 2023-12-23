package product

import "github.com/zazhiladhf/newbie-ecommerce/domain/merchant"

type CreateProductRequest struct {
	Name        string `json:"name"`
	Description string `db:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	CategoryId  int    `json:"category_id"`
	ImageURL    string `json:"image_url"`
}

type GetListProductResponse struct {
	Id          int    `json:"id"`
	Sku         string `db:"sku"`
	Name        string `json:"name"`
	Description string `db:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Category    string `json:"category"`
	ImageURL    string `json:"image_url"`
}

type UpdateProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	CategoryId  int    `json:"category_id"`
	ImageURL    string `json:"image_url"`
}

type GetDetailProductResponse struct {
	Id          int    `json:"id"`
	Sku         string `json:"sku"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Category    string `json:"category"`
	CategoryId  int    `json:"category_id"`
	ImageURL    string `json:"image_url"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type GetDetailProductUserPerspectiveResponse struct {
	Id          int                       `json:"id"`
	Sku         string                    `json:"sku"`
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	Price       int                       `json:"price"`
	Stock       int                       `json:"stock"`
	Category    string                    `json:"category"`
	CategoryId  int                       `json:"category_id"`
	Merchant    merchant.MerchantResponse `json:"merchant"`
	ImageURL    string                    `json:"image_url"`
	CreatedAt   string                    `json:"created_at"`
	UpdatedAt   string                    `json:"updated_at"`
}
