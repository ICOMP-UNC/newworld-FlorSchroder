package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ICOMP-UNC/newworld-FlorSchroder/internal/models"
	"github.com/ICOMP-UNC/newworld-FlorSchroder/internal/services"
	"github.com/golang-jwt/jwt"
	"github.com/jackc/pgx/v4/pgxpool"
)

// func TestDatabaseConnection(t *testing.T) {
// 	_, err := pgxpool.Connect(context.Background(), "postgres://florxha:mydb123@localhost:5432/florxha_tp3")
// 	if err != nil {
// 		t.Fatalf("failed to connect to the database: %v", err)
// 	}
// }

// func TestRoutesInitialization(t *testing.T) {
// 	app := fiber.New()
// 	dbPool, err := pgxpool.Connect(context.Background(), "postgres://florxha:mydb123@localhost:5432/florxha_tp3")
// 	if err != nil {
// 		t.Fatalf("failed to connect to the database: %v", err)
// 	}
// 	defer dbPool.Close()

// 	routes.InitRoutes(app, dbPool)
// }

// func TestRegisterNewUser(t *testing.T) {
// 	//initialize the database
// 	dbPool, err := pgxpool.Connect(context.Background(), "postgres://florxha:mydb123@localhost:5432/florxha_tp3")
// 	if err != nil {
// 		t.Fatalf("failed to connect to the database: %v", err)
// 	}
// 	defer dbPool.Close()

// 	register := models.Register{
// 		Username: "testuser",
// 		Email:    "test@test.com",
// 		Password: "testpassword",
// 	}

// 	err1 := services.AddUser(register)
// 	if err1 != nil {
// 		t.Fatalf("failed to add user: %v", err1)
// 	}

// }

// func TestRegisterNewUser(t *testing.T) {
// 	app := fiber.New()
// 	dbPool, err := pgxpool.Connect(context.Background(), "postgres://florxha:mydb123@localhost:5432/florxha_tp3")
// 	if err != nil {
// 		t.Fatalf("failed to connect to the database: %v", err)
// 	}
// 	defer dbPool.Close()

// 	routes.InitRoutes(app, dbPool)

// 	register := models.Register{
// 		Username: "testuser",
// 		Email:    "test@test.com",
// 		Password: "testpassword",
// 	}
// 	body, err := json.Marshal(register)
// 	if err != nil {
// 		t.Fatalf("failed to marshal register: %v", err)
// 	}

// 	req := httptest.NewRequest(http.MethodPost, "/auth/register", bytes.NewReader(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, err := app.Test(req)
// 	if err != nil {
// 		t.Fatalf("failed to make request: %v", err)
// 	}
// 	if resp.StatusCode != http.StatusCreated {
// 		t.Fatalf("expected status code 201, got %d", resp.StatusCode)
// 	}
// }

var testEmail string

func TestAddUser(t *testing.T) {
	dbPool, err := pgxpool.Connect(context.Background(), "postgres://florxha:mydb123@localhost:5432/florxha_tp3")
	if err != nil {
		t.Fatalf("failed to connect to the database: %v", err)
	}
	defer dbPool.Close()

	// Set the global database connection to the one you just created
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
	dbPool, err := pgxpool.Connect(context.Background(), "postgres://florxha:mydb123@localhost:5432/florxha_tp3")
	if err != nil {
		t.Fatalf("failed to connect to the database: %v", err)
	}
	defer dbPool.Close()

	// Set the global database connection to the one you just created
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

	// Generate a JWT
	tokenString, err := services.GenerateJWT(username, role)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Parse the JWT
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

	// Generate a JWT
	tokenString, err := services.GenerateJWT(username, role)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Check the JWT
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

	// Generate a JWT
	tokenString, err := services.GenerateJWT(username, role)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Check the JWT role
	err = services.CheckJWTRole(tokenString)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
