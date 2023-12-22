package merchant

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

func (r PostgreSqlxRepository) InsertMerchant(ctx context.Context, merchant Merchant) (err error) {
	query := `
		INSERT INTO merchants (
			name, auth_id, phone_number, address, image_url, city
		) VALUES (
			:name, :auth_id, :phone_number, :address, :image_url, :city
		)
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, merchant)
	if err != nil {
		return
	}

	return
}

func (r PostgreSqlxRepository) GetMerchantById(ctx context.Context, id int) (merchant Merchant, err error) {
	query := `
		SELECT 
			id, name, phone_number, address, image_url, city, auth_id
		FROM merchants
		WHERE id = $1
	`

	err = r.db.GetContext(ctx, &merchant, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return Merchant{}, nil
		}
	}

	return
}

func (r PostgreSqlxRepository) UpdateMerchant(ctx context.Context, req Merchant) (err error) {
	query := `
		UPDATE merchants
		SET name = :name, phone_number = :phone_number, address = :address, image_url = :image_url, city = :city
		WHERE id = :auth_id
	`

	_, err = r.db.NamedExecContext(ctx, query, req)
	if err != nil {
		return
	}

	return
}
