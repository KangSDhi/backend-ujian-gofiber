package config

import (
	"backend-ujian-gofiber/src/models"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strconv"
	"time"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateJWT(pengguna models.Pengguna) (string, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   pengguna.ID,
		"role": pengguna.RolePengguna,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})

	return token.SignedString(privateKey)
}
