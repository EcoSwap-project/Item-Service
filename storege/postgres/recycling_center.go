package postgres

import (
	"context"
	"database/sql"

	exchange "item_service/genproto/item_service"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type RecyclingCenterRepo struct{
	db *sql.DB
}

func NewRecyclingCenterRepository(db *sql.DB) *RecyclingCenterRepo {
	return &RecyclingCenterRepo{db}
}


// Recycling Centers
func (r *EcoExchangeRepo) AddRecyclingCenter(ctx context.Context, req *exchange.AddRecyclingCenterRequest) (*exchange.AddRecyclingCenterResponse, error) {
	query := `
		INSERT INTO recycling_centers (id, name, address, accepted_materials, working_hours, contact_number)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, name, address, accepted_materials, working_hours, contact_number
	`
	id := uuid.New().String()
	row := r.db.QueryRowContext(ctx, query, id, req.Name, req.Address, req.AcceptedMaterials, req.WorkingHours, req.ContactNumber)

	var center exchange.RecyclingCenter
	err := row.Scan(&center.Id, &center.Name, &center.Address, &center.AcceptedMaterials, &center.WorkingHours,  &center.ContactNumber)
	if err != nil {
		return nil, err
	}

	return &exchange.AddRecyclingCenterResponse{Center: &center}, nil
}

func (r *EcoExchangeRepo) GetRecyclingCenters(ctx context.Context, req *exchange.GetRecyclingCentersRequest) (*exchange.GetRecyclingCentersResponse, error) {
	query := `
		SELECT id, name, address, contact_info, created_at, updated_at
		FROM recycling_centers
		where deleted_at is null
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.QueryContext(ctx, query, req.Limit, (req.Page-1)*req.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var centers []*exchange.RecyclingCenter
	for rows.Next() {
		var center exchange.RecyclingCenter
		err := rows.Scan(&center.Id, &center.Name, &center.Address, &center.ContactNumber, &center.CreatedAt, &center.UpdatedAt)
		if err != nil {
			return nil, err
		}
		centers = append(centers, &center)
	}

	var total int32
	err = r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM recycling_centers where deleted_at is null").Scan(&total)
	if err != nil {
		return nil, err
	}

	return &exchange.GetRecyclingCentersResponse{
		Centers: centers,
		Total:   total,
		Page:    req.Page,
		Limit:   req.Limit,
	}, nil
}

func (r *EcoExchangeRepo) AddRecyclingSubmission(ctx context.Context, req *exchange.AddRecyclingSubmissionRequest) (*exchange.AddRecyclingSubmissionResponse, error) {
	query := `
		INSERT INTO recycling_submissions (id, user_id, center_id, items, eco_points_earned)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, user_id, center_id, items, eco_points_earned, created_at
	`
	id := uuid.New().String()
	row := r.db.QueryRowContext(ctx, query, id, req.UserId, req.CenterId, pq.Array(req.Items), 0 )

	var submission exchange.RecyclingSubmission
	err := row.Scan(&submission.Id, &submission.UserId, &submission.CenterId, pq.Array(&submission.Items), &submission.EcoPointsEarned, &submission.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &exchange.AddRecyclingSubmissionResponse{Submission: &submission}, nil
}
