package order

type CheckoutRequest struct {
	AuthId    int
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CheckoutResponse struct {
	InvoiceUrl string `json:"invoice_url"`
}
