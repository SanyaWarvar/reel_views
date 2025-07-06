package reviews

import (
	"context"
	"rv/internal/domain/dto/reviews"
	apperrors "rv/internal/errors"
	"rv/internal/infrastructure/repository/common"
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

// CreateReview создает новый отзыв
func (repo *Repository) CreateReview(ctx context.Context, review reviews.Review) error {
	query, args, err := squirrel.Insert("reviews").
		Columns("id", "movie_id", "user_id", "description", "rating", "created_at").
		Values(review.Id, review.MovieId, review.UserId, review.Description, review.Rating, review.CreatedAt).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return errors.Wrap(err, "squirrel.ToSql")
	}

	_, err = repo.conn.Exec(ctx, query, args...)
	if err != nil {
		if common.IsUniqueErr(err) {
			return apperrors.NotUnique
		}
	}
	return nil
}

// GetReviewByID возвращает отзыв по ID
func (repo *Repository) GetReviewByID(ctx context.Context, id uuid.UUID) (*reviews.Review, error) {
	query, args, err := squirrel.Select("id", "movie_id", "user_id", "description", "rating", "created_at").
		From("reviews").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "squirrel.ToSql")
	}

	var review reviews.Review
	err = repo.conn.QueryRow(ctx, query, args...).Scan(
		&review.Id,
		&review.MovieId,
		&review.UserId,
		&review.Description,
		&review.Rating,
		&review.CreatedAt,
	)
	if err != nil {
		return nil, errors.Wrap(err, "repo.conn.QueryRow")
	}

	return &review, nil
}

// GetReviewsByMovie возвращает отзывы для фильма
func (repo *Repository) GetReviewsByMovie(ctx context.Context, movieID uuid.UUID, limit, offset uint64) ([]reviews.Review, error) {
	query := squirrel.Select("id", "movie_id", "user_id", "description", "rating", "created_at").
		From("reviews").
		Where(squirrel.Eq{"movie_id": movieID}).
		OrderBy("created_at DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	sql, args, err := query.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "squirrel.ToSql")
	}

	rows, err := repo.conn.Query(ctx, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "repo.conn.Query")
	}
	defer rows.Close()

	var result []reviews.Review
	for rows.Next() {
		var review reviews.Review
		if err := rows.Scan(
			&review.Id,
			&review.MovieId,
			&review.UserId,
			&review.Description,
			&review.Rating,
			&review.CreatedAt,
		); err != nil {
			return nil, errors.Wrap(err, "rows.Scan")
		}
		result = append(result, review)
	}

	return result, nil
}

// GetReviewsByUser возвращает отзывы пользователя
func (repo *Repository) GetReviewsByUser(ctx context.Context, userId uuid.UUID, limit, offset uint64) ([]reviews.Review, error) {
	query := squirrel.Select("id", "movie_id", "user_id", "description", "rating", "created_at").
		From("reviews").
		Where(squirrel.Eq{"user_id": userId}).
		OrderBy("created_at DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	sql, args, err := query.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "squirrel.ToSql")
	}

	rows, err := repo.conn.Query(ctx, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "repo.conn.Query")
	}
	defer rows.Close()

	var result []reviews.Review
	for rows.Next() {
		var review reviews.Review
		if err := rows.Scan(
			&review.Id,
			&review.MovieId,
			&review.UserId,
			&review.Description,
			&review.Rating,
			&review.CreatedAt,
		); err != nil {
			return nil, errors.Wrap(err, "rows.Scan")
		}
		result = append(result, review)
	}

	return result, nil
}

// UpdateReview обновляет отзыв
func (repo *Repository) UpdateReview(ctx context.Context, review reviews.Review) error {
	query, args, err := squirrel.Update("reviews").
		Set("description", review.Description).
		Set("rating", review.Rating).
		Where(squirrel.Eq{"id": review.Id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return errors.Wrap(err, "squirrel.ToSql")
	}

	_, err = repo.conn.Exec(ctx, query, args...)
	return errors.Wrap(err, "repo.conn.Exec")
}

// DeleteReview удаляет отзыв
func (repo *Repository) DeleteReview(ctx context.Context, id uuid.UUID) error {
	query, args, err := squirrel.Delete("reviews").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return errors.Wrap(err, "squirrel.ToSql")
	}

	_, err = repo.conn.Exec(ctx, query, args...)
	return errors.Wrap(err, "repo.conn.Exec")
}

// GetAverageRating возвращает средний рейтинг фильма
func (repo *Repository) GetAverageRating(ctx context.Context, movieID uuid.UUID) (float64, error) {
	query, args, err := squirrel.Select("COALESCE(AVG(rating), 0)").
		From("reviews").
		Where(squirrel.Eq{"movie_id": movieID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "squirrel.ToSql")
	}

	var avgRating float64
	err = repo.conn.QueryRow(ctx, query, args...).Scan(&avgRating)
	if err != nil {
		return 0, errors.Wrap(err, "repo.conn.QueryRow")
	}

	return avgRating, nil
}
