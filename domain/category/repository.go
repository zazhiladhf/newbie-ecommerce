package category

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/zazhiladhf/newbie-ecommerce/pkg/helper"
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
			id, name 
    	FROM categories
    `

	err = r.db.SelectContext(ctx, &categories, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.ErrCategoriesNotFound
		}
		return
	}

	if len(categories) == 0 {
		return nil, helper.ErrCategoriesNotFound
	}

	return
}

func (r PostgreSqlxRepository) Create(ctx context.Context, category Category) (err error) {
	query := `
	INSERT INTO categories (
			name 
		) VALUES (
			:name
		)
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, category)
	if err != nil {
		return
	}

	return
}
