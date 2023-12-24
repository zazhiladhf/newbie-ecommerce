package meili

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConnectionMeilisearch(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client, err := ConnectMeilisearch("http://localhost:7700", "ThisIsMasterKey")
		require.Nil(t, err)
		require.NotNil(t, client)
	})

}
