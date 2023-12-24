package database

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zazhiladhf/newbie-ecommerce/config"
)

func TestConnectionPostgres(t *testing.T) {
	cfg := config.DB{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "74712331",
		Name:     "newbie-ecommerce",
	}

	t.Run("success connect", func(t *testing.T) {
		db, err := ConnectPostgresSqlx(cfg)
		require.Nil(t, err)
		require.NotNil(t, db)
	})

	cfg = config.DB{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "invalid password",
		Name:     "newbie-ecommerce",
	}

	t.Run("invalid password", func(t *testing.T) {
		_, err := ConnectPostgresSqlx(cfg)
		require.NotNil(t, err)
	})
}
