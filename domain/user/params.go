package user

type RequestBodyCreateProfileUser struct {
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	ImageUrl    string `json:"image_url"`
	// Role        string `json:"role"`
	// CreatedBy   string `json:"created_by"`
}

type GetUserResponse struct {
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	ImageUrl    string `json:"image_url"`
}
