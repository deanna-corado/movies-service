package middlewares

// import (
// 	"log"
// 	"os"

// 	"github.com/gin-gonic/gin"
// )

// func AuthRequired() gin.HandlerFunc {

// 	user := os.Getenv("ADMIN_USER")
// 	pass := os.Getenv("ADMIN_PASS")

// 	if user == "" || pass == "" {
// 		log.Fatal("ADMIN_USER or ADMIN_PASS environment variables are not set")

// 	}

// 	return func(c *gin.Context) {
// 		username, password, ok := c.Request.BasicAuth()

// 		if !ok || username != user || password != pass {
// 			c.JSON(401, gin.H{"error": "Unauthorized"})
// 			c.Abort()
// 			return
// 		}

// 		c.Next()
// 	}
// }
