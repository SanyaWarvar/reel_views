package movies

import (
	"context"
	"rv/internal/domain/dto/movies"
	"rv/pkg/database/postgres"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

type Repository struct {
	conn postgres.Connection
}

func NewRepository(conn postgres.Connection) *Repository {
	return &Repository{conn: conn}
}

func (repo *Repository) CreateMovie(ctx context.Context, movie Movie) error {
	query, args, err := squirrel.Insert("movies").
		Values(movie.Id, movie.Title, movie.Description, movie.ImgUrl, movie.Meta, movie.CreatedAt).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return errors.Wrap(err, "squirrel.ToSql")
	}

	_, err = repo.conn.Exec(ctx, query, args...)
	return err
}

func (repo *Repository) GetMoviesShort(ctx context.Context, movieFilter MovieFilter, offset, limit uint64) ([]movies.MoviesShort, error) {
	// Субзапрос для жанров
	genreSubquery, _, err := squirrel.Select("array_agg(g.name)").
		From("movie_genre mg").
		Join("genres g ON g.id = mg.genre_id").
		Where("mg.movie_id = m.id").ToSql()

	// Основной запрос
	query := squirrel.Select(
		"m.id, m.title, m.img_url",
		"("+genreSubquery+") AS genres",
		"ROUND(AVG(r.rating), 2) AS avg_rating",
	).
		From("movies m").
		RightJoin("reviews r ON m.id = r.movie_id").
		GroupBy("m.id").
		OrderBy("m.created_at")

	// Получаем SQL и аргументы
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	// Выполняем запрос
	rows, err := repo.conn.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var output []movies.MoviesShort
	for rows.Next() {
		var m movies.MoviesShort
		err := rows.Scan(
			&m.Id,
			&m.Title,
			&m.ImgUrl,
			&m.Genres,
			&m.AvgRating,
		)
		if err != nil {
			return nil, err
		}
		output = append(output, m)
	}

	return output, nil
}
