package tokens
 
import (jwt "github.com/dgrijalva/jwt-go"
"time"
    "github.com/cristian-fioravanti/ecommerceGo/models"
"fmt"                                  
    "gorm.io/gorm"                        )


var jwtKey = []byte("my_secret_key") 

// TokenGenerator crea un token JWT e un refresh token
func TokenGenerator(email, firstName, lastName, userID string) (string, string, error) {
    // Dati del token
    claims := &models.SignedDetails{
        Email:     email,
        FirstName: firstName,
        LastName:  lastName,
        ID:        userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), // Imposta la scadenza del token a 1 ora
        },
    }

    // Creazione dell'access token
    accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := accessToken.SignedString(jwtKey)
    if err != nil {
        return "", "", err
    }

    // Creazione del refresh token (di lunga durata)
    refreshClaims := &models.SignedDetails{
        Email:     email,
        FirstName: firstName,
        LastName:  lastName,
        ID:        userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(), // Imposta la scadenza del refresh token a 7 giorni
        },
    }
    refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
    refreshTokenString, err := refreshToken.SignedString(jwtKey)
    if err != nil {
        return "", "", err
    }

    return tokenString, refreshTokenString, nil
}

// ValidateToken verifica se il token JWT è valido
func ValidateToken(tokenString string) (*models.SignedDetails, error) {
    claims := &models.SignedDetails{}

    // Parse e verifica la firma del token
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil || !token.Valid {
        return nil, fmt.Errorf("invalid token")
    }

    return claims, nil
}

// UpdateAllTokens aggiorna i token dell'utente nel database
func UpdateAllTokens(db *gorm.DB, user *models.User, token, refreshToken string) error {
   
    user.Token = &token
    user.RefreshToken = &refreshToken

    if err := db.Save(&user).Error; err != nil {
        return err 
    }

    return nil 
}

