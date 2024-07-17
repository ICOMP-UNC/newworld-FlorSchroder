package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/ICOMP-UNC/newworld-FlorSchroder/internal/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var dbPool *pgxpool.Pool

// SetDB sets the database pool for analytics package
func SetDB(pool *pgxpool.Pool) {
	dbPool = pool
}

func AddUser(register models.Register) error {
	if dbPool == nil {
		return errors.New("database pool is not initialized")
	}

	// Determine the role based on the username and password
	role := "normal"
	if register.Username == "ubuntu" && register.Password == "ubuntu" {
		role = "admin"
	}

	// GenerateJWT(register.Username)
	jwt, err0 := GenerateJWT(register.Username, role)
	if err0 != nil {
		return err0
	}

	_, err := dbPool.Exec(context.Background(), "INSERT INTO users (username, email, password, jwt, role) VALUES ($1, $2, $3, $4, $5)", register.Username, register.Email, register.Password, jwt, role)
	if err != nil {
		return err
	}

	return nil
}

func Login(login models.Login) (string, error) {
	if dbPool == nil {
		return "", errors.New("database pool is not initialized")
	}

	var jwtKey string
	err := dbPool.QueryRow(context.Background(), "SELECT jwt FROM users WHERE email = $1 AND password = $2", login.Email, login.Password).Scan(&jwtKey)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("user does not exist or invalid credentials")
		}
		return "", err
	}

	return jwtKey, nil
}

func GenerateJWT(username string, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func AddOffer(offer models.Offer) error {
	if dbPool == nil {
		return errors.New("database pool is not initialized")
	}

	if offer.Quantity <= 0 || offer.Price <= 0 {
		return errors.New("quantity and price must be greater than 0")
	}

	validCategories := []string{"medicine", "food", "ammo"}
	if !contains(validCategories, offer.Category) {
		return errors.New("category must be either 'medicine', 'food', or 'ammo'")
	}

	_, err := dbPool.Exec(context.Background(), "INSERT INTO offers (name, quantity, price, category) VALUES ($1, $2, $3, $4)", offer.Name, offer.Quantity, offer.Price, offer.Category)
	if err != nil {
		return err
	}

	return nil
}

func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

func GetOffers(jwtString string) ([]models.OfferWithID, error) {
	if dbPool == nil {
		return nil, errors.New("database pool is not initialized")
	}

	// Check the JWT using the checkJWT function
	_, err := checkJWT(jwtString)
	if err != nil {
		return nil, err
	}

	rows, err := dbPool.Query(context.Background(), "SELECT id, name, quantity, price, category FROM offers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	offers := []models.OfferWithID{}
	for rows.Next() {
		var offer models.OfferWithID
		err := rows.Scan(&offer.ID, &offer.Name, &offer.Quantity, &offer.Price, &offer.Category)
		if err != nil {
			return nil, err
		}
		offers = append(offers, offer)
	}

	return offers, nil
}

func Checkout(order models.Order, jwt string) (models.Message, error) {
	fmt.Println("Starting Checkout")

	if dbPool == nil {
		fmt.Println("Error: database pool is not initialized")
		return models.Message{}, errors.New("database pool is not initialized")
	}

	// check valid JWT
	claims, err := checkJWT(jwt)
	if err != nil {
		fmt.Println("Error:", err)
		return models.Message{}, err
	}

	// print username in console
	fmt.Println("Username:", claims["username"], "is checking out")

	var total float64
	for _, item := range order.Items {
		var price float64
		var quantity int
		err := dbPool.QueryRow(context.Background(), "SELECT price, quantity FROM offers WHERE id = $1", item.ProductID).Scan(&price, &quantity)
		if err != nil {
			fmt.Println("Error:", err)
			return models.Message{}, err
		}

		fmt.Println("Price:", price, "Quantity:", quantity)

		if item.Quantity > quantity {
			fmt.Println("Error: not enough quantity available")
			return models.Message{}, errors.New("not enough quantity available")
		}

		total += price * float64(item.Quantity)

		newQuantity := quantity - item.Quantity
		if newQuantity == 0 {
			_, err = dbPool.Exec(context.Background(), "DELETE FROM offers WHERE id = $1", item.ProductID)
			if err != nil {
				fmt.Println("Error:", err)
				return models.Message{}, err
			}
		} else {
			_, err = dbPool.Exec(context.Background(), "UPDATE offers SET quantity = $1 WHERE id = $2", newQuantity, item.ProductID)
			if err != nil {
				fmt.Println("Error:", err)
				return models.Message{}, err
			}
		}
	}

	fmt.Println("Total:", total)

	_, err = dbPool.Exec(context.Background(), "INSERT INTO orders (items, status) VALUES ($1, $2)", order.Items, "preparing")
	if err != nil {
		fmt.Println("Error:", err)
		return models.Message{}, err
	}

	fmt.Println("Order placed")

	return models.Message{Total: total, Status: "preparing"}, nil
}

func checkJWT(jwtString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("your_secret_key"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Si el token es válido, devuelve las claims
		return claims, nil
	} else {
		// Si el token no es válido, devuelve un error
		return nil, errors.New("invalid token")
	}
}

func checkJWTRole(jwtString string) error {
	mapClaims, err := checkJWT(jwtString)
	if err != nil {
		return err
	}

	role, ok := mapClaims["role"].(string)
	if !ok {
		return errors.New("invalid JWT role claim")
	}

	if role != "admin" {
		return errors.New("user is not an admin")
	}

	return nil
}

func GetDashboard(jwt string) (models.Dashboard, error) {
	// this funcition must return all de orders from the table orders of the database if the user is admin
	// if the user is not admin, it must return an error
	if dbPool == nil {
		return models.Dashboard{}, errors.New("database pool is not initialized")
	}

	// check valid JWT
	err := checkJWTRole(jwt)
	if err != nil {
		return models.Dashboard{}, err
	}

	rows, err := dbPool.Query(context.Background(), "SELECT items FROM orders")
	if err != nil {
		return models.Dashboard{}, err
	}
	defer rows.Close()

	dashboard := models.Dashboard{}
	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.Items)
		if err != nil {
			return models.Dashboard{}, err
		}
		dashboard.Orders = append(dashboard.Orders, order)
	}

	return dashboard, nil
}

func GetOrderStatus(id int, jwt string) (string, error) {
	if dbPool == nil {
		return "", errors.New("database pool is not initialized")
	}

	// check valid JWT and role
	err := checkJWTRole(jwt)
	if err != nil {
		return "", err
	}

	var status string
	err = dbPool.QueryRow(context.Background(), "SELECT status FROM orders WHERE id = $1", id).Scan(&status)
	if err != nil {
		return "", err
	}

	return status, nil
}

func UpdateOrderStatus(id int, status string, jwt string) error {
	if dbPool == nil {
		return errors.New("database pool is not initialized")
	}

	// check valid JWT and role
	err := checkJWTRole(jwt)
	if err != nil {
		return err
	}

	_, err = dbPool.Exec(context.Background(), "UPDATE orders SET status = $1 WHERE id = $2", status, id)
	if err != nil {
		return err
	}

	return nil
}
