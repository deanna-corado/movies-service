// CONTROLLER - BUILD RESPONSES TO REQUESTS AND INTERACT WITH THE DATABASE
package controllers

import (
	"net/http"
	"strconv"

	"movies-service/models"
	"movies-service/services"
	"movies-service/utils"

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

// GetMovies godoc
// @Summary Get all movies
// @Description Get a list of all movies from the database
// @Tags movies
// @Accept json
// @Produces json
// @Success 200 {array} models.Movie
// @Failure 404 {object} map[string]string
// @Router /movies [get]
func (c *MovieController) GetMovies(ctx *gin.Context) {
	movies, err := c.service.GetMovies()
	if err != nil {

		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, movies)
}

// GetMovieByID godoc
// @Summary Get a movie by ID
// @Description Get a movie by its ID from the database
// @Tags movies
// @Accept json
// @Param id path int true "Movie ID"
// @Produces json
// @Success 200 {object} models.Movie
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /movies/{id} [get]
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

// AddMovie godoc
// @Summary Add a new movie
// @Description Add a new movie to the database
// @Tags admin-movies
// @Security BasicAuth
// @Accept json
// @Param movie body models.Movie true "Add Movie"
// @Produces json
// @Success 201 {object} models.Movie
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /admin/movies [post]
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

// UpdateMovie godoc
// @Summary Update an existing movie
// @Description Update an existing movie in the database
// @Tags admin-movies
// @Security BasicAuth
// @Accept json
// @Param id path int true "Movie ID"
// @Param movie body models.Movie true "Update Movie"
// @Produces json
// @Success 200 {object} models.Movie
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /admin/movies/{id} [put]
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

// DeleteMovie godoc
// @Summary Delete a movie
// @Description Delete a movie from the database
// @Tags admin-movies
// @Security BasicAuth
// @Accept json
// @Param id path int true "Movie ID"
// @Produces json
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /admin/movies/{id} [delete]
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
