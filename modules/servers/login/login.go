package login

import (
	"database/sql"
	"users/modules/login/handlers"
	"users/modules/login/repositories"
	"users/modules/login/services"

	"github.com/gin-gonic/gin"
)

func Login(router *gin.Engine, db *sql.DB) {

	r := repositories.NewRepositorie(db)
	s := services.NewService(r)
	h := handlers.NewHandler(s)

	router.POST("/api/login", h.LoginRequestHandler)
}


