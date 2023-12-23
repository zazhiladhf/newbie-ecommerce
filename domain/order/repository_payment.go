package order

import (
	"context"

	paymentgateway "github.com/zazhiladhf/newbie-ecommerce/pkg/payment-gateway"
)

type PaymentAdapter interface {
	GetBalance(ctx context.Context) (balance paymentgateway.Balance, err error)
	CreateInvoice(ctx context.Context, req paymentgateway.Invoice) (resp paymentgateway.InvoiceResponse, err error)
}

type paymentGatewayRepository struct {
	payment PaymentAdapter
}

func newPaymentGatewayRepository(payment PaymentAdapter) paymentGatewayRepository {
	return paymentGatewayRepository{
		payment: payment,
	}
}

func (p paymentGatewayRepository) GetBalance(ctx context.Context) (balance float64, err error) {
	myBalance, err := p.payment.GetBalance(ctx)
	if err != nil {
		return
	}

	balance = myBalance.Balance
	return
}

func (p paymentGatewayRepository) CreateInvoice(ctx context.Context, order Order) (invoiceUrl string, err error) {
	invoiceRequest := order.parseToInvoicePaymentRequest()

	resp, err := p.payment.CreateInvoice(ctx, invoiceRequest)
	if err != nil {
		return
	}

	return resp.InvoiceURL, nil
}
