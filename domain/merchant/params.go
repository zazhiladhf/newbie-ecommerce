package merchant

type RequestBodyCreateMerchant struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	ImageUrl    string `json:"image_url"`
	City        string `json:"city"`
}

type GetMerchantResponse struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	ImageUrl    string `json:"image_url"`
	City        string `json:"city"`
}
