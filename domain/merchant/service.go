package merchant

import (
	"context"
	"log"

	"github.com/zazhiladhf/newbie-ecommerce/domain/auth"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/helper"
)

type postgreSqlxRepository interface {
	InsertMerchant(ctx context.Context, merchant Merchant) (err error)
	GetMerchantById(ctx context.Context, id int) (merchant Merchant, err error)
	UpdateMerchant(ctx context.Context, req Merchant) (err error)
}

type authRepository interface {
	GetAuthByEmail(ctx context.Context, email string) (auth auth.Auth, err error)
}

type MerchantService struct {
	merchantRepo postgreSqlxRepository
	authRepo     authRepository
}

func NewService(merchantRepo postgreSqlxRepository, authrepo authRepository) MerchantService {
	return MerchantService{
		merchantRepo: merchantRepo,
		authRepo:     authrepo,
	}
}

func (s MerchantService) CreateProfileMerchant(ctx context.Context, req Merchant, email string) (merchant Merchant, err error) {
	auth, err := s.authRepo.GetAuthByEmail(ctx, email)
	if err != nil {
		log.Println("error when try to get auth by email with error", err)
		return
	}

	if auth.Role != "Merchant" {
		log.Println("auth:", auth)
		return merchant, helper.ErrInvalidRole
	}

	err = s.merchantRepo.InsertMerchant(ctx, req)
	if err != nil {
		log.Println("error when try to insert merchant to db with error", err)
		return
	}

	return
}

func (s MerchantService) GetMerchantById(ctx context.Context, id int) (resp GetMerchantResponse, err error) {
	merchant, err := s.merchantRepo.GetMerchantById(ctx, id)
	if err != nil {
		log.Println("error when try to get user by id with error", err)
		return
	}

	if merchant.Id == 0 {
		return resp, helper.ErrUserNotFound
	}

	resp = NewMerchant().MerchantResponse(merchant)

	return
}

func (s MerchantService) UpdateProfileMerchant(ctx context.Context, req Merchant, email string) (err error) {
	auth, err := s.authRepo.GetAuthByEmail(ctx, email)
	if err != nil {
		log.Println("error when try to get auth by email with error", err)
		return
	}

	if auth.Role != "Merchant" {
		log.Println("auth:", auth)
		return helper.ErrInvalidRole
	}

	err = s.merchantRepo.UpdateMerchant(ctx, req)
	if err != nil {
		log.Println("error when try to update merchant with error", err)
		return
	}

	return
}
