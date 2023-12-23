package category

import (
	"context"

	"github.com/zazhiladhf/newbie-ecommerce/pkg/helper"
)

type repository interface {
	FindAllCategory(ctx context.Context) (categories []Category, err error)
	Create(ctx context.Context, category Category) (err error)
}

type CategoryService struct {
	repo repository
}

func newService(repo repository) CategoryService {
	return CategoryService{
		repo: repo,
	}
}

func (s CategoryService) GetListCategories(ctx context.Context) (list []Category, err error) {
	listCategories, err := s.repo.FindAllCategory(ctx)
	if err != nil {
		if err == helper.ErrCategoriesNotFound {
			return []Category{}, nil
		}
		return nil, err
	}
	return listCategories, nil
}

func (s CategoryService) CreateCategory(ctx context.Context, req Category) (err error) {
	err = s.repo.Create(ctx, req)
	if err != nil {
		return
	}

	return
}
