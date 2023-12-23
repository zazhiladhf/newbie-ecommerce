package search

import (
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
