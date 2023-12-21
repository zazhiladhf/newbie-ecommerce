package product

import (
	"context"
	"log"
)

type ProductRepository interface {
	InsertProduct(ctx context.Context, model Product) (id int, err error)
	// FindAllProducts(ctx context.Context) (list []Product, err error)
	FindProductByEmail(ctx context.Context, queryParam string, email string, limit int, page int) (list []Product, totalData int, err error)
}

type SearchEngineInterface interface {
	SyncPartial(ctx context.Context, productList []Product) (statusId int, err error)
	// SyncAll(ctx context.Context, productList []Product) (statusId int, err error)
	// GetAll(ctx context.Context) (productList []Product, err error)
}

// type AuthRepository interface {
// 	GetAuthByEmail(ctx context.Context, email string) (auth auth.Auth, err error)
// }

type Service struct {
	repo   ProductRepository
	search SearchEngineInterface
	// authRepo    AuthRepository
}

func NewService(repo ProductRepository, search SearchEngineInterface) Service {
	return Service{
		repo:   repo,
		search: search,
		// authRepo:    authRepo,
	}
}

func (s Service) createProduct(ctx context.Context, req Product) (err error) {
	// product, err := newFromRequest(req)
	// if err != nil {
	// 	return
	// }
	// product.AuthEmail = token

	// if err = req.(); err != nil {
	// 	log.Println("erro when try to validate request with error")
	// 	return
	// }GetAuthByEmail

	// auth, err := s.authRepo.GetAuthByEmail(ctx, token)
	// if err != nil {
	// 	log.Println("error when try to get auth by email with error :", err.Error(), token)
	// 	return err
	// }

	// model := Product{
	// 	Name:       product.Name,
	// 	Stock:      product.Stock,
	// 	Price:      product.Price,
	// 	CategoryId: product.CategoryId,
	// 	ImageURL:   product.ImageURL,
	// 	AuthEmail:  auth.Email,
	// }

	id, err := s.repo.InsertProduct(ctx, req)
	if err != nil {
		log.Println("error when try to insert product to database with error :", err.Error())
		return err
	}

	req.Id = id

	status, err := s.search.SyncPartial(ctx, []Product{req})
	if err != nil {
		log.Println("error when try to sync pertial with error :", err.Error())
		return err
	}

	log.Println("status id", status)

	return
}

// func (s Service) GetProducts(ctx context.Context) (list []Product, err error) {
// 	listProducts, err := s.repo.FindAll(ctx)
// 	if err != nil {
// 		if err == ErrCategoriesNotFound {
// 			return []Product{}, nil
// 		}
// 		return nil, err
// 	}

// 	return listProducts, nil

// }

func (s Service) GetProductsByEmail(ctx context.Context, queryParam string, email string, limit int, page int) (resp []GetListProductResponse, totalData int, err error) {
	list, totalData, err := s.repo.FindProductByEmail(ctx, queryParam, email, limit, page)
	if err != nil {
		return
	}

	resp = NewProduct().ProductResponse(list)

	return
}

// func (s Service) UpdateProduct(ctx context.Context, req Product, param int) (err error) {
// 	if err = req.Validate(); err != nil {
// 		log.Println("erro when try to validate request with error")
// 		return
// 	}

// 	if err = s.repo.Update(ctx, param, req); err != nil {
// 		log.Println("error when try to Update to database with error :", err.Error())
// 		return
// 	}
// 	return

// }

// func (s Service) DeleteProduct(ctx context.Context, model Product, param int) (err error) {
// 	if err = s.repo.DeleteProduct(ctx, model, param); err != nil {
// 		log.Println("error when try to Delete to database with error :", err.Error())
// 		return
// 	}
// 	return

// }
