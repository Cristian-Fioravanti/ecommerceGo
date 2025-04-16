package controllers

import (
	"fmt"
	"net/http"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
"gorm.io/gorm"
generate "github.com/cristian-fioravanti/ecommerceGo/tokens"
"github.com/cristian-fioravanti/ecommerceGo/models"
"log"
"strconv")

var validate = validator.New()
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Panic(err)
	}
	return string(hashedPassword), nil
}

func VerifyPassword(userPassword string, givenPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(givenPassword), []byte(userPassword))
	if err != nil {
		return false, "Password is incorrect"
	}
	return true, "Password is correct"
}
func Signup(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var user models.User

        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        if validationErr := validate.Struct(user); validationErr != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
            return
        }

        // Check if user already exists
        var existingUser models.User
        if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
            c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
            return
        }

        // Hash password
        hashedPassword, err := HashPassword(user.Password)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
            return
        }
        user.Password = hashedPassword

        // Generate tokens
        token, refreshToken, _ := generate.TokenGenerator(user.Email, user.FirstName, user.LastName, "") // puoi passare l'ID dopo insert

        user.Token = &token
        user.RefreshToken = &refreshToken

        // Insert user into DB
        if err := db.Create(&user).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
            return
        }

        // Ora user.ID è valorizzato (auto-increment)
        c.JSON(http.StatusCreated, gin.H{"user": user})
    }
}

func Login(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var loginData struct {
            Email    string `json:"email" binding:"required,email"`
            Password string `json:"password" binding:"required"`
        }

        if err := c.ShouldBindJSON(&loginData); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        var user models.User
        if err := db.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }

        isPasswordValid, msg := VerifyPassword(loginData.Password, user.Password)
        if !isPasswordValid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
            return
        }

        // Generate new tokens
        token, refreshToken, _ := generate.TokenGenerator(user.Email, user.FirstName, user.LastName, fmt.Sprint(user.ID))
        
        if err := generate.UpdateAllTokens(db, &user, token, refreshToken); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update tokens"})
            return
        }
       
		user.Password = ""
        c.JSON(http.StatusOK, gin.H{"user": user})
    }
}


func GetProductById(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        productIdParam := c.Param("productId")

        // Converti il parametro stringa in intero
        productId, err := strconv.Atoi(productIdParam)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
            return
        }

        var product models.Product
        if err := db.First(&product, productId).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
            return
        }

        c.JSON(http.StatusOK, product)
    }
}

func GetProducts(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var products []models.Product

        // Recupera tutti i prodotti
        if err := db.Find(&products).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
            return
        }

        c.JSON(http.StatusOK, products)
    }
}

func CreateProduct(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var product models.Product

        if err := c.ShouldBindJSON(&product); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if validationErr := validate.Struct(product); validationErr != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error validation": validationErr.Error()})
            return
        }
        if err := db.Create(&product).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
            return
        }

        c.JSON(http.StatusCreated, gin.H{"product": product})
    }
}
func UpdateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		productIdParam := c.Param("productId")
		productId, err := strconv.Atoi(productIdParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
		}
    
		var product models.Product
		if err := db.First(&product, productId).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
        if validationErr := validate.Struct(product); validationErr != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error validation": validationErr.Error()})
            return
        }
		if err := db.Save(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"product": product})
	}
}

func DeleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		productIdParam := c.Param("productId")

		productId, err := strconv.Atoi(productIdParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
		}

		var product models.Product
		if err := db.First(&product, productId).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		if err := db.Delete(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
	}
}