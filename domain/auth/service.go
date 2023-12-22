package auth

import (
	"context"
	"log"
	"strconv"

	"github.com/zazhiladhf/newbie-ecommerce/config"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/jwt"
)

type postgreSqlxRepository interface {
	StoreAuth(ctx context.Context, auth Auth) (err error)
	GetAuthByEmail(ctx context.Context, email string) (auth Auth, err error)
	UpdateRoleAuth(ctx context.Context, id int) (err error)
}

type redisRepository interface {
	Set(ctx context.Context, email string, token string, lifeTime int) (err error)
	Get(ctx context.Context, email string) (token string, err error)
}

type AuthService struct {
	repo  postgreSqlxRepository
	redis redisRepository
}

func NewService(repo postgreSqlxRepository, redis redisRepository) AuthService {
	return AuthService{
		repo:  repo,
		redis: redis,
	}
}

func (s AuthService) RegisterAuth(ctx context.Context, req Auth) (err error) {
	_, err = s.repo.GetAuthByEmail(ctx, req.Email)
	if err != nil {
		log.Println("error when try to getAuthByEmail with error", err)
		return
	}

	err = req.EncryptPassword()
	if err != nil {
		log.Println("error when try to encrypt password with error", err)
		return
	}

	err = s.repo.StoreAuth(ctx, req)
	if err != nil {
		log.Println("error when try to store auth with error", err)
		return
	}

	return
}

func (s AuthService) Login(ctx context.Context, req Auth) (item Auth, token string, err error) {
	itemAuth, err := s.repo.GetAuthByEmail(ctx, req.Email)
	if err != nil {
		log.Println("error when try to getAuthByEmail with error", err)
		return
	}

	if itemAuth.Email != req.Email {
		return item, token, ErrInvalidEmail
	}

	ok, err := itemAuth.ValidatePassword(req.Password)
	if err != nil {
		log.Println("error when try to validate password with error", err)
		return req, token, err
	}

	if !ok {
		log.Println("error when try to !ok with error", err)
		return req, token, err
	}

	idString := strconv.Itoa(itemAuth.Id)
	token, err = jwt.GenerateToken(idString, itemAuth.Email, itemAuth.Role)
	if err != nil {
		log.Println("error when trying to generate token with error:", err)
	}

	err = s.redis.Set(ctx, itemAuth.Email, token, config.Cfg.Redis.LifeTime)
	if err != nil {
		log.Println("error when try to set data to redis with message :", err)
	}

	return itemAuth, token, err

}

func (s AuthService) UpdateRoleToMerchant(ctx context.Context, email string) (err error) {
	auth, err := s.repo.GetAuthByEmail(ctx, email)
	if err != nil {
		log.Println("error when try to getAuthByEmail with error", err)
		return
	}

	if auth.Role == "Merchant" {
		return ErrUserAlreadyMerchant
	}

	err = s.repo.UpdateRoleAuth(ctx, auth.Id)
	if err != nil {
		log.Println("error when try to update role auth with error", err)
		return
	}

	return
}
