package product

import (
	"context"
	"encoding/json"

	"github.com/meilisearch/meilisearch-go"
)

type MeiliRepository struct {
	client *meilisearch.Client
}

func NewMeiliRepository(clent *meilisearch.Client) MeiliRepository {
	return MeiliRepository{
		client: clent,
	}
}

// // GetAll implements SearchEngineInterface.
// func (m meiliRepository) GetAll(ctx context.Context) (productList []Product, err error) {
// 	resp, err := m.client.Index("products").Search("book", &meilisearch.SearchRequest{
// 		Facets: []string{"category"},
// 	})
// 	if err != nil {
// 		return
// 	}

// 	hitByte, err := json.Marshal(resp.Hits)
// 	if err != nil {
// 		return
// 	}

// 	err = json.Unmarshal(hitByte, &productList)
// 	if err != nil {
// 		return
// 	}

// 	// response.FacetDistribution adalah total
// 	// data dari hasil pengelompokan
// 	log.Printf("%+v\n", resp.FacetDistribution)
// 	return
// 	// panic("")
// }

// // SyncAll implements SearchEngineInterface.
// func (m meiliRepository) SyncAll(ctx context.Context, productList []Product) (statusId int, err error) {
// 	dataByte, err := json.Marshal(productList)
// 	if err != nil {
// 		return
// 	}
// 	task, err := m.client.Index("products").UpdateDocuments(dataByte, "id")
// 	if err != nil {
// 		return -1, err
// 	}
// 	return int(task.TaskUID), nil
// }

// SyncPartial implements SearchEngineInterface.
func (m MeiliRepository) SyncPartial(ctx context.Context, productList []Product) (statusId int, err error) {
	dataByte, err := json.Marshal(productList)
	if err != nil {
		return
	}
	task, err := m.client.Index("products").UpdateDocuments(dataByte, "id")
	if err != nil {
		return -1, err
	}
	return int(task.TaskUID), nil
}
