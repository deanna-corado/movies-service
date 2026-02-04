// ROUTES - FOR CALLING CONTROLLERS BASED ON ENDPOINTS; CALL API
package routes

import (
	"movies-service/controllers"
	"movies-service/middlewares"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func RegisterRoutes(
	r *gin.Engine,
	movieController *controllers.MovieController, credController *controllers.CredentialController,
) {

	r.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
	))

	//for api versioning
	v1 := r.Group("/api/v1")

	movies := v1.Group("/movies", middlewares.ClientCredentialAuth())
	{
		movies.GET("", movieController.GetMovies)
		movies.GET("/:id", movieController.GetMovieByID)
		movies.POST("", movieController.AddMovie)
		movies.PUT("/:id", movieController.UpdateMovie)
		movies.DELETE("/:id", movieController.DeleteMovie)
	}

	v1.POST("/credentials/validate", credController.Validate)
}
