package tests

import (
	"context"
	"testing"

	"github.com/ICOMP-UNC/newworld-FlorSchroder/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

func TestDatabaseConnection(t *testing.T) {
	_, err := pgxpool.Connect(context.Background(), "postgres://florxha:mydb123@localhost:5432/florxha_tp3")
	if err != nil {
		t.Fatalf("failed to connect to the database: %v", err)
	}
}

func TestRoutesInitialization(t *testing.T) {
	app := fiber.New()
	dbPool, err := pgxpool.Connect(context.Background(), "postgres://florxha:mydb123@localhost:5432/florxha_tp3")
	if err != nil {
		t.Fatalf("failed to connect to the database: %v", err)
	}
	defer dbPool.Close()

	routes.InitRoutes(app, dbPool)
}

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
