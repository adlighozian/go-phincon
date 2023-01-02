package middlewares

import (
	"net/http"

	"github.com/adlighozian/go-belajar/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware(c *gin.Context) {
	y, err := c.Cookie("token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Status": "Error", "Message": "Unauthorized"})
		return
	}

	tokenString := y
	claims := &config.JWTClaim{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return config.JWT_KEY, nil
	})

	if err != nil {
		v, _ := err.(*jwt.ValidationError)
		switch v.Errors {
		case jwt.ValidationErrorSignatureInvalid:
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Status": "Error", "Message": "Unauthorized"})
			return
		case jwt.ValidationErrorExpired:
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Status": "Error", "Message": "token expired"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Status": "Error", "Message": "Unauthorized"})
			return
		}
	}

	if !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Status": "Error", "Message": "error"})
		return
	}

}
