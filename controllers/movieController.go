// CONTROLLER - BUILD RESPONSES TO REQUESTS AND INTERACT WITH THE DATABASE
package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"go-gin-mysql/models"
	"go-gin-mysql/repositories"

	"github.com/gin-gonic/gin"
)

type MovieController struct {
	repo *repositories.MovieRepository
}

func NewMovieController(repo *repositories.MovieRepository) *MovieController {
	return &MovieController{
		repo: repo,
	}
}

func (c *MovieController) GetMovies(ctx *gin.Context) {
	movies, err := c.repo.GetMovies()
	if err != nil {

		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, movies)
}

func (c *MovieController) GetMovieByID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	movie, err := c.repo.GetMovieByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	ctx.JSON(http.StatusOK, movie)
}

func (c *MovieController) AddMovie(ctx *gin.Context) {
	var movie models.Movie
	//DIFFERENCE BETWEEN THESE 2?
	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON",
		})
		return
	}

	if err := c.repo.AddMovie(&movie); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, movie)
}

func (c *MovieController) UpdateMovie(ctx *gin.Context) {
	var movie models.Movie

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	movie.ID = id

	if err := c.repo.UpdateMovie(&movie); err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, movie)
}

func (c *MovieController) DeleteMovie(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	if err := c.repo.DeleteMovie(id); err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Movie deleted"})
}
