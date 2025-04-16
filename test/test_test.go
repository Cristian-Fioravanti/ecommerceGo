package tests

import (
	"bytes"
	"encoding/json"
	// "fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	
	"github.com/cristian-fioravanti/ecommerceGo/models"
	"github.com/cristian-fioravanti/ecommerceGo/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SetupTestDB crea un database SQLite in-memory con utenti admin e normali
func SetupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to test database")
	}

	// Migra i modelli
	db.AutoMigrate(&models.User{}, &models.Product{})

	return db
}


// setupRouter imposta il router Gin per il test
func setupRouter(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	routes.UserRoutes(r, db) // Aggiungi tutte le rotte necessarie
	return r
}

// TestSignup verifica la registrazione di un nuovo utente
func TestSignup(t *testing.T) {
	db := SetupTestDB()
	router := setupRouter(db)

	payload := map[string]string{
		"first_name": "Cristian",
		"last_name":  "Fiss",
		"email":      "cristian.fiss@example.com",
		"password":   "test1234",
	}
	jsonPayload, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/users/signup", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

// TestLoginCristianFiss verifica il login di un utente normale
func TestLoginCristianFiss(t *testing.T) {
	db := SetupTestDB()
	router := setupRouter(db)
	payloadSign := map[string]string{
		"first_name": "Cristian",
		"last_name":  "Fiss",
		"email":      "cristian.fiss@example.com",
		"password":   "test1234",
	}
	jsonPayloadSign, _ := json.Marshal(payloadSign)

	reqSign, _ := http.NewRequest("POST", "/users/signup", bytes.NewBuffer(jsonPayloadSign))
	reqSign.Header.Set("Content-Type", "application/json")

	payload := map[string]string{
		"email":    "cristian.fiss@example.com",
		"password": "test1234",
	}
	jsonPayload, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/users/login", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// // TestGetProductByIdMagliaZara verifica che si ottenga un prodotto per ID
// func TestGetProductByIdMagliaZara(t *testing.T) {
// 	db := SetupTestDB()
// 	router := setupRouter(db)

// 	// Crea il prodotto "Maglia Zara" nel DB
// 	product := models.Product{
// 		Name:        "Maglia Zara",
// 		Description: "Maglia a maniche lunghe in cotone biologico.",
// 		Price:       39.90,
// 		Quantity:    100,
// 	}
// 	if err := db.Create(&product).Error; err != nil {
// 		t.Fatalf("Errore durante la creazione del prodotto: %v", err)
// 	}

// 	// Esegui la richiesta GET sul prodotto appena creato
// 	req, _ := http.NewRequest("GET", fmt.Sprintf("/users/products/%d", product.ID), nil)
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)

// 	// Controlla che lo status sia 200 OK
// 	assert.Equal(t, http.StatusOK, w.Code)

// 	// Verifica che il body contenga "Maglia Zara"
// 	assert.Contains(t, w.Body.String(), "Maglia Zara")
// 	assert.Contains(t, w.Body.String(), "cotone biologico")
// }

// // TestGetAllProducts verifica che vengano restituiti tutti i prodotti
// func TestGetAllProducts(t *testing.T) {
// 	db := SetupTestDB()
// 	router := setupRouter(db)

// 	// Crea alcuni prodotti nel DB
// 	products := []models.Product{
// 		{
// 			Name:        "Maglia Zara",
// 			Description: "Maglia a maniche lunghe in cotone biologico.",
// 			Price:       39.90,
// 			Quantity:    100,
// 		},
// 		{
// 			Name:        "Pantaloni Levi's",
// 			Description: "Pantaloni in denim di alta qualità.",
// 			Price:       89.90,
// 			Quantity:    50,
// 		},
// 	}
// 	for _, product := range products {
// 		if err := db.Create(&product).Error; err != nil {
// 			t.Fatalf("Errore durante la creazione del prodotto: %v", err)
// 		}
// 	}

// 	// Esegui la richiesta GET per ottenere tutti i prodotti
// 	req, _ := http.NewRequest("GET", "/users/products", nil)
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)

// 	// Controlla che lo status sia 200 OK
// 	assert.Equal(t, http.StatusOK, w.Code)

// 	// Verifica che il body contenga i prodotti creati
// 	for _, product := range products {
// 		assert.Contains(t, w.Body.String(), product.Name)
// 		assert.Contains(t, w.Body.String(), product.Description)
// 	}
// }
