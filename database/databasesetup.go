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
	urlDBNoSchema := "root:root@tcp(localhost:3306)/"
	dbNoSchema, err := gorm.Open(mysql.Open(urlDBNoSchema), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database server: %v", err)
	}

	// Crea il database "ecommerce" se non esiste
	createDBQuery := "CREATE DATABASE IF NOT EXISTS ecommerce"
	if err := dbNoSchema.Exec(createDBQuery).Error; err != nil {
		return nil, fmt.Errorf("failed to create database: %v", err)
	}

	// Dettagli di connessione MySQL con il database "ecommerce"
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

	// Crea le tabelle se non esistono già
	err = db.AutoMigrate(&models.User{}, &models.Product{})
	if err != nil {
		fmt.Println("Errore durante la migrazione delle tabelle: %v", err)
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