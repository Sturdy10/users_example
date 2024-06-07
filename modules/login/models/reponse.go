package models

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type LoginResponse struct {
	OrgmbTitle   string `json:"title"`
	OrgmbName    string `json:"name"`
	OrgmbSurname string `json:"surname"`
	OrgmbEmail   string `json:"email"`
	OrgdpName    string `json:"department"`
}

// JwtResponse represents the JWT claims

type JwtResponse struct {
	OrgmbTitle   string `json:"title"`
	OrgmbName    string `json:"name"`
	OrgmbSurname string `json:"surname"`
	OrgmbEmail   string `json:"email"`
	jwt.StandardClaims
}

// Valid checks the validity of the token
func (c JwtResponse) Valid() error {
	if c.StandardClaims.ExpiresAt < time.Now().Unix() {
		return errors.New("token has expired")
	}
	return nil
}
