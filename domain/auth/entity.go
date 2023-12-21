package auth

import (
	"errors"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailEmpty      = errors.New("email required")
	ErrInvalidEmail    = errors.New("invalid email")
	ErrPasswordEmpty   = errors.New("password required")
	ErrInvalidPassword = errors.New("invalid password")
	ErrDuplicateEmail  = errors.New("email already used")
	ErrRepository      = errors.New("error repository")
	ErrInternalServer  = errors.New("unknown error")
)

type Auth struct {
	Id       int    `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Role     string `db:"role"`
}

func NewAuth() Auth {
	return Auth{}
}

func (a Auth) ValidateFormRegister(req registerRequest) (Auth, error) {
	if req.Email == "" {
		return a, ErrEmailEmpty
	}

	if !valid(req.Email) {
		return a, ErrInvalidEmail
	}

	if req.Password == "" {
		return a, ErrPasswordEmpty
	}

	if len(req.Password) < 6 {
		return a, ErrInvalidPassword
	}

	a.Email = req.Email
	a.Password = req.Password
	a.Role = "merchant"
	return a, nil
}

func (a Auth) ValidateFormLogin(req loginRequest) (Auth, error) {
	if req.Email == "" {
		return a, ErrEmailEmpty
	}

	if !valid(req.Email) {
		return a, ErrInvalidEmail
	}

	if req.Password == "" {
		return a, ErrPasswordEmpty
	}

	if len(req.Password) < 6 {
		return a, ErrInvalidPassword
	}

	a.Email = req.Email
	a.Password = req.Password
	return a, nil
}

func (a *Auth) EncryptPassword() (err error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	a.Password = string(encrypted)
	return
}

func (a Auth) ValidatePassword(password string) (ok bool, err error) {
	err = bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password))
	if err != nil {
		return ok, err
	}
	ok = true
	return
}

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
