package postgres

import (
	"context"
	"database/sql"

	exchange "item_service/genproto/item_service"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type ItemRepo struct {
	db *sql.DB
}

func NewItemRepo(db *sql.DB) *ItemRepo {
	return &ItemRepo{db}
}

// Items
func (r *ItemRepo) AddItem(ctx context.Context, req *exchange.AddItemRequest) (*exchange.AddItemResponse, error) {
	query := `
		INSERT INTO items (id, name, description, category_id, condition, swap_preference, images, owner_id, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, now(), now())
		RETURNING id, name, description, category_id, condition, swap_preference, owner_id, status, created_at, updated_at
	`
	id := uuid.New().String()
	row := r.db.QueryRowContext(ctx, query, id, req.Name, req.Description, req.CategoryId, req.Condition, pq.Array(req.SwapPreference), pq.Array(req.Images), "", "active")

	var item exchange.Item
	err := row.Scan(&item.Id, &item.Name, &item.Description, &item.CategoryId, &item.Condition, pq.Array(&item.SwapPreference), &item.OwnerId, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &exchange.AddItemResponse{Item: &item}, nil
}

func (r *ItemRepo) UpdateItem(ctx context.Context, req *exchange.UpdateItemRequest) (*exchange.UpdateItemResponse, error) {
	query := `
		UPDATE items
		SET name = $1, condition = $2, updated_at = now()
		WHERE id = $3 AND deleted_at is null
		RETURNING id, name, description, category_id, condition, swap_preference, owner_id, status, created_at, updated_at
	`
	row := r.db.QueryRowContext(ctx, query, req.Name, req.Condition, req.Id)

	var item exchange.Item
	err := row.Scan(&item.Id, &item.Name, &item.Description, &item.CategoryId, &item.Condition, pq.Array(&item.SwapPreference), &item.OwnerId, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &exchange.UpdateItemResponse{Item: &item}, nil
}

func (r *ItemRepo) DeleteItem(ctx context.Context, req *exchange.DeleteItemRequest) (*exchange.DeleteItemResponse, error) {
	query := `
		UPDATE items
		SET deleted_at = now()
		WHERE id = $1 AND deleted_at is null
	`
	_, err := r.db.ExecContext(ctx, query, req.Id)
	if err != nil {
		return nil, err
	}

	return &exchange.DeleteItemResponse{Message: "Item deleted successfully"}, nil
}

func (r *ItemRepo) GetItems(ctx context.Context, req *exchange.GetItemsRequest) (*exchange.GetItemsResponse, error) {
	query := `
		SELECT id, name, description, category_id, condition, swap_preference, owner_id, status, created_at, updated_at
		FROM items WHERE deleted_at is null
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.QueryContext(ctx, query, req.Limit, (req.Page-1)*req.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*exchange.Item
	for rows.Next() {
		var item exchange.Item
		err := rows.Scan(&item.Id, &item.Name, &item.Description, &item.CategoryId, &item.Condition, pq.Array(&item.SwapPreference), &item.OwnerId, &item.Status, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	var total int32
	err = r.db.QueryRowContext(ctx,  "SELECT COUNT(*) FROM items WHERE deleted_at is null").Scan(&total)
	if err != nil {
		return nil, err
	}

	return &exchange.GetItemsResponse{
		Items: items,
		Total: total,
		Page:  req.Page,
		Limit: req.Limit,
	}, nil
}

func (r *ItemRepo) GetItem(ctx context.Context, req *exchange.GetItemRequest) (*exchange.GetItemResponse, error) {
	query := `
		SELECT id, name, description, category_id, condition, swap_preference, owner_id, status, created_at, updated_at
		FROM items
		WHERE id = $1 and deleted_at is null		
	`
	row := r.db.QueryRowContext(ctx, query, req.Id)

	var item exchange.Item
	err := row.Scan(&item.Id, &item.Name, &item.Description, &item.CategoryId, &item.Condition, pq.Array(&item.SwapPreference), &item.OwnerId, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &exchange.GetItemResponse{Item: &item}, nil
}

func (r *ItemRepo) SearchItems(ctx context.Context, req *exchange.SearchItemsRequest) (*exchange.SearchItemsResponse, error) {
	query := `
		SELECT id, name, description, category_id, condition, swap_preference, owner_id, status, created_at, updated_at
		FROM items
		WHERE name ILIKE '%' || $1 || '%' AND category_id = $2 AND condition = $3 and deleted_at is null
		LIMIT $4 OFFSET $5
	`
	rows, err := r.db.QueryContext(ctx, query, req.Query, req.Category, req.Condition, req.Limit, (req.Page-1)*req.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*exchange.Item
	for rows.Next() {
		var item exchange.Item
		err := rows.Scan(&item.Id, &item.Name, &item.Description, &item.CategoryId, &item.Condition, pq.Array(&item.SwapPreference), &item.OwnerId, &item.Status, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	var total int32
	err = r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM items WHERE name ILIKE '%' || $1 || '%' AND category_id = $2 AND condition = $3 and deleted_at is null", req.Query, req.Category, req.Condition).Scan(&total)
	if err != nil {
		return nil, err
	}

	return &exchange.SearchItemsResponse{
		Items: items,
		Total: total,
		Page:  req.Page,
		Limit: req.Limit,
	}, nil
}
