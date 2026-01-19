package repositories

import (
	"go-gin-mysql/models"

	"gorm.io/gorm"
)

type MovieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{
		db: db,
	}
}

func (r *MovieRepository) GetMovies() ([]models.Movie, error) {

	var movies []models.Movie
	if err := r.db.Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

func (r *MovieRepository) GetMovieByID(id int) (*models.Movie, error) {
	var movie models.Movie
	if err := r.db.First(&movie, id).Error; err != nil {
		return nil, err
	}
	return &movie, nil
}

func (r *MovieRepository) AddMovie(movie *models.Movie) error {
	return r.db.Create(movie).Error
}

func (r *MovieRepository) UpdateMovie(movie *models.Movie) error {
	result := r.db.Save(movie)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

func (r *MovieRepository) DeleteMovie(id int) error {
	result := r.db.Delete(&models.Movie{}, id)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}
