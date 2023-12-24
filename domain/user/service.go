package user

import (
	"context"
	"log"

	"github.com/zazhiladhf/newbie-ecommerce/domain/auth"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/helper"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user User) (err error)
	GetUserById(ctx context.Context, id int) (user User, err error)
	UpdateUser(ctx context.Context, req User) (err error)
	GetUserByAuthId(ctx context.Context, authId int) (user User, err error)
}

type authRepository interface {
	GetAuthByEmail(ctx context.Context, email string) (auth auth.Auth, err error)
}

type UserService struct {
	userRepo UserRepository
	authRepo authRepository
}

func NewService(userRepo UserRepository, authrepo authRepository) UserService {
	return UserService{
		userRepo: userRepo,
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

	err = s.userRepo.InsertUser(ctx, req)
	if err != nil {
		log.Println("error when try to insert user to db with error", err)
		return
	}

	return
}

func (s UserService) GetUserById(ctx context.Context, id int) (resp GetUserResponse, err error) {
	user, err := s.userRepo.GetUserById(ctx, id)
	if err != nil {
		log.Println("error when try to get user by id with error", err)
		return
	}

	if user.Id == 0 {
		return resp, helper.ErrUserNotFound
	}

	resp = NewUser().UserResponse(user)

	return
}

func (s UserService) UpdateProfileUser(ctx context.Context, req User, email string) (err error) {
	auth, err := s.authRepo.GetAuthByEmail(ctx, email)
	if err != nil {
		log.Println("error when try to get auth by email with error", err)
		return
	}

	if auth.Role != "User" {
		log.Println("auth:", auth)
		return helper.ErrInvalidRole
	}

	// user, err := s.repo.GetUserById(ctx, req.AuthId)
	// if err != nil {
	// 	log.Println("error when try to get user by id with error", err)
	// 	return
	// }

	// if user.Id == 0 {
	// 	return helper.ErrUserNotFound
	// }

	err = s.userRepo.UpdateUser(ctx, req)
	if err != nil {
		log.Println("error when try to update user with error", err)
		return
	}

	return
}
