package middlewares

import (
	"movies-service/config"
	"movies-service/models"

	"github.com/gin-gonic/gin"
)

func ClientCredentialAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		clientID := c.GetHeader("X-Client-ID")
		secret := c.GetHeader("X-Client-Secret")

		if clientID == "" || secret == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Missing client credentials"})
			return
		}

		var cred models.Credential
		if err := config.DB.
			Where("client_id = ? AND secret_key = ?", clientID, secret).
			First(&cred).Error; err != nil {

			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid client credentials"})
			return
		}

		c.Next()
	}
}
