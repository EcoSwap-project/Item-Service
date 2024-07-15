package postgres

import (
	"context"
	"database/sql"
	"item_service/genproto/item_service"
	"time"

	"github.com/google/uuid"
)

type RatingRepo struct {
	db *sql.DB
}

func NewRatingRepo(db *sql.DB) *RatingRepo {
	return &RatingRepo{db}
}

func (r *RatingRepo) AddUserRating(ctx context.Context, req *item_service.AddUserRatingRequest) (*item_service.AddUserRatingResponse, error) {
	id := uuid.NewString()
	createdAt := time.Now().Format(time.RFC3339)

	query := `
		INSERT INTO ratings (id, user_id, rater_id, rating, comment, swap_id, created_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.ExecContext(ctx, query, id, req.UserId, req.RaterId, req.Rating, req.Comment, req.SwapId, createdAt)
	if err != nil {
		return nil, err
	}

	return &item_service.AddUserRatingResponse{
		Rating: &item_service.Rating{
			Id:        id,
			UserId:    req.UserId,
			RaterId:   req.RaterId,
			Rating:    req.Rating,
			Comment:   req.Comment,
			SwapId:    req.SwapId,
			CreatedAt: createdAt,
		},
	}, nil
}

func (r *RatingRepo) GetUserRatings(ctx context.Context, req *item_service.GetUserRatingsRequest) (*item_service.GetUserRatingsResponse, error) {
	query := `
		SELECT id, user_id, rater_id, rating, comment, swap_id, created_at 
		FROM ratings 
		WHERE user_id = $1 
		LIMIT $2 OFFSET $3
	`
	rows, err := r.db.QueryContext(ctx, query, req.UserId, req.Limit, req.Page*req.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ratings []*item_service.Rating
	for rows.Next() {
		var rating item_service.Rating
		err := rows.Scan(&rating.Id, &rating.UserId, &rating.RaterId, &rating.Rating, &rating.Comment, &rating.SwapId, &rating.CreatedAt)
		if err != nil {
			return nil, err
		}
		ratings = append(ratings, &rating)
	}

	return &item_service.GetUserRatingsResponse{
		Ratings: ratings,
		Page:    req.Page,
		Limit:   req.Limit,
	}, nil
}
