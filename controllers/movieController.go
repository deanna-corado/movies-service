// CONTROLLER - BUILD RESPONSES TO REQUESTS AND INTERACT WITH THE DATABASE
package controllers

import (
	"net/http"
	"strconv"

	"go-gin-mysql/models"
	"go-gin-mysql/services"
	"go-gin-mysql/utils"

	"github.com/gin-gonic/gin"
)

type MovieController struct {
	service *services.MovieService
}

func NewMovieController(service *services.MovieService) *MovieController {
	return &MovieController{
		service: service,
	}
}

func (c *MovieController) GetMovies(ctx *gin.Context) {
	movies, err := c.service.GetMovies()
	if err != nil {

		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, movies)
}

func (c *MovieController) GetMovieByID(ctx *gin.Context) {

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid movie ID"})
		return
	}

	movie, err := c.service.GetMovieByID(id)
	if err != nil {

		utils.HandleMovieError(ctx, err)
		return
	}
	ctx.JSON(200, movie)
}

func (c *MovieController) AddMovie(ctx *gin.Context) {
	var movie models.Movie

	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid JSON",
		})
		return
	}

	if err := c.service.AddMovie(&movie); err != nil {
		utils.HandleMovieError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, movie)
}
func (c *MovieController) UpdateMovie(ctx *gin.Context) {
	var movie models.Movie

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid movie ID"})
		return
	}

	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := c.service.UpdateMovie(id, &movie); err != nil {
		utils.HandleMovieError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, movie)
}

func (c *MovieController) DeleteMovie(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid movie ID"})
		return
	}

	if err := c.service.DeleteMovie(id); err != nil {
		utils.HandleMovieError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{"message": "Movie deleted"})
}
