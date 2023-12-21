package product

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PostgresSQLXRepository struct {
	db *sqlx.DB
}

func NewPostgresSQLXRepository(db *sqlx.DB) PostgresSQLXRepository {
	return PostgresSQLXRepository{
		db: db,
	}
}

func (r PostgresSQLXRepository) InsertProduct(ctx context.Context, model Product) (id int, err error) {
	query := `
		INSERT INTO products (
			name, image_url, stock, price, category_id, email_auth
		) VALUES (
			:name, :image_url, :stock, :price, :category_id, :email_auth
		)
		RETURNING id
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.GetContext(ctx, &id, model)
	if err != nil {
		return
	}

	return
}

// func (r PostgresSQLXRepository) FindAllProducts(ctx context.Context) (list []Product, err error) {
// 	query := `
// 		SELECT
// 			id, name, image_url, stock, price, category_id, email_auth, category
// 		FROM products
// 	`

// 	err = r.db.SelectContext(ctx, &list, query)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, ErrNotFound
// 		}
// 		return
// 	}

// 	if len(list) == 0 {
// 		return nil, err
// 	}

// 	return list, nil
// }

func (r PostgresSQLXRepository) FindProductByEmail(ctx context.Context, queryParam string, email string, limit int, page int) (list []Product, totalData int, err error) {
	queryByEmail := `
		SELECT 
			p.id, p.name, p.image_url, p.stock, p.price, c.category_name as category, p.email_auth
		FROM products as p
		JOIN categories as c
			ON c.id = p.category_id
		WHERE email_auth = $1
	`

	queryCountByEmail := `
		SELECT COUNT(p.id) as total_data 
		FROM products as p
		JOIN categories as c 
			ON c.id = p.category_id
		WHERE p.email_auth = $1
	`

	filter := mappingQueryFilter(queryParam)
	offset := (page - 1) * limit
	queryLimitOffset := fmt.Sprintf("LIMIT %d OFFSET %d", limit, offset)

	queryCount := fmt.Sprintf("%s %s %s", queryCountByEmail, filter, queryLimitOffset)
	query := fmt.Sprintf("%s %s %s", queryByEmail, filter, queryLimitOffset)

	err = r.db.SelectContext(ctx, &list, query, email)
	if err != nil {
		return
	}

	err = r.db.GetContext(ctx, &totalData, queryCount, email)
	if err != nil {
		if err == sql.ErrNoRows {
			totalData = 0
			return []Product{}, totalData, nil
		}
		return
	}

	return
}

func mappingQueryFilter(queryParam string) string {
	filter := ""

	if queryParam != "" {
		filter = fmt.Sprintf("%s AND name ILIKE '%%%s%%'", filter, queryParam)
	}

	return filter
}
