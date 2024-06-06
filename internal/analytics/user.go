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

func Login(login models.Login) error {
	if dbPool == nil {
		return errors.New("database pool is not initialized")
	}

	var userExists bool
	err := dbPool.QueryRow(context.Background(), "SELECT EXISTS (SELECT 1 FROM users WHERE email = $1 AND password = $2)", login.Email, login.Password).Scan(&userExists)
	if err != nil {
		return err
	}

	if !userExists {
		return errors.New("user does not exist or invalid credentials")
	}

	return nil
}
