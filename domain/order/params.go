package order

import (
	"time"
)

type CheckoutRequest struct {
	AuthId    int
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CheckoutResponse struct {
	InvoiceUrl string `json:"invoice_url"`
}

type GetOrderHistoriesRequest struct {
	Limit  int `query:"limit"`
	Page   int `query:"page"`
	AuthId int
}

type GetOrderHistoriesResponse struct {
	CreatedAt   time.Time    `json:"created_at"`
	GrandTotal  float32      `json:"grand_total"`
	Id          string       `json:"id"`
	InvoiceUrl  string       `json:"invoice_url"`
	Merchant    MerchantData `json:"merchant"`
	PlatformFee float32      `json:"paltfrom_fee"`
	Price       float32      `json:"price"`
	Product     ProductData  `json:"product"`
	Quantity    int          `json:"quantity"`
	Status      string       `json:"status"`
	SubTotal    float32      `json:"sub_total"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Uuid        string       `json:"uuid"`
}

type MerchantData struct {
	Id       int    `json:"id"`
	ImageUrl string `json:"image_url"`
	Name     string `json:"name"`
}

type ProductData struct {
	Id          int     `json:"id"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	ImageUrl    string  `json:"image_url"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Stock       int     `json:"stock"`
}

type GeListOrdersResponse struct {
	CreatedAt   time.Time   `json:"created_at"`
	GrandTotal  float32     `json:"grand_total"`
	Id          string      `json:"id"`
	PlatformFee float32     `json:"paltfrom_fee"`
	Price       float32     `json:"price"`
	Product     ProductData `json:"product"`
	Quantity    int         `json:"quantity"`
	Status      string      `json:"status"`
	SubTotal    float32     `json:"sub_total"`
	UpdatedAt   time.Time   `json:"updated_at"`
	Uuid        string      `json:"uuid"`
}
