package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"item_service/config/logger"
	exchange "item_service/genproto/item_service"
)

type statisticsRepo struct {
	db *sql.DB
	lg *slog.Logger
}

func NewStatisticsRepository(db *sql.DB) *statisticsRepo {
	return &statisticsRepo{db: db, lg: logger.NewLogger()}
}

// Statistics
func (r *statisticsRepo) GetStatistics(ctx context.Context, req *exchange.GetStatisticsRequest) (*exchange.GetStatisticsResponse, error) {
	startDate := req.StartDate
	endDate := req.EndDate

	var totalSwaps, totalRecycledItems, totalEcoPointsEarned int32

	err := r.db.QueryRow(`SELECT COUNT(*) FROM swaps WHERE created_at BETWEEN $1 AND $2`, startDate, endDate).Scan(&totalSwaps)
	if err != nil {

		r.lg.Error(fmt.Sprintf("message get total swaps error -> %v", err))
		return nil, err
	}

	err = r.db.QueryRow(`SELECT COUNT(*) FROM recycled_items WHERE created_at BETWEEN $1 AND $2`, startDate, endDate).Scan(&totalRecycledItems)
	if err != nil {
		r.lg.Error(fmt.Sprintf("message get total recycled items error -> %v", err))
		return nil, err
	}

	err = r.db.QueryRow(`SELECT SUM(eco_points) FROM eco_points WHERE created_at BETWEEN $1 AND $2`, startDate, endDate).Scan(&totalEcoPointsEarned)
	if err != nil {
		r.lg.Error(fmt.Sprintf("message get total eco points earned error -> %v", err))
		return nil, err
	}

	rows, err := r.db.Query(`SELECT id, name, COUNT(*) as swap_count FROM item_categories
	JOIN swaps ON item_categories.id = swaps.category_id WHERE swaps.created_at BETWEEN $1 AND $2
	GROUP BY id, name ORDER BY swap_count DESC LIMIT 10`, startDate, endDate)
	if err != nil {
		r.lg.Error(fmt.Sprintf("message get top categories error -> %v", err))
		return nil, err
	}
	defer rows.Close()

	var topCategories []*exchange.TopCategory
	for rows.Next() {
		var category exchange.TopCategory
		err := rows.Scan(&category.Id, &category.Name, &category.SwapCount)
		if err != nil {
			r.lg.Error(fmt.Sprintf("message scan category error -> %v", err))
			return nil, err
		}
		topCategories = append(topCategories, &category)
	}

	rows, err = r.db.Query(`SELECT id, name, COUNT(*) as submissions_count FROM recycling_centers
	JOIN recycled_items ON recycling_centers.id = recycled_items.center_id WHERE recycled_items.created_at BETWEEN $1 AND $2
	GROUP BY id, name ORDER BY submissions_count DESC LIMIT 10`, startDate, endDate)
	if err != nil {
		r.lg.Error(fmt.Sprintf("message get top recycling centers error -> %v", err))
		return nil, err
	}
	defer rows.Close()

	var topRecyclingCenters []*exchange.TopRecyclingCenter
	for rows.Next() {
		var center exchange.TopRecyclingCenter
		err := rows.Scan(&center.Id, &center.Name, &center.SubmissionsCount)
		if err != nil {
			r.lg.Error(fmt.Sprintf("message scan center error -> %v", err))
			return nil, err
		}
		topRecyclingCenters = append(topRecyclingCenters, &center)
	}

	response := &exchange.GetStatisticsResponse{
		TotalSwaps:           totalSwaps,
		TotalRecycledItems:   totalRecycledItems,
		TotalEcoPointsEarned: totalEcoPointsEarned,
		TopCategories:        topCategories,
		TopRecyclingCenters:  topRecyclingCenters,
	}

	return response, nil
}
