package models

import (jwt "github.com/dgrijalva/jwt-go"
)

type User struct {
    ID           uint   `gorm:"primaryKey;autoIncrement" json:"id"`
    FirstName    string `json:"first_name" validate:"required,min=2,max=100"`
    LastName     string `json:"last_name" validate:"required,min=2,max=100"`
    Email        string `gorm:"unique" json:"email" validate:"required,email"`
    Password     string `json:"password" validate:"required,min=6"`
    Token        *string `json:"token"`
    RefreshToken *string `json:"refresh_token"`
	Role 	   string `json:"role" gorm:"default:'user'"` 
}
type Product struct {
	ID           uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string   `json:"name" gorm:"not null" validate:"required,min=2,max=100"`
	Description string   `json:"description" gorm:"not null" validate:"min=10,max=500"`
	Price       float64  `json:"price" gorm:"not null" validate:"required,gt=0"`
	Quantity    int      `json:"quantity" gorm:"not null" validate:"required,gte=0"`
}
type SignedDetails struct {
	Email string	`json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	ID string `json:"id"`
	jwt.StandardClaims
}