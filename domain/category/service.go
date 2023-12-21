package category

import (
	"context"
)

type repository interface {
	FindAllCategory(ctx context.Context) (categories []Category, err error)
}

type CategoryService struct {
	repo repository
}

func newService(repo repository) CategoryService {
	return CategoryService{
		repo: repo,
	}
}

func (s CategoryService) getListCategories(ctx context.Context) (list []Category, err error) {
	listCategories, err := s.repo.FindAllCategory(ctx)
	if err != nil {
		if err == ErrCategoriesNotFound {
			return []Category{}, nil
		}
		return nil, err
	}
	return listCategories, nil
}
