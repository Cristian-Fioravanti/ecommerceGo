package middleware

import (
	token "github.com/cristian-fioravanti/ecommerceGo/tokens"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"gorm.io/gorm"
	"github.com/cristian-fioravanti/ecommerceGo/models"
)

func Authentication() gin.HandlerFunc{
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization header provided"})
			c.Abort()
			return
		}
		tokenString := strings.Split(clientToken, "Bearer ")[1]
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}
		claims, err := token.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("id", claims.ID)
		c.Next()
	}
}
func AdminMiddleware(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization header provided"})
			c.Abort()
			return
		}
		tokenString := strings.Split(clientToken, "Bearer ")[1]
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}
		claims, err := token.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

        var user models.User
        if err := db.Where("email = ?", claims.Email).First(&user).Error; err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
            c.Abort()
            return
        }

        if user.Role != "admin" {
            c.JSON(http.StatusForbidden, gin.H{"error": "Access denied: Admins only"})
            c.Abort()
            return
        }

        c.Set("email", claims.Email)
		c.Set("id", claims.ID)
		c.Next()
    }
}