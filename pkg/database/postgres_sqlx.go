package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/zazhiladhf/newbie-ecommerce/config"
)

func ConnectPostgresSqlx(cfg config.DB) (db *sqlx.DB, err error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name,
	)

	db, err = sqlx.Open("postgres", dsn)
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}

	if db == nil {
		log.Println("error when to try connect db with error :", err.Error())
		panic("db not connected")
	}

	log.Println("database connect success 🚀🚀🚀")
	log.Println("dsn :", dsn)

	return
}

func Migrate(db *sqlx.DB) (err error) {
	query := `
		CREATE TABLE IF NOT EXISTS auths (
			id SERIAL PRIMARY KEY,
			email varchar(100) NOT NULL,
			password varchar(100) NOT NULL,
			role varchar(100) NOT NULL,
			UNIQUE (email)
		);

		CREATE TABLE IF NOT EXISTS categories (
			id SERIAL PRIMARY KEY,
			name varchar(100) NOT NULL,
			UNIQUE (name)
		);
    
		CREATE TABLE IF NOT EXISTS merchants (
			id SERIAL PRIMARY KEY,
			auth_id int NOT NULL,
			name VARCHAR(100) NOT NULL,
			phone_number VARCHAR(20) NOT NULL,
			address VARCHAR(100) NOT NULL,
			image_url VARCHAR(100) NOT NULL,
			city varchar(100) NOT NULL,
			FOREIGN KEY ("auth_id") REFERENCES "auths" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
			UNIQUE (auth_id)
		);

		CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			name varchar(100) NOT NULL,
			description varchar(255) NOT NULL,
			stock int NOT NULL,
			price int NOT NULL,
			category_id int NOT NULL,
			image_url varchar(100) NOT NULL,
			merchant_id int NOT NULL,
			sku varchar(100) NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP WITH TIME ZONE NULL DEFAULT NOW(),
			FOREIGN KEY ("category_id") REFERENCES "categories" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
			FOREIGN KEY ("merchant_id") REFERENCES "merchants" ("id") ON DELETE CASCADE ON UPDATE CASCADE
		);

		CREATE TYPE genders AS ENUM ('male', 'female');

		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			auth_id int NOT NULL,
			name VARCHAR(100) NOT NULL,
			date_of_birth DATE NOT NULL,
			phone_number VARCHAR(20) NOT NULL,
			gender genders NOT NULL,
			address VARCHAR(100) NOT NULL,
			image_url VARCHAR(100) NOT NULL,
			FOREIGN KEY ("auth_id") REFERENCES "auths" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
			UNIQUE (auth_id)
		);
	`
	_, err = db.Exec(query)

	return
}
