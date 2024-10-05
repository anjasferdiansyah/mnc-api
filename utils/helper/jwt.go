package helper

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

const secretKey = "access-login"

func GenerateToken(id uint, email string, role string) (string, error) {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"role": role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToke(c *gin.Context) (jwt.MapClaims, error) {
	errResponse := errors.New("token invalid")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	// Pisahkan token dari "Bearer"
	stringToken := strings.Split(headerToken, " ")[1]

	// Parsing token
	token, err := jwt.Parse(stringToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		// Menggunakan secretKey untuk memvalidasi token
		return []byte(secretKey), nil
	})

	// Jika parsing token gagal atau token tidak valid
	if err != nil || !token.Valid {
		return nil, errResponse
	}

	// Ambil claims dari token dan pastikan bentuknya adalah jwt.MapClaims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errResponse
	}

	return claims, nil
}
