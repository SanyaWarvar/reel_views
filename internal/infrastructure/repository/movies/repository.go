package movies

import (
	"context"
	"fmt"
	"rv/internal/domain/dto/movies"
	"rv/pkg/database/postgres"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
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

	subBuilder := squirrel.Select("array_agg(g.name)").
		From("movie_genre mg").
		Join("genres g ON g.id = mg.genre_id").
		Where("mg.movie_id = m.id")
	genreSubquery, args, _ := subBuilder.PlaceholderFormat(squirrel.Dollar).ToSql()
	// Основной запрос
	query := squirrel.Select(
		"m.id, m.title, m.img_url",
		"("+genreSubquery+") AS genres",
		"coalesce(ROUND(AVG(r.rating), 2), 0) AS avg_rating",
	).
		From("movies m").
		LeftJoin("reviews r ON m.id = r.movie_id").
		GroupBy("m.id").
		OrderBy("m.created_at")

	if movieFilter.Id != nil {
		query = query.Where(squirrel.Eq{"m.id": movieFilter.Id})
	}
	if len(movieFilter.TitleIn) > 0 {
		query = query.Where(squirrel.Eq{"m.title": movieFilter.TitleIn})
	}
	if movieFilter.RatingGOE != nil {
		query = query.Where(squirrel.GtOrEq{"avg_rating": *movieFilter.RatingGOE})
	}
	if movieFilter.Search != "" {
		query = query.Where(squirrel.Like{"m.title": fmt.Sprintf("%%%s%%", movieFilter.Search)})
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	if limit > 0 {
		query = query.Limit(limit)
	}

	// Получаем SQL и аргументы
	sql, args, err := query.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

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

func (repo *Repository) GetRecomendationsForMovie(ctx context.Context, movieId uuid.UUID) ([]movies.MoviesShort, error) {
	query := `
	select m.id, m.title, m.img_url, 
	(select array_agg(g.name) from movie_genre mg join genres g on g.id = mg.genre_id where mg.movie_id = m.id) as genres,
	coalesce(ROUND(AVG(r2.rating), 2), 0) AS avg_rating,
	ROUND(r.similarity_score::numeric, 2) as similarity_score
	from get_similar_movies($1) r 
	join movies m on m.id = r.similar_movie_id
	left join reviews r2 on r2.movie_id = m.id
	group by m.id, r.similarity_score
	order by r.similarity_score desc
	`
	rows, err := repo.conn.Query(ctx, query, movieId)
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
			&m.SimilarityScore,
		)
		if err != nil {
			return nil, err
		}
		output = append(output, m)
	}
	return output, err
}

func (repo *Repository) GetRecomendationsForUser(ctx context.Context, userId uuid.UUID) ([]movies.MoviesShort, error) {
	query := `
	select m.id, m.title, m.img_url, 
	(select array_agg(g.name) from movie_genre mg join genres g on g.id = mg.genre_id where mg.movie_id = m.id) as genres,
	coalesce(ROUND(AVG(r2.rating), 2), 0) AS avg_rating,
	ROUND(r.recommendation_score::numeric, 2) as similarity_score
	from get_user_recommendations($1) r 
	join movies m on m.id = r.movie_id
	left join reviews r2 on r2.movie_id = m.id
	group by m.id, r.recommendation_score
	order by r.recommendation_score desc;
	`
	rows, err := repo.conn.Query(ctx, query, userId)
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
			&m.SimilarityScore,
		)
		if err != nil {
			return nil, err
		}
		output = append(output, m)
	}
	return output, err
}
