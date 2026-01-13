// ROUTES - FOR CALLING CONTROLLERS BASED ON ENDPOINTS; CALL API
package routes

import (
	// "database/sql"
	// "go-gin-mysql/controllers"
	// "go-gin-mysql/repositories"

	"go-gin-mysql/controllers"

	"github.com/gin-gonic/gin"
)

// FIX BY GROUP; UNDESTAND MORE
func RegisterRoutes(
	r *gin.Engine,
	movieController *controllers.MovieController,
) {
	r.GET("/movies", movieController.GetMovies)
	r.GET("/movies/:id", movieController.GetMovieByID)
	r.POST("/movies", movieController.AddMovie)
	r.PUT("/movies/:id", movieController.UpdateMovie)
	r.DELETE("/movies/:id", movieController.DeleteMovie)
}
