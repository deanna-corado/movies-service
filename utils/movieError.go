package utils

import (
	"errors"

	"movies-service/services"

	"github.com/gin-gonic/gin"
)

func HandleMovieError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, services.ErrInvalidMovieID),
		errors.Is(err, services.ErrMissingMovieData):
		ctx.JSON(400, gin.H{"error": err.Error()})

	case errors.Is(err, services.ErrMovieNotFound):
		ctx.JSON(404, gin.H{"error": err.Error()})

	default:
		ctx.JSON(500, gin.H{"error": err.Error()})

	}
}
