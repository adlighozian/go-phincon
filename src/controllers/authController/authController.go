package authController

import (
	"net/http"
	"time"

	"github.com/adlighozian/go-belajar/config"
	"github.com/adlighozian/go-belajar/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Error", "message": err.Error()})
		return
	}
	var auth models.User
	if err := models.DB.Where("username = ?", user.Username).First(&auth).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "Error", "message": "username atau password salah"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": err.Error()})
			return
		}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(user.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "username atau password salah"})
		return
	}
	expTime := time.Now().Add(time.Minute * 60)
	claims := &config.JWTClaim{
		Username: auth.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-belajar",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Status": "Error", "Message": err.Error()})
		return
	}
	c.SetCookie("token", token, 0, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"Status": "Login berhasil", "Token": config.JWT_KEY})
}

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": "Error", "Message": err.Error()})
		return
	}
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashPassword)
	models.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"Status": "Success", "User": user})
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "token", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"Status": "Logout berhasil"})
}
