package services

import (
	"log"
	"time"
	"users/modules/login/models"
	"users/modules/login/repositories"
	"users/pkg/auth"

	"github.com/dgrijalva/jwt-go"
)

type IService interface {
	LoginRequestService(login models.LoginRequest) (string, error)
}

type service struct {
	r repositories.IRepositorie
}

func NewService(r repositories.IRepositorie) IService {
	return &service{r: r}
}

func (s *service) LoginRequestService(login models.LoginRequest) (string, error) {
	userDetails, err := s.r.LoginRequestRepository(login)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// Set the expiration time for the JWT token
	expirationTime := time.Now().Add(time.Minute * 1).Unix()
	claims := models.JwtResponse{
		OrgmbTitle:   userDetails.OrgmbTitle,
		OrgmbName:    userDetails.OrgmbName,
		OrgmbSurname: userDetails.OrgmbSurname,
		OrgmbEmail:   userDetails.OrgmbEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}

	// Create JWT token
	tokenJWT, err := auth.CreateJWT(claims)
	if err != nil {
		log.Println("Failed to create token:", err)
		return "", err
	}

	return tokenJWT, nil
}
