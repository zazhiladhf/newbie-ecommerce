package search

import (
	"context"
	"encoding/json"
	"log"

	"github.com/meilisearch/meilisearch-go"
	"github.com/zazhiladhf/newbie-ecommerce/domain/product"
)

// type MeiliClient interface {
// 	GetAll(ctx context.Context) (productList []product.Product, err error)
// 	SyncPartial(ctx context.Context, productList []product.Product) (statusId int, err error)
// 	SyncAll(ctx context.Context, productList []product.Product) (statusId int, err error)
// }

type MeiliService struct {
	client *meilisearch.Client
}

func NewMeiliService(client *meilisearch.Client) MeiliService {
	return MeiliService{
		client: client,
	}
}

// GetAll implements SearchEngineInterface.
func (s MeiliService) Search(ctx context.Context, req SearchProductModel) (products []product.Product, err error) {
	resp, err := s.client.Index("products").Search(req.Query, &meilisearch.SearchRequest{
		HitsPerPage: int64(req.Limit),
		Page:        int64(req.Page),
		Filter:      req.Filter,
		Facets:      req.Facets,
	})

	if err != nil {
		return
	}

	hitByte, err := json.Marshal(resp.Hits)
	if err != nil {
		return
	}

	err = json.Unmarshal(hitByte, &products)
	if err != nil {
		return
	}

	// response.FacetDistribution adalah total
	// data dari hasil pengelompokan
	log.Printf("%+v\n", resp.FacetDistribution)
	return
}

// // SyncAll implements SearchEngineInterface.
func (s MeiliService) SyncAll(ctx context.Context, productList []product.Product) (statusId int, err error) {
	dataByte, err := json.Marshal(productList)
	if err != nil {
		return
	}
	task, err := s.client.Index("products").UpdateDocuments(dataByte, "id")
	if err != nil {
		return -1, err
	}
	return int(task.TaskUID), nil
}

// SyncPartial implements SearchEngineInterface.
func (s MeiliService) SyncPartial(ctx context.Context, productList []product.Product) (statusId int, err error) {
	dataByte, err := json.Marshal(productList)
	if err != nil {
		return
	}
	task, err := s.client.Index("products").UpdateDocuments(dataByte, "id")
	if err != nil {
		return -1, err
	}
	return int(task.TaskUID), nil
}
