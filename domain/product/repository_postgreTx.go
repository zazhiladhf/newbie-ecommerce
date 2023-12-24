package product

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type txPostgresSQLXRepository struct {
	db *sqlx.Tx
}

func NewTXPostgresSQLXRepository(db *sqlx.Tx) *txPostgresSQLXRepository {
	return &txPostgresSQLXRepository{
		db: db,
	}
}

// GetProductByIdForUpdate implements ProductRepositoryTx.
func (r *txPostgresSQLXRepository) GetProductByIdForUpdate(ctx context.Context, id int) (product Product, err error) {
	query := `
		SELECT p.id, p.sku, p.name, p.description, p.price, p.stock, c.name as category, p.category_id, p.image_url, p.created_at, p.updated_at
		FROM products as p
		JOIN categories as c ON c.id = p.category_id
		WHERE p.id = $1
		FOR UPDATE
	`
	err = r.db.GetContext(ctx, &product, query, id)
	if err != nil {
		return
	}

	return
}

// UpdateProductStok implements ProductRepositoryTx.
func (r *txPostgresSQLXRepository) UpdateProductStok(ctx context.Context, product Product) (err error) {
	query := `
		UPDATE products 
		SET 
			stock = :stock, updated_at = NOW() 
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

var _ ProductRepositoryTx = (*txPostgresSQLXRepository)(nil)
