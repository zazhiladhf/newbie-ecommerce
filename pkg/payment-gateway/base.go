package paymentgateway

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/xendit/xendit-go/v3"
	"github.com/xendit/xendit-go/v3/invoice"
	"github.com/zazhiladhf/newbie-ecommerce/config"
)

type Xendit struct {
	client    *xendit.APIClient
	secretKey string
	redirect  redirect
}

type redirect struct {
	Success string `json:"success_redirect_url"`
	Failure string `json:"failure_redirect_url"`
}

type Balance struct {
	Balance float64 `json:"balance"`
}

type Invoice struct {
	ExternalId      string   `json:"external_id"`
	Amount          float32  `json:"amount"`
	Description     string   `json:"description"`
	InvoiceDuration int      `json:"invoice_duration"`
	Customer        Customer `json:"customer"`
	Items           []Item   `json:"items"`
	AdditionalFee   []Fee    `json:"fees"`
}

type Customer struct {
	Name        string `json:"given_names"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"mobile_number"`
}

type Item struct {
	Name     string  `json:"name"`
	Quantity float32 `json:"quantity"`
	Price    float32 `json:"price"`
	Category string  `json:"category"`
	URL      string  `json:"url"`
}

func (i Item) parseToXenditInvoiceItem() invoice.InvoiceItem {
	return *invoice.NewInvoiceItem(i.Name, i.Price, i.Quantity)
}

type Fee struct {
	Type  string  `json:"type"`
	Value float32 `json:"value"`
}

func (f Fee) parseToXenditInvoiceFee() invoice.InvoiceFee {
	return *invoice.NewInvoiceFee(f.Type, f.Value)
}

func (i Invoice) parseToXenditRequest(successRedirectURI, failureRedirectURI string) *invoice.CreateInvoiceRequest {

	duration := strconv.Itoa(i.InvoiceDuration)
	shouldSendEmail := true

	items := []invoice.InvoiceItem{}

	for _, item := range i.Items {
		// panggil method parse invoice item
		items = append(items, item.parseToXenditInvoiceItem())
	}

	fees := []invoice.InvoiceFee{}
	for _, fee := range i.AdditionalFee {
		// panggil method parse invoice fee
		fees = append(fees, fee.parseToXenditInvoiceFee())
	}
	return &invoice.CreateInvoiceRequest{
		ExternalId:      i.ExternalId,
		Amount:          i.Amount,
		PayerEmail:      &i.Customer.Email,
		Description:     &i.Description,
		InvoiceDuration: &duration,
		ShouldSendEmail: &shouldSendEmail,
		Customer: &invoice.CustomerObject{
			GivenNames:  *invoice.NewNullableString(&i.Customer.Name),
			Surname:     *invoice.NewNullableString(&i.Customer.Surname),
			Email:       *invoice.NewNullableString(&i.Customer.Email),
			PhoneNumber: *invoice.NewNullableString(&i.Customer.PhoneNumber),
		},
		SuccessRedirectUrl: &successRedirectURI,
		FailureRedirectUrl: &failureRedirectURI,
		Items:              items,
		Fees:               fees,
	}
}

type InvoiceResponse struct {
	InvoiceURL string `json:"invoice_url"`
}

func newInvoiceResponseFromXenditResponse(invoice invoice.Invoice) (resp InvoiceResponse) {
	return InvoiceResponse{
		InvoiceURL: invoice.InvoiceUrl,
	}
}

func NewXendit(secretKey string) Xendit {
	client := xendit.NewClient(secretKey)
	return Xendit{
		client:    client,
		secretKey: secretKey,
	}
}

func (x *Xendit) SetConfig(cfg config.Payment) *Xendit {
	x.redirect.Success = cfg.Redirect.Success
	x.redirect.Failure = cfg.Redirect.Failure

	return x
}

func (x Xendit) GetBalance(ctx context.Context) (myBalance Balance, err error) {
	// hit api xendit, dan akan nge return 3 value
	balance, httpResp, errXendit := x.client.BalanceApi.GetBalance(ctx).Execute()
	if err != nil {
		b, _ := json.Marshal(errXendit.FullError())
		fmt.Printf("Error when try to get balance with error detail : %v\n", string(b))
		fmt.Printf("Full HTTP response: %v\n", httpResp)
		err = errXendit
		return
	}

	myBalance = Balance{
		Balance: float64(balance.GetBalance()),
	}
	return
}

func (x Xendit) CreateInvoice(ctx context.Context, req Invoice) (resp InvoiceResponse, err error) {
	// karena request nya adalah object Invoice yang kita buat sendiri
	// maka kita perlu parse ke request yg di butuhkan oleh xendit
	// disini menggunakan redirect success dan failure sebagai parameter
	xenditReq := req.parseToXenditRequest(x.redirect.Success, x.redirect.Failure)

	// setelah sudah dibuat object xendit request, lalu panggil xendit client
	invoice, httpResp, errXendit := x.client.InvoiceApi.CreateInvoice(ctx).CreateInvoiceRequest(*xenditReq).Execute()

	// jika muncul error xendit
	if errXendit != nil {
		b, _ := json.Marshal(errXendit.FullError())
		fmt.Printf("Error when try to get balance with error detail : %v\n", string(b))
		fmt.Printf("Full HTTP response: %v\n", httpResp)
		err = errXendit
		return
	}

	// jika invoice nya nil
	if invoice == nil {
		err = errors.New("invoice xendit is nil")
		fmt.Printf("Full HTTP response: %v\n", httpResp)
		fmt.Printf("Error when try to get balance with error detail : %v\n", err.Error())
		return
	}

	// return invoice
	resp = newInvoiceResponseFromXenditResponse(*invoice)
	return
}
