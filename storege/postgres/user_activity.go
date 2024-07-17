package postgres

import (
	"context"
	"database/sql"

	exchange "item_service/genproto/item_service"
)

type userActivityRepo struct {
	db *sql.DB
}

func NewUserActivityRepository(db *sql.DB) *userActivityRepo {
	return &userActivityRepo{db}
}

// User Activity
func (r *userActivityRepo) GetUserActivity(ctx context.Context, req *exchange.GetUserActivityRequest) (*exchange.GetUserActivityResponse, error) {
	
	var swapInitiated, swapCompleted, itemListed, recyclingSubmissions, ecoPointsEarned int32
	var userId string

	err := r.db.QueryRow(`SELECT COUNT(*) FROM swaps WHERE user_id = $1 AND created_at BETWEEN $2 AND $3`, req.UserId, req.StartDate, req.EndDate).Scan(&swapInitiated)
	if err != nil {
		return nil, err
	}

	err = r.db.QueryRow(`SELECT COUNT(*) FROM swaps WHERE user_id = $1 AND completed_at BETWEEN $2 AND $3`, req.UserId, req.StartDate, req.EndDate).Scan(&swapCompleted)
	if err != nil {
		return nil, err
	}

	err = r.db.QueryRow(`SELECT COUNT(*) FROM items WHERE user_id = $1 AND listed_at BETWEEN $2 AND $3`, req.UserId, req.StartDate, req.EndDate).Scan(&itemListed)
	if err != nil {
		return nil, err
	}

	err = r.db.QueryRow(`SELECT COUNT(*) FROM recycled_items WHERE user_id = $1 AND submitted_at BETWEEN $2 AND $3`, req.UserId, req.StartDate, req.EndDate).Scan(&recyclingSubmissions)
	if err != nil {
		return nil, err
	}

	err = r.db.QueryRow(`SELECT SUM(eco_points) FROM eco_points WHERE user_id = $1 AND earned_at BETWEEN $2 AND $3`, req.UserId, req.StartDate, req.EndDate).Scan(&ecoPointsEarned)
	if err != nil {
		return nil, err
	}

	response := &exchange.GetUserActivityResponse{
		UserId:               userId,
		SwapsInitiated: swapInitiated,
		SwapsCompleted: swapCompleted,
		ItemsListed:          itemListed,
		RecyclingSubmissions: recyclingSubmissions,
		EcoPointsEarned:      ecoPointsEarned,
	}

	return response, nil
}
