package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

// ConnectDB establishes a connection to the PostgreSQL database.
func ConnectDB() (*pgxpool.Pool, error) {
	databaseUrl := "postgres://florxha:mydb123@localhost:5432/florxha_tp3"

	config, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		return nil, fmt.Errorf("unable to parse database url: %v", err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %v", err)
	}

	return pool, nil
}

// InitDB initializes the database with necessary tables.
func InitDB(pool *pgxpool.Pool) error {
	_, err := pool.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		role TEXT NOT NULL,
		jwt TEXT
	);

	CREATE TABLE IF NOT EXISTS offers (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		quantity INT NOT NULL,
		price FLOAT NOT NULL,
		category TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		items JSONB NOT NULL,
		status TEXT NOT NULL
	);
	`)
	if err != nil {
		return fmt.Errorf("failed to initialize database: %v", err)
	}
	return nil
}
