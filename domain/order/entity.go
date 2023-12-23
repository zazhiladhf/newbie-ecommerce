package order

import (
	"time"

	"github.com/zazhiladhf/newbie-ecommerce/domain/product"
	paymentgateway "github.com/zazhiladhf/newbie-ecommerce/pkg/payment-gateway"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	Id          primitive.ObjectID `bson:"_id"`
	ProductId   int                `bson:"product_id"`
	ExternalId  string             `bson:"external_id"`
	UserId      int                `bson:"user_id"`
	Uuid        string             `bson:"uuid"`
	UserEmail   string             `json:"user_email"`
	UserName    string             `json:"user_name"`
	PhoneNumber string             `json:"phone_number"`

	Quantity        int             `bson:"quantity"`
	Price           float32         `bson:"price"`
	SubTotal        float32         `bson:"sub_total"`
	AdditionalFee   []AdditionalFee `bson:"additional_feee"`
	GrandTotal      float32         `bson:"grand_total"`
	InvoiceDuration int             `json:"invoice_duration"`
	InvoiceUrl      string          `bson:"invoice_url"`
	Description     string          `bson:"description,omitempty"`
	Status          string          `bson:"status"`
	Product         product.Product `bson:"product"`
	CreatedAt       time.Time       `bson:"created_at"`
	UpdatedAt       time.Time       `bson:"updated_at"`
}

type AdditionalFee struct {
	Value float32 `bson:"value"`
	Type  string  `bson:"type"`
}

func (o *Order) setTotal() *Order {
	var total float32 = o.SubTotal

	for _, fee := range o.AdditionalFee {
		total += fee.Value
	}

	o.GrandTotal = total

	return o
}

func (o Order) parseToInvoicePaymentRequest() paymentgateway.Invoice {
	var items = []paymentgateway.Item{
		{
			Name:     o.Product.Name,
			Category: o.Product.Category,
			Quantity: float32(o.Quantity),
			Price:    float32(o.Product.Price),
		},
	}

	var fees = []paymentgateway.Fee{}
	for _, fee := range o.AdditionalFee {
		fees = append(fees, paymentgateway.Fee{
			Value: fee.Value,
			Type:  fee.Type,
		})
	}

	return paymentgateway.Invoice{
		ExternalId:      o.ExternalId,
		Amount:          o.GrandTotal,
		Description:     o.Description,
		InvoiceDuration: o.InvoiceDuration,
		Customer: paymentgateway.Customer{
			Email:       o.UserEmail,
			Name:        o.UserName,
			Surname:     o.UserName,
			PhoneNumber: o.PhoneNumber,
		},
		Items:         items,
		AdditionalFee: fees,
	}
}
