package controllers

import (
	"movies-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CredentialController struct {
	db *gorm.DB
}

func NewCredentialController(db *gorm.DB) *CredentialController {
	return &CredentialController{db: db}
}

// validate client credentials
func (cc *CredentialController) Validate(c *gin.Context) {

	var body struct {
		ClientID string `json:"client_id"`
		Secret   string `json:"secret"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	//get credential from db
	var cred models.Credential
	if err := cc.db.First(&cred, "client_id = ?", body.ClientID).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid client"})
		return
	}

	if cred.SecretKey != body.Secret {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid secret"})
		return
	}

	token := body.ClientID + "-token"
	c.JSON(http.StatusOK, gin.H{"token": token})
}
