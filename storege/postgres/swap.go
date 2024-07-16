package postgres

import (
	"context"
	"database/sql"
	"item_service/genproto/item_service"
	"time"

	"github.com/google/uuid"
)

type SwapRepo struct {
	db *sql.DB
}

func NewSwapRepo(db *sql.DB) *SwapRepo {
	return &SwapRepo{db}
}

func (r *SwapRepo) CreateSwap(ctx context.Context, req *item_service.CreateSwapRequest) (*item_service.CreateSwapResponse, error) {
	id := uuid.NewString()
	createdAt := time.Now().Format(time.RFC3339)
	updatedAt := createdAt

	query := `
		INSERT INTO swaps (id, offered_item_id, requested_item_id, requester_id, owner_id, status, message) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := r.db.ExecContext(ctx, query, id, req.OfferedItemId, req.RequestedItemId, req.RequesterId, req.OwnerId, "pending", req.Message, createdAt, updatedAt)
	if err != nil {
		return nil, err
	}

	return &item_service.CreateSwapResponse{
		Swap: &item_service.Swap{
			Id:              id,
			OfferedItemId:   req.OfferedItemId,
			RequestedItemId: req.RequestedItemId,
			RequesterId:     req.RequesterId,
			OwnerId:         req.OwnerId,
			Status:          "pending",
			Message:         req.Message,
			CreatedAt:       createdAt,
			UpdatedAt:       updatedAt,
		},
	}, nil
}

func (r *SwapRepo) UpdateSwapStatus(ctx context.Context, req *item_service.UpdateSwapStatusRequest) (*item_service.UpdateSwapStatusResponse, error) {
	updatedAt := time.Now().Format(time.RFC3339)
	query := `
		UPDATE swaps 
		SET status = $1, message = $2, updated_at = $3 
		WHERE id = $4 and deleted_at is null
		RETURNING id, offered_item_id, requested_item_id, requester_id, owner_id, status, message, created_at, updated_at
	`
	var swap item_service.Swap
	err := r.db.QueryRowContext(ctx, query, req.Status, req.Reason, updatedAt, req.Id).
		Scan(&swap.Id, &swap.OfferedItemId, &swap.RequestedItemId, &swap.RequesterId, &swap.OwnerId, &swap.Status, &swap.Message, &swap.CreatedAt, &swap.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &item_service.UpdateSwapStatusResponse{
		Swap: &swap,
	}, nil
}

func (r *SwapRepo) GetSwaps(ctx context.Context, req *item_service.GetSwapsRequest) (*item_service.GetSwapsResponse, error) {
	query := `
		SELECT id, offered_item_id, requested_item_id, requester_id, owner_id, status, message, created_at, updated_at 
		FROM swaps 
		WHERE status = $1 and deleted_at is null
		LIMIT $2 OFFSET $3
	`
	rows, err := r.db.QueryContext(ctx, query, req.Status, req.Limit, req.Page*req.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var swaps []*item_service.Swap
	for rows.Next() {
		var swap item_service.Swap
		err := rows.Scan(&swap.Id, &swap.OfferedItemId, &swap.RequestedItemId, &swap.RequesterId, &swap.OwnerId, &swap.Status, &swap.Message, &swap.CreatedAt, &swap.UpdatedAt)
		if err != nil {
			return nil, err
		}
		swaps = append(swaps, &swap)
	}

	return &item_service.GetSwapsResponse{
		Swaps: swaps,
		Page:  req.Page,
		Limit: req.Limit,
	}, nil
}
