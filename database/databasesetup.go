package database

import (
	"fmt"
	"github.com/cristian-fioravanti/ecommerceGo/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Client è il puntatore globale al database GORM
var Client *gorm.DB

// DBSet stabilisce la connessione al database utilizzando GORM
func DBSet() (*gorm.DB, error) {
	// Dettagli di connessione MySQL
	urlDB := "root:root@tcp(localhost:3306)/ecommerce"
	db, err := gorm.Open(mysql.Open(urlDB), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Convalida della connessione
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get raw database connection: %v", err)
	}

	// Verifica che la connessione al database sia attiva
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	fmt.Println("Successfully connected to MySQL database")
	return db, nil
}

// Inizializza la connessione al database
func init() {
	var err error
	Client, err = DBSet()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		// Gestisci l'errore (potresti voler eseguire un panic o uscire)
	}
}

func UserData() ([]models.User, error) {
	var users []models.User
	if err := Client.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func ProductData() ([]models.Product, error) {
	var products []models.Product
	if err := Client.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
