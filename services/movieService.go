package services

import (
	"database/sql"
	"errors"
	"go-gin-mysql/models"
	"go-gin-mysql/repositories"

	"gorm.io/gorm"
)

var ErrMovieNotFound = errors.New("Movie not found")
var ErrInvalidMovieID = errors.New("Invalid movie ID")
var ErrMissingMovieData = errors.New("Movie must include title and director")

type MovieService struct {
	repo *repositories.MovieRepository
}

func NewMovieService(repo *repositories.MovieRepository) *MovieService {
	return &MovieService{
		repo: repo,
	}
}

func (s *MovieService) GetMovies() ([]models.Movie, error) {
	return s.repo.GetMovies()
}

func (s *MovieService) GetMovieByID(id int) (*models.Movie, error) {

	if id <= 0 {
		return nil, ErrInvalidMovieID
	}

	movie, err := s.repo.GetMovieByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMovieNotFound
		}
		return nil, err
	}

	return movie, nil
}

func (s *MovieService) AddMovie(movie *models.Movie) error {
	if movie.Title == "" || movie.Director == "" {
		return ErrMissingMovieData
	}

	return s.repo.AddMovie(movie)
}
func (s *MovieService) UpdateMovie(id int, movie *models.Movie) error {

	if id <= 0 {
		return ErrInvalidMovieID
	}

	if movie.Title == "" || movie.Director == "" {
		return ErrMissingMovieData
	}

	_, err := s.repo.GetMovieByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrMovieNotFound
		}
		return err
	}

	movie.ID = id
	return s.repo.UpdateMovie(movie)
}

func (s *MovieService) DeleteMovie(id int) error {
	if id <= 0 {
		return ErrInvalidMovieID
	}

	err := s.repo.DeleteMovie(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrMovieNotFound
		}
		return err
	}

	return nil
}
