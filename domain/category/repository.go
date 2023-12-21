package category

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type PostgreSqlxRepository struct {
	db *sqlx.DB
}

func NewPostgreSqlxRepository(db *sqlx.DB) PostgreSqlxRepository {
	return PostgreSqlxRepository{
		db: db,
	}
}

func (r PostgreSqlxRepository) FindAllCategory(ctx context.Context) (categories []Category, err error) {
	query := `
    	SELECT 
			id, category_name 
    	FROM categories
    `

	err = r.db.SelectContext(ctx, &categories, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrCategoriesNotFound
		}
		return
	}

	if len(categories) == 0 {
		return nil, ErrCategoriesNotFound
	}

	return
}
