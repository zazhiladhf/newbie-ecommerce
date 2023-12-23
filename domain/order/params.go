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

type WebhookInvoiceRequest struct {
	Id                 string    `json:"id"`
	ExternalId         string    `json:"external_id"`
	UserId             string    `json:"user_id"`
	IsHigh             bool      `json:"is_high"`
	Status             string    `json:"status"`
	MerchantName       string    `json:"merchant_name"`
	Amount             float64   `json:"amount"`
	PaidAmount         float64   `json:"paid_amount"`
	PayerEmail         string    `json:"payer_email"`
	Description        string    `json:"description"`
	UpdatedAt          time.Time `json:"updated_at"`
	CreatedAt          time.Time `json:"created_at"`
	PaidAt             time.Time `json:"paid_at"`
	Currency           string    `json:"currency"`
	PaymentChannel     string    `json:"payment_channel"`
	PaymentMethod      string    `json:"payment_method"`
	PaymentDestination string    `json:"payment_destination"`
	PaymentId          string    `json:"payment_id"`
}

func (r WebhookInvoiceRequest) parseToInvoice() Invoice {
	return Invoice{
		Id:                 r.Id,
		ExternalId:         r.ExternalId,
		UserId:             r.UserId,
		IsHigh:             r.IsHigh,
		Status:             r.Status,
		MerchantName:       r.MerchantName,
		Amount:             r.Amount,
		PaidAmount:         r.PaidAmount,
		PayerEmail:         r.PayerEmail,
		Description:        r.Description,
		UpdatedAt:          r.UpdatedAt,
		CreatedAt:          r.CreatedAt,
		PaidAt:             r.PaidAt,
		Currency:           r.Currency,
		PaymentChannel:     r.PaymentChannel,
		PaymentMethod:      r.PaymentMethod,
		PaymentDestination: r.PaymentDestination,
		PaymentId:          r.PaymentId,
	}
}
