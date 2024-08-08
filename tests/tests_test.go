package tests

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/ICOMP-UNC/newworld-FlorSchroder/internal/models"
	"github.com/ICOMP-UNC/newworld-FlorSchroder/internal/services"
	"github.com/golang-jwt/jwt"
	"github.com/jackc/pgx/v4/pgxpool"
)

var testEmail string

func getDatabaseURL() string {
	// Try to get the DATABASE_URL from the environment, fall back to localhost if not set
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://florxha:mydb123@localhost:5432/florxha_tp3"
	}

	// create users table
	dbPool, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer dbPool.Close()

	_, err = dbPool.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS users (username TEXT, email TEXT, password TEXT, jwt TEXT, role TEXT)")
	if err != nil {
		log.Fatalf("failed to create users table: %v", err)
	}

	return dbURL
}

func TestAddUser(t *testing.T) {
	dbPool, err := pgxpool.Connect(context.Background(), getDatabaseURL())
	if err != nil {
		t.Fatalf("failed to connect to the database: %v", err)
	}
	defer dbPool.Close()

	services.SetDB(dbPool)

	testEmail = fmt.Sprintf("test%d@test.com", time.Now().Unix())

	user := models.Register{
		Username: "testuser",
		Email:    testEmail,
		Password: "testpassword",
	}

	err = services.AddUser(user)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestLogin(t *testing.T) {
	dbPool, err := pgxpool.Connect(context.Background(), getDatabaseURL())
	if err != nil {
		t.Fatalf("failed to connect to the database: %v", err)
	}
	defer dbPool.Close()

	services.SetDB(dbPool)

	user := models.Login{
		Email:    testEmail,
		Password: "testpassword",
	}

	token, err := services.Login(user)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if token == "" {
		t.Fatalf("expected a token, got an empty string")
	}

}

func TestGenerateJWT(t *testing.T) {
	username := "testuserjwt"
	role := "testrole"

	tokenString, err := services.GenerateJWT(username, role)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("your_secret_key"), nil
	})

	if err != nil {
		t.Fatalf("error parsing token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["username"] != username || claims["role"] != role {
			t.Fatalf("unexpected claims in token: %v", claims)
		}
	} else {
		t.Fatalf("invalid token")
	}
}

func TestCheckJWT(t *testing.T) {
	username := "testuserjwt"
	role := "testrole"

	tokenString, err := services.GenerateJWT(username, role)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	claims, err := services.CheckJWT(tokenString)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if claims["username"] != username || claims["role"] != role {
		t.Fatalf("unexpected claims in token: %v", claims)
	}
}

func TestCheckJWTRole(t *testing.T) {
	username := "testuserjwt"
	role := "admin"

	tokenString, err := services.GenerateJWT(username, role)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = services.CheckJWTRole(tokenString)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
