package repositories

import (
	"database/sql"
	"go-gin-mysql/models"
)

// medyo lito pa dito
type MovieRepository struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{
		db: db,
	}
}

func (r *MovieRepository) GetMovies() ([]models.Movie, error) {
	query := "SELECT id, title, director FROM movies"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Director)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return movies, nil
}

func (r *MovieRepository) GetMovieByID(id int) (*models.Movie, error) {
	query := "SELECT id, title, director FROM movies WHERE id = ?"

	row := r.db.QueryRow(query, id)

	var movie models.Movie
	err := row.Scan(&movie.ID, &movie.Title, &movie.Director)
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func (r *MovieRepository) AddMovie(movie *models.Movie) error {
	query := "INSERT INTO movies (title, director) VALUES (?, ?)"

	result, err := r.db.Exec(query, movie.Title, movie.Director)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	movie.ID = int(id)
	return nil
}

func (r *MovieRepository) UpdateMovie(movie *models.Movie) error {

	query := "UPDATE movies SET title = ?, director = ? WHERE id = ?"
	result, err := r.db.Exec(query, movie.Title, movie.Director, movie.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *MovieRepository) DeleteMovie(id int) error {
	query := "DELETE FROM movies WHERE id = ?"

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

// func atoi(s string) int {
// 	i, _ := strconv.Atoi(s)
// 	return i
// }
