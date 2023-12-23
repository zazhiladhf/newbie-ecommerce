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

func (r PostgresSQLXRepository) InsertProduct(ctx context.Context, product Product) (id int, err error) {
	query := `
		INSERT INTO products (
			name, description, price, stock, category_id, merchant_id, image_url, sku
		) VALUES (
			:name, :description, :price, :stock, :category_id, :merchant_id, :image_url, :sku
		)
		RETURNING id
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.GetContext(ctx, &id, product)
	if err != nil {
		return
	}

	return
}

func (r PostgresSQLXRepository) GetProducts(ctx context.Context, queryParam string, id int, limit int, page int) (list []Product, totalData int, err error) {
	queryGet := `
		SELECT 
			p.id, p.name, p.description, p.price, p.stock, c.name as category, p.image_url, p.sku
		FROM products as p
		JOIN categories as c
			ON c.id = p.category_id
		WHERE p.merchant_id = $1
	`

	queryCountByEmail := `
		SELECT COUNT(p.id) as total_data 
		FROM products as p
		JOIN categories as c 
			ON c.id = p.category_id
		WHERE p.merchant_id = $1
	`

	filter := mappingQueryFilter(queryParam)
	offset := (page - 1) * limit
	queryLimitOffset := fmt.Sprintf("LIMIT %d OFFSET %d", limit, offset)

	queryCount := fmt.Sprintf("%s %s %s", queryCountByEmail, filter, queryLimitOffset)
	query := fmt.Sprintf("%s %s %s", queryGet, filter, queryLimitOffset)

	err = r.db.SelectContext(ctx, &list, query, id)
	if err != nil {
		return
	}

	err = r.db.GetContext(ctx, &totalData, queryCount, id)
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

func (r PostgresSQLXRepository) GetProductById(ctx context.Context, id int) (product Product, err error) {
	query := `
	SELECT
		p.id,
		p.sku,
		p.name,
		p.description,
		p.price,
		p.stock,
		c.name as category,
		p.category_id,
		p.image_url,
		p.created_at,
		p.updated_at
	FROM products as p
	JOIN categories as c ON c.id = p.category_id
	WHERE p.id = $1
	`
	err = r.db.GetContext(ctx, &product, query, id)
	if err != nil {
		return
	}

	return
}

func (r PostgresSQLXRepository) UpdateProduct(ctx context.Context, product Product) (err error) {
	query := `
	UPDATE products SET 
		name = :name, 
		description = :description, 
		price = :price, 
		stock = :stock, 
		category_id = :category_id, 
		image_url = :image_url, 
		updated_at = NOW() 
	WHERE id = :id
	`
	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, product)
	if err != nil {
		return
	}

	return
}

func (r PostgresSQLXRepository) GetProductBySku(ctx context.Context, sku string) (product Product, err error) {
	query := `
	SELECT p.id, p.sku, p.name, p.description, p.price, p.stock, c.name as category, p.category_id, p.merchant_id, p.image_url, p.created_at, p.updated_at, m.name as merchant_name, m.city as merchant_city
	FROM products as p
	JOIN categories as c 
		ON c.id = p.category_id
	JOIN merchants as m 
		ON m.id = p.merchant_id
	WHERE p.sku = $1
`

	err = r.db.GetContext(ctx, &product, query, sku)
	if err != nil {
		return
	}

	return
}
