package user

import (
	"strings"
	"time"

	"github.com/zazhiladhf/newbie-ecommerce/domain/auth"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/helper"
)

// Gender type represents the gender of a user.
type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

type User struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	DateOfBirth string `db:"date_of_birth"`
	PhoneNumber string `db:"phone_number"`
	Gender      Gender `db:"gender"`
	Address     string `db:"address"`
	ImageUrl    string `db:"image_url"`
	AuthId      int    `db:"auth_id"`
	auth.Auth
}

func NewUser() User {
	return User{}
}

func (u User) newFromRequest(req RequestBodyCreateProfileUser) (user User, err error) {
	user = User{
		Name:        req.Name,
		DateOfBirth: req.DateOfBirth,
		PhoneNumber: req.PhoneNumber,
		Gender:      Gender(req.Gender),
		Address:     req.Address,
		ImageUrl:    req.ImageUrl,
		AuthId:      user.Auth.Id,
	}

	err = user.ValidateRequestCreateUser()
	return
}

func (u User) ValidateRequestCreateUser() (err error) {
	if u.Gender == "" {
		return helper.ErrGenderEmpty
	}

	if !isValidGender(string(u.Gender)) {
		return helper.ErrInvalidGender
	}

	if u.PhoneNumber == "" {
		return helper.ErrEmptyPhoneNumber
	}

	if len(u.PhoneNumber) < 10 {
		return helper.ErrInvalidPhoneNumber
	}

	if u.Name == "" {
		return helper.ErrEmptyNameUser
	}

	if u.Address == "" {
		return helper.ErrEmptyAddress
	}

	if u.DateOfBirth == "" {
		return helper.ErrEmptyDateOfBirth
	}

	if !isValidDateFormat(u.DateOfBirth) {
		return helper.ErrInvalidDateOfBirth
	}

	if u.ImageUrl == "" {
		return helper.ErrEmptyImageURLUser
	}

	if err != nil {
		return helper.ErrUserNotFound
	}

	return
}

func isValidDateFormat(dateString string) bool {
	_, err := time.Parse("2006-01-02", dateString)
	return err == nil
}

func isValidGender(gender string) bool {
	genderLower := strings.ToLower(gender)
	return genderLower == "male" || genderLower == "female"
}

func (u User) UserResponse(user User) GetUserResponse {
	resp := GetUserResponse{
		Name:        user.Name,
		DateOfBirth: user.DateOfBirth,
		PhoneNumber: user.PhoneNumber,
		Gender:      string(user.Gender),
		Address:     user.Address,
		ImageUrl:    user.ImageUrl,
	}

	return resp
}
