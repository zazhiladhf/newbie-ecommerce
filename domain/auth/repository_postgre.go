package auth

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

func (r PostgreSqlxRepository) StoreAuth(ctx context.Context, auth Auth) (err error) {
	query := `
		INSERT INTO auths (
			email, password, role
		) VALUES (
			:email, :password, :role
		)
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, auth)
	if err != nil {
		return
	}

	return
}

func (r PostgreSqlxRepository) GetAuthByEmail(ctx context.Context, email string) (auth Auth, err error) {
	query := `
		SELECT id, email, password, role 
		FROM auths 
		WHERE email = $1
	`

	err = r.db.GetContext(ctx, &auth, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return Auth{}, nil
		}
	}

	return
}
