package user

import (
	"context"
	"log"

	"github.com/zazhiladhf/newbie-ecommerce/domain/auth"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/helper"
)

type postgreSqlxRepository interface {
	InsertUser(ctx context.Context, user User) (err error)
	GetUserById(ctx context.Context, id int) (user User, err error)
}

type authRepository interface {
	GetAuthByEmail(ctx context.Context, email string) (auth auth.Auth, err error)
}

type UserService struct {
	repo     postgreSqlxRepository
	authRepo authRepository
}

func NewService(repo postgreSqlxRepository, authrepo authRepository) UserService {
	return UserService{
		repo:     repo,
		authRepo: authrepo,
	}
}

func (s UserService) CreateProfileUser(ctx context.Context, req User, email string) (user User, err error) {
	auth, err := s.authRepo.GetAuthByEmail(ctx, email)
	if err != nil {
		log.Println("error when try to get auth by email with error", err)
		return
	}

	if auth.Role != "User" {
		log.Println("auth:", auth)
		return user, helper.ErrInvalidRole
	}

	// user, err = s.repo.GetUserById(ctx, req.Id)
	// if err != nil {
	// 	log.Println("error when try to get user by id with error", err)
	// 	return
	// }

	// if user.Auth.Role == "merchant" {
	// 	return user, ErrUnauthorized
	// }

	err = s.repo.InsertUser(ctx, req)
	if err != nil {
		log.Println("error when try to insert user to db with error", err)
		return
	}

	return
}

// func (s UserService) Login(ctx context.Context, req Auth) (item Auth, token string, err error) {
// 	itemAuth, err := s.repo.GetAuthByEmail(ctx, req.Email)
// 	if err != nil {
// 		log.Println("error when try to getAuthByEmail with error", err)
// 		return
// 	}

// 	if itemAuth.Email != req.Email {
// 		return item, token, ErrInvalidEmail
// 	}

// 	ok, err := itemAuth.ValidatePassword(req.Password)
// 	if err != nil {
// 		log.Println("error when try to validate password with error", err)
// 		return req, token, err
// 	}

// 	if !ok {
// 		log.Println("error when try to !ok with error", err)
// 		return req, token, err
// 	}

// 	token, err = jwt.GenerateToken(itemAuth.Email)
// 	if err != nil {
// 		log.Println("error when trying to generate token with error:", err)
// 	}

// 	err = s.redis.Set(ctx, itemAuth.Email, token, config.Cfg.Redis.LifeTime)
// 	if err != nil {
// 		log.Println("error when try to set data to redis with message :", err)
// 	}

// 	return itemAuth, token, err

// }
