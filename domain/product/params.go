package product

type CreateProductRequest struct {
	Name       string `json:"name"`
	ImageURL   string `json:"image_url"`
	Stock      int    `json:"stock"`
	Price      int    `json:"price"`
	CategoryId int    `json:"category_id"`
}

type GetListProductResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Stock    int    `json:"stock"`
	Category string `json:"category"`
	ImageUrl string `json:"image_url"`
}
