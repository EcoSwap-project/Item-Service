package postgres

import (
	"context"
	"database/sql"
	"item_service/genproto/item_service"
	"time"

	"github.com/google/uuid"
)

type ItemRepo struct {
	db *sql.DB
}

func NewItemRepo(db *sql.DB) *ItemRepo {
	return &ItemRepo{db}
}

func (r *ItemRepo) AddItem(ctx context.Context, req *item_service.AddItemRequest) (*item_service.AddItemResponse, error) {
	id := uuid.NewString()
	createdAt := time.Now().Format(time.RFC3339)
	updatedAt := createdAt

	query := `
		INSERT INTO items (id, name, description, category_id, condition, owner_id, status, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := r.db.ExecContext(ctx, query, id, req.Name, req.Description, req.CategoryId, req.Condition, "", "available", createdAt, updatedAt)
	if err != nil {
		return nil, err
	}

	return &item_service.AddItemResponse{
		Item: &item_service.Item{
			Id:           id,
			Name:         req.Name,
			Description:  req.Description,
			CategoryId:   req.CategoryId,
			Condition:    req.Condition,
			Status:       "available",
			CreatedAt:    createdAt,
			UpdatedAt:    updatedAt,
		},
	}, nil
}

func (r *ItemRepo) UpdateItem(ctx context.Context, req *item_service.UpdateItemRequest) (*item_service.UpdateItemResponse, error) {
	updatedAt := time.Now().Format(time.RFC3339)
	query := `
		UPDATE items 
		SET name = $1, condition = $2, updated_at = $3 
		WHERE id = $4
		RETURNING id, name, description, category_id, condition, owner_id, status, created_at, updated_at
	`
	var item item_service.Item
	err := r.db.QueryRowContext(ctx, query, req.Name, req.Condition, updatedAt, req.Id).
		Scan(&item.Id, &item.Name, &item.Description, &item.CategoryId, &item.Condition, &item.OwnerId, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &item_service.UpdateItemResponse{
		Item: &item,
	}, nil
}

func (r *ItemRepo) DeleteItem(ctx context.Context, req *item_service.DeleteItemRequest) (*item_service.DeleteItemResponse, error) {
	query := `DELETE FROM items WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, req.Id)
	if err != nil {
		return nil, err
	}

	return &item_service.DeleteItemResponse{
		Message: "Item deleted successfully",
	}, nil
}

func (r *ItemRepo) GetItems(ctx context.Context, req *item_service.GetItemsRequest) (*item_service.GetItemsResponse, error) {
	query := `
		SELECT id, name, description, category_id, condition, owner_id, status, created_at, updated_at 
		FROM items 
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.QueryContext(ctx, query, req.Limit, req.Page*req.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*item_service.Item
	for rows.Next() {
		var item item_service.Item
		err := rows.Scan(&item.Id, &item.Name, &item.Description, &item.CategoryId, &item.Condition, &item.OwnerId, &item.Status, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	return &item_service.GetItemsResponse{
		Items: items,
		Page:  req.Page,
		Limit: req.Limit,
	}, nil
}

func (r *ItemRepo) GetItem(ctx context.Context, req *item_service.GetItemRequest) (*item_service.GetItemResponse, error) {
	query := `
		SELECT id, name, description, category_id, condition, owner_id, status, created_at, updated_at 
		FROM items 
		WHERE id = $1
	`
	var item item_service.Item
	err := r.db.QueryRowContext(ctx, query, req.Id).
		Scan(&item.Id, &item.Name, &item.Description, &item.CategoryId, &item.Condition, &item.OwnerId, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &item_service.GetItemResponse{
		Item: &item,
	}, nil
}
