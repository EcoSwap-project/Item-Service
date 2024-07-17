package postgres

import (
	"context"
	"database/sql"

	exchange "item_service/genproto/item_service"

	"github.com/google/uuid"
)

type itemCategoryRepo struct {
	db *sql.DB
}

func NewItemCategoryRepository(db *sql.DB) *itemCategoryRepo {
	return &itemCategoryRepo{db}
}

// Item Categories
func (r *itemCategoryRepo) AddItemCategory(ctx context.Context, req *exchange.AddItemCategoryRequest) (*exchange.AddItemCategoryResponse, error) {
	query := `
		INSERT INTO item_categories (id, name, description)
		VALUES ($1, $2, $3)
		RETURNING id, name, description, created_at
	`
	id := uuid.New().String()
	row := r.db.QueryRowContext(ctx, query, id, req.Name, req.Description)

	var category exchange.ItemCategory
	err := row.Scan(&category.Id, &category.Name, &category.Description, &category.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &exchange.AddItemCategoryResponse{Category: &category}, nil
}
