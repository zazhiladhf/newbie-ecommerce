package search

import (
	"errors"

	"github.com/meilisearch/meilisearch-go"
)

func ConnectMeilisearch(host string, apiKey string) (client *meilisearch.Client, err error) {
	client = meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   host, // like http://localhost:7700
		APIKey: apiKey,
	})

	if !client.IsHealthy() {
		err = errors.New("meilisearch not healthy")
		return
	}
	// s.Meilisearch = client
	return
}
