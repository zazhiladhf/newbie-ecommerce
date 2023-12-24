package product

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zazhiladhf/newbie-ecommerce/config"
	"github.com/zazhiladhf/newbie-ecommerce/domain/auth"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/database"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/meili"
)

var svc Service

func init() {
	cfg := config.DB{
		Host:     "103.193.176.215",
		Port:     "5432",
		User:     "postgres",
		Password: "74712331",
		Name:     "newbie-ecommerce",
	}

	db, err := database.ConnectPostgresSqlx(cfg)
	if err != nil {
		panic(err)
	}

	client, err := meili.ConnectMeilisearch("http://localhost:7700", "ThisIsMasterKey")
	if err != nil {
		panic(err)
	}

	meiliRepo := NewMeiliRepository(client)
	pRepo := NewPostgresSQLXRepository(db)
	aRepo := auth.NewPostgreSqlxRepository(db)
	svc = NewService(pRepo, aRepo, meiliRepo)
}

func TestCreate(t *testing.T) {
	ctx := context.Background()
	req := Product{
		Name:        "Book of Avatar",
		Description: "deskripsi",
		CategoryId:  1,
		Price:       10_000,
		Stock:       100,
		ImageURL:    "image",
	}
	email := "zazhil@gmail.com"
	err := svc.CreateProductByMerchant(ctx, req, email)
	require.Nil(t, err)
}

// func TestGetAll(t *testing.T) {
// 	ctx := context.Background()
// 	productList, total, err := svc.GetListProductsMerchant(ctx)
// 	log.Println(productList)
// 	require.Nil(t, err)
// 	require.NotNil(t, productList, total)
// }

func TestSyncAll(t *testing.T) {
	ctx := context.Background()
	taskId, err := svc.syncAll(ctx)
	log.Println(taskId)
	require.Nil(t, err)
}
