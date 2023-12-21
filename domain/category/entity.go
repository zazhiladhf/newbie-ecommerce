package category

import "errors"

var (
	ErrRepository         = errors.New("error repository")
	ErrInternalServer     = errors.New("unknown error")
	ErrCategoriesNotFound = errors.New("categories not found")
)

type Category struct {
	Id           int    `db:"id"`
	CategoryName string `db:"category_name"`
}
