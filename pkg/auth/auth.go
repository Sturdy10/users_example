package auth

import (
	"os"
	"time"
	"users/modules/login/models"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// CreateTokenI สร้าง JWT token จาก email โดยมีอายุ 1 นาที
func CreateTokenI(email string) (string, error) {
	// สร้าง claims สำหรับ token โดยใช้ email และกำหนดเวลา expiration
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Minute * 1).Unix(), // token มีอายุ 1 นาที
	}

	// สร้าง token ด้วย method HS256 และ claims ที่กำหนด
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// ลงนาม token ด้วย secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CreateJWT(claims models.JwtResponse) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("thenilalive"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
