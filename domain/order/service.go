package order

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/zazhiladhf/newbie-ecommerce/config"
	"github.com/zazhiladhf/newbie-ecommerce/domain/product"
	"github.com/zazhiladhf/newbie-ecommerce/domain/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, payload Order) (order Order, err error)
	GetOrderHistories(ctx context.Context, limit, page, userId int) (orders []Order, totalPage int, err error)
}

type PaymentRepository interface {
	GetBalance(ctx context.Context) (balance float64, err error)
	CreateInvoice(ctx context.Context, req Order) (invoiceUrl string, err error)
}

type service struct {
	productRepo product.ProductRepository
	userRepo    user.UserRepository
	orderRepo   OrderRepository
	paymentRepo PaymentRepository
}

func NewService(productRepo product.ProductRepository, userRepo user.UserRepository, orderRepo OrderRepository, paymentRepo PaymentRepository) service {
	return service{
		productRepo: productRepo,
		orderRepo:   orderRepo,
		paymentRepo: paymentRepo,
		userRepo:    userRepo,
	}
}

func (s service) GetOrderHistories(ctx context.Context, req GetOrderHistoriesRequest) (resp []GetOrderHistoriesResponse, totalPage int, err error) {
	user, err := s.userRepo.GetUserByAuthId(ctx, req.AuthId)
	if err != nil {
		log.Println("error when try to get user with error", err)
		return resp, totalPage, err
	}

	orders, totalPage, err := s.orderRepo.GetOrderHistories(ctx, req.Limit, req.Page, user.Id)
	if err != nil {
		return nil, totalPage, err
	}

	for _, v := range orders {
		var platformFee float32 = 0

		for _, x := range v.AdditionalFee {
			if x.Type == "plarform_fee" {
				platformFee = x.Value
			}
		}

		temp := GetOrderHistoriesResponse{
			CreatedAt:  v.CreatedAt,
			GrandTotal: v.GrandTotal,
			Id:         v.Id.Hex(),
			InvoiceUrl: v.InvoiceUrl,
			Merchant: MerchantData{
				Id:   v.Product.MerchantId,
				Name: v.Product.MerchantName,
			},
			PlatformFee: platformFee,
			Price:       v.Price,
			Product: ProductData{
				Id:          v.Product.Id,
				Category:    v.Product.Category,
				Description: v.Product.Description,
				ImageUrl:    v.Product.ImageURL,
				Name:        v.Product.Name,
				Price:       float32(v.Product.Price),
				Stock:       v.Product.Stock,
			},
			Quantity:  v.Quantity,
			Status:    v.Status,
			SubTotal:  v.SubTotal,
			UpdatedAt: v.UpdatedAt,
			Uuid:      v.Uuid,
		}

		resp = append(resp, temp)
	}

	return resp, totalPage, nil
}

func (s service) Checkout(ctx context.Context, req CheckoutRequest) (resp CheckoutResponse, err error) {
	user, err := s.userRepo.GetUserByAuthId(ctx, req.AuthId)
	if err != nil {
		log.Println("error when try to get user with error", err)
		return resp, err
	}

	// start transaction
	tx, err := s.productRepo.CheckoutProduct(ctx, req.ProductId, req.Quantity)
	if err != nil {
		log.Println("error when try to checkout product with error", err)
		return resp, err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			log.Println("Error occurred, rolling back transaction:", err)
			tx.Rollback()
		} else {
			err = tx.Commit()
			if err != nil {
				log.Println("error committing transaction", err)
			}
		}
	}()

	// get product by id
	product, err := s.productRepo.GetProductById(ctx, 1)
	if err != nil {
		log.Println("error when try to get product by id with error", err)
		return resp, err
	}

	payload := Order{
		Id:              primitive.NewObjectID(),
		ProductId:       product.Id,
		ExternalId:      uuid.NewString(),
		UserId:          user.Id,
		Uuid:            uuid.NewString(),
		Price:           float32(product.Price),
		AdditionalFee:   nil,
		SubTotal:        float32(product.Price * req.Quantity),
		UserEmail:       user.Auth.Email,
		UserName:        user.Name,
		PhoneNumber:     user.PhoneNumber,
		Description:     product.Description,
		InvoiceDuration: config.Cfg.Payment.InvoiceDuration,
		Product:         product,
		Quantity:        req.Quantity,
		Status:          "UNPAID",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	payload.setTotal()

	// create order
	order, err := s.orderRepo.CreateOrder(ctx, payload)
	if err != nil {
		log.Println("error when try to create order with error", err)
		return resp, err
	}

	// i have to create invoice first
	invoiceUrl, err := s.paymentRepo.CreateInvoice(ctx, order)
	if err != nil {
		log.Println("error when try to get product by id with error", err)
		return resp, err
	}

	return CheckoutResponse{
		InvoiceUrl: invoiceUrl,
	}, err
}
