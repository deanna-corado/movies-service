// ROUTES - FOR CALLING CONTROLLERS BASED ON ENDPOINTS; CALL API
package routes

import (
	"movies-service/controllers"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func RegisterRoutes(
	r *gin.Engine,
	movieController *controllers.MovieController,
) {

	r.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
	))

	//for api versioning
	v1 := r.Group("/api/v1")

	//PUBLIC ROUTES
	movies := v1.Group("/movies")
	{
		movies.GET("", movieController.GetMovies)
		movies.GET("/:id", movieController.GetMovieByID)
		movies.POST("", movieController.AddMovie)
		movies.PUT("/:id", movieController.UpdateMovie)
		movies.DELETE("/:id", movieController.DeleteMovie)
	}

	//PRIVATE ROUTES (ADMIN ACCESS)
	// 	adminMovies := v1.Group("/admin/movies", middlewares.AuthRequired())
	// 	{
	// 		adminMovies.POST("", movieController.AddMovie)
	// 		adminMovies.PUT("/:id", movieController.UpdateMovie)
	// 		adminMovies.DELETE("/:id", movieController.DeleteMovie)
	// 	}
	// }
}
