package analytics

import (
	"context"
	"errors"

	"github.com/ICOMP-UNC/newworld-FlorSchroder/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

var dbPool *pgxpool.Pool

// SetDB sets the database pool for analytics package
func SetDB(pool *pgxpool.Pool) {
	dbPool = pool
}

// AddUser adds a new user to the database
func AddUser(register models.Register) error {
	if dbPool == nil {
		return errors.New("database pool is not initialized")
	}

	_, err := dbPool.Exec(context.Background(), "INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", register.Username, register.Email, register.Password)
	if err != nil {
		return err
	}

	return nil
}
