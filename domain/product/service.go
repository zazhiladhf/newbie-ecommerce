package product

import (
	"context"
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/zazhiladhf/newbie-ecommerce/domain/auth"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/helper"
)

type ProductRepositoryTx interface {
	GetProductByIdForUpdate(ctx context.Context, id int) (product Product, err error)
	UpdateProductStok(ctx context.Context, product Product) (err error)
}

type ProductRepository interface {
	InsertProduct(ctx context.Context, product Product) (id int, err error)
	GetAllProducts(ctx context.Context) (productList []Product, err error)
	GetProducts(ctx context.Context, queryParam string, id int, limit int, page int) (list []Product, totalData int, err error)
	GetProductById(ctx context.Context, id int) (product Product, err error)
	UpdateProduct(ctx context.Context, product Product) (err error)
	GetProductBySku(ctx context.Context, sku string) (product Product, err error)
	CheckoutProduct(ctx context.Context, id, quantity int) (tx *sqlx.Tx, err error)
}

type SearchEngineInterface interface {
	GetAll(ctx context.Context) (productList []Product, err error)
	SyncPartial(ctx context.Context, productList []Product) (statusId int, err error)
	SyncAll(ctx context.Context, productList []Product) (statusId int, err error)
}

type AuthRepository interface {
	GetAuthByEmail(ctx context.Context, email string) (auth auth.Auth, err error)
}

type Service struct {
	pRepo  ProductRepository
	aRepo  AuthRepository
	search SearchEngineInterface
	// authRepo    AuthRepository
}

func NewService(pRepo ProductRepository, aRepo AuthRepository, search SearchEngineInterface) Service {
	return Service{
		pRepo:  pRepo,
		aRepo:  aRepo,
		search: search,
	}
}

func (s Service) CreateProductByMerchant(ctx context.Context, req Product, email string) (err error) {
	auth, err := s.aRepo.GetAuthByEmail(ctx, email)
	if err != nil {
		log.Println("error when try to get auth by email with error", err)
		return
	}

	if auth.Role != "Merchant" {
		return helper.ErrInvalidRole
	}

	id, err := s.pRepo.InsertProduct(ctx, req)
	if err != nil {
		log.Println("error when try to insert product to database with error :", err.Error())
		return err
	}
	req.Id = id

	status, err := s.search.SyncPartial(ctx, []Product{req})
	if err != nil {
		log.Println("error when try to sync partial with error :", err.Error())
		return err
	}

	log.Println("status id", status)

	return
}

func (s Service) GetAllProducts(ctx context.Context) (taskId int, err error) {
	products, err := s.pRepo.GetAllProducts(ctx)
	if err != nil {
		return -1, err
	}

	taskId, err = s.search.SyncAll(ctx, products)
	if err != nil {
		return -1, err
	}

	log.Println("status id", taskId)
	return
}

func (s Service) syncAll(ctx context.Context) (taskId int, err error) {
	productList, err := s.pRepo.GetAllProducts(ctx)
	if err != nil {
		return -1, err
	}
	taskId, err = s.search.SyncAll(ctx, productList)
	if err != nil {
		return -1, err
	}

	log.Println("status id", taskId)
	return

}

func (s Service) GetListProductsMerchant(ctx context.Context, queryParam string, email string, limit int, page int) (resp []GetListProductResponse, totalData int, err error) {
	auth, err := s.aRepo.GetAuthByEmail(ctx, email)
	if err != nil {
		log.Println("error when try to get auth by email with error", err)
		return
	}

	if auth.Role != "Merchant" {
		return nil, 0, helper.ErrInvalidRole
	}

	list, totalData, err := s.pRepo.GetProducts(ctx, queryParam, auth.Id, limit, page)
	if err != nil {
		return
	}

	resp = NewProduct().ProductResponse(list)

	return
}

func (s Service) GetDetailProductById(ctx context.Context, id int, email string) (resp GetDetailProductResponse, err error) {
	auth, err := s.aRepo.GetAuthByEmail(ctx, email)
	if err != nil {
		log.Println("error when try to get auth by email with error", err)
		return
	}

	if auth.Role != "Merchant" {
		return resp, helper.ErrInvalidRole
	}

	product, err := s.pRepo.GetProductById(ctx, id)
	if err != nil {
		if sql.ErrNoRows == err {
			err = helper.ErrNotFound
		}
		return
	}

	resp = NewProduct().ProductDetailResponse(product)

	return
}

func (s Service) UpdateProduct(ctx context.Context, req Product, email string) (err error) {
	auth, err := s.aRepo.GetAuthByEmail(ctx, email)
	if err != nil {
		log.Println("error when try to get auth by email with error", err)
		return
	}

	if auth.Role != "Merchant" {
		return helper.ErrInvalidRole
	}

	product, err := s.pRepo.GetProductById(ctx, req.Id)
	if err != nil {
		log.Println("error when try to get product by id with error", err)
		if sql.ErrNoRows == err {
			err = helper.ErrNotFound
		}
		return
	}

	if product.Id == 0 {
		return helper.ErrNotFound
	}

	err = s.pRepo.UpdateProduct(ctx, req)
	if err != nil {
		log.Println("error when try to update product with error", err)
		return
	}

	return
}

func (s Service) GetDetailProductUserPerspective(ctx context.Context, sku string) (resp GetDetailProductUserPerspectiveResponse, err error) {
	product, err := s.pRepo.GetProductBySku(ctx, sku)
	if err != nil {
		if sql.ErrNoRows == err {
			err = helper.ErrNotFound
		}
		return
	}

	resp = NewProduct().ProductDetailUserPerspectiveResponse(product)

	return
}
