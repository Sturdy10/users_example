package password

import (
	"database/sql"
	"users/modules/password/handlers"
	"users/modules/password/repositories"
	"users/modules/password/services"

	"github.com/gin-gonic/gin"
)

func Password(router *gin.Engine, db *sql.DB) {

	r := repositories.NewRepositorie(db)
	s := services.NewService(r)
	h := handlers.NewHandler(s)

	router.PATCH("/api/newPassword", h.InitPasswordHandler)
	router.PATCH("/api/changePassword", h.ChangePasswordHandler)
}
