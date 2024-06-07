package handlers

import (
	"net/http"
	"users/modules/login/models"
	"users/modules/login/services"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	LoginRequestHandler(c *gin.Context)
}

type handler struct {
	s services.IService
}

func NewHandler(s services.IService) IHandler {
	return &handler{s: s}
}

func (h *handler) LoginRequestHandler(c *gin.Context) {
	var login models.LoginRequest
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Error", "message": err.Error()})
		return
	}
	tokenJWT, err := h.s.LoginRequestService(login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Login successfully", "token": tokenJWT})
}
