package postgres

import (
	"context"
	"database/sql"
	"item_service/genproto/item_service"
	"time"

	"github.com/google/uuid"
)

type RecyclingCenterRepo struct {
	db *sql.DB
}

func NewRecyclingCenterRepo(db *sql.DB) *RecyclingCenterRepo {
	return &RecyclingCenterRepo{db}
}

func (r *RecyclingCenterRepo) AddRecyclingCenter(ctx context.Context, req *item_service.AddRecyclingCenterRequest) (*item_service.AddRecyclingCenterResponse, error) {
	id := uuid.NewString()
	createdAt := time.Now().Format(time.RFC3339)
	updatedAt := createdAt

	query := `
		INSERT INTO recycling_centers (id, name, address, accepted_materials, working_hours, contact_number, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := r.db.ExecContext(ctx, query, id, req.Name, req.Address, req.AcceptedMaterials, req.WorkingHours, req.ContactNumber, createdAt, updatedAt)
	if err != nil {
		return nil, err
	}

	return &item_service.AddRecyclingCenterResponse{
		Center: &item_service.RecyclingCenter{
			Id:               id,
			Name:             req.Name,
			Address:          req.Address,
			AcceptedMaterials: req.AcceptedMaterials,
			WorkingHours:     req.WorkingHours,
			ContactNumber:    req.ContactNumber,
			CreatedAt:        createdAt,
			UpdatedAt:        updatedAt,
		},
	}, nil
}

func (r *RecyclingCenterRepo) GetRecyclingCenters(ctx context.Context, req *item_service.GetRecyclingCentersRequest) (*item_service.GetRecyclingCentersResponse, error) {
	query := `
		SELECT id, name, address, accepted_materials, working_hours, contact_number, created_at, updated_at 
		FROM recycling_centers 
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.QueryContext(ctx, query, req.Limit, req.Page*req.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var centers []*item_service.RecyclingCenter
	for rows.Next() {
		var center item_service.RecyclingCenter
		err := rows.Scan(&center.Id, &center.Name, &center.Address, &center.AcceptedMaterials, &center.WorkingHours, &center.ContactNumber, &center.CreatedAt, &center.UpdatedAt)
		if err != nil {
			return nil, err
		}
		centers = append(centers, &center)
	}

	return &item_service.GetRecyclingCentersResponse{
		Centers: centers,
		Page:    req.Page,
		Limit:   req.Limit,
	}, nil
}
