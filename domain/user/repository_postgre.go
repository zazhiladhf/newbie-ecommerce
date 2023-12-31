package user

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

// GetUserByAuthId implements UserRepository.
func (r PostgreSqlxRepository) GetUserByAuthId(ctx context.Context, authId int) (user User, err error) {
	query := `
		SELECT u.id, u.name, u.date_of_birth, u.phone_number, u.gender, u.address, u.image_url,
			u.auth_id, a.id AS auth_id, a.email, a.password, a.role
		FROM users u
		JOIN auths a ON u.auth_id = a.id
		WHERE u.auth_id = $1
	`

	err = r.db.GetContext(ctx, &user, query, authId)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, nil
		}
	}

	return
}

func (r PostgreSqlxRepository) InsertUser(ctx context.Context, user User) (err error) {
	query := `
		INSERT INTO users (
			name, auth_id, date_of_birth, phone_number, gender, address, image_url
		) VALUES (
			:name, :auth_id, :date_of_birth, :phone_number, :gender, :address, :image_url
		)
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, user)
	if err != nil {
		return
	}

	return
}

func (r PostgreSqlxRepository) GetUserById(ctx context.Context, id int) (user User, err error) {
	query := `
		SELECT 
			id, name, date_of_birth, phone_number, gender, address, image_url, auth_id
		FROM users
		WHERE id = $1
	`

	err = r.db.GetContext(ctx, &user, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, nil
		}
	}

	return
}

func (r PostgreSqlxRepository) UpdateUser(ctx context.Context, req User) (err error) {
	query := `
		UPDATE users
		SET name = :name, date_of_birth = :date_of_birth, phone_number = :phone_number,
			gender = :gender, address = :address, image_url = :image_url
		WHERE id = :auth_id
	`

	_, err = r.db.NamedExecContext(ctx, query, req)
	if err != nil {
		return
	}

	return
}
