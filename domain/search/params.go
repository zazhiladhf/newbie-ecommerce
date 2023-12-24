package search

type SearchProductModel struct {
	HitsPerPage int      `json:"hits_per_page"`
	Query       string   `json:"query"`
	Facets      []string `json:"facets"`
	Filter      string   `json:"filter"`
	Sort        []string `json:"sort"`

	Pagination
}

type Pagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}
