package repositories

import (
	"database/sql"
	"errors"
	"log"
	"users/modules/login/models"

	"golang.org/x/crypto/bcrypt"
)

type IRepositorie interface {
	LoginRequestRepository(login models.LoginRequest) (models.JwtResponse, error) 
}

type repository struct {
	db *sql.DB
}

func NewRepositorie(db *sql.DB) IRepositorie {
	return &repository{db: db}
}

func (r *repository) LoginRequestRepository(login models.LoginRequest) (models.JwtResponse, error) {
    var orgmbID string
    err := r.db.QueryRow("SELECT orgmb_id FROM organize_member WHERE orgmb_email = $1", login.OrgmbEmail).Scan(&orgmbID)
    if err != nil {
        log.Println("failed to find organize_member:", err)
        return models.JwtResponse{}, errors.New("failed to find email address")
    }

    var orgmbcrPassword string
    err = r.db.QueryRow("SELECT orgmbcr_password FROM organize_member_credential WHERE orgmbcr_orgmb_id = $1", orgmbID).Scan(&orgmbcrPassword)
    if err != nil {
        log.Println("failed to update organize_member_credential:", err)
        return models.JwtResponse{}, err
    }

    err = bcrypt.CompareHashAndPassword([]byte(orgmbcrPassword), []byte(login.OrgmbcrPassword))
    if err != nil {
        log.Println("the password is incorrect:", err)
        return models.JwtResponse{}, errors.New("the password is incorrect")
    }

    // Query additional details required for JwtResponse
    var jwtResponse models.JwtResponse
    err = r.db.QueryRow("SELECT orgmb_title, orgmb_name, orgmb_surname, orgmb_email FROM organize_member WHERE orgmb_id = $1", orgmbID).
        Scan(&jwtResponse.OrgmbTitle, &jwtResponse.OrgmbName, &jwtResponse.OrgmbSurname, &jwtResponse.OrgmbEmail)
    if err != nil {
        log.Println("failed to get user details:", err)
        return models.JwtResponse{}, errors.New("failed to get user details")
    }

    return jwtResponse, nil
}
