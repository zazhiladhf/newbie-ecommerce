package category

import "github.com/zazhiladhf/newbie-ecommerce/pkg/helper"

type Category struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

func NewCategory() Category {
	return Category{}
}

func (c Category) Validate(req CreateCategoryRequest) (Category, error) {
	if req.Name == "" {
		return c, helper.ErrEmptyCategoryName
	}

	c.Name = req.Name

	return c, nil
}
