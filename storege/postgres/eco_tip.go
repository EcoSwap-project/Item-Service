package postgres

import (
	"context"
	"database/sql"

	exchange "item_service/genproto/item_service"

	"github.com/google/uuid"
)

type EcoTipRepo struct {
	db *sql.DB
}

func NewEcoTipRepository(db *sql.DB) *EcoTipRepo {
	return &EcoTipRepo{db}
}

// Eco Tips
func (r *EcoTipRepo) AddEcoTip(ctx context.Context, req *exchange.AddEcoTipRequest) (*exchange.AddEcoTipResponse, error) {
	query := `
		INSERT INTO eco_tips (id, title, content, created_at, updated_at)
		VALUES ($1, $2, $3, now(), now())
		RETURNING id, title, content, created_at
	`
	id := uuid.New().String()
	row := r.db.QueryRowContext(ctx, query, id, req.Title, req.Content)

	var tip exchange.EcoTip
	err := row.Scan(&tip.Id, &tip.Title, &tip.Content, &tip.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &exchange.AddEcoTipResponse{Tip: &tip}, nil
}

func (r *EcoTipRepo) GetEcoTips(ctx context.Context, req *exchange.GetEcoTipsRequest) (*exchange.GetEcoTipsResponse, error) {
	query := `
		SELECT id, title, content, created_at, updated_at
		FROM eco_tips
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.QueryContext(ctx, query, req.Limit, (req.Page-1)*req.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tips []*exchange.EcoTip
	for rows.Next() {
		var tip exchange.EcoTip
		err := rows.Scan(&tip.Id, &tip.Title, &tip.Content, &tip.CreatedAt)
		if err != nil {
			return nil, err
		}
		tips = append(tips, &tip)
	}

	var total int32
	err = r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM eco_tips").Scan(&total)
	if err != nil {
		return nil, err
	}

	return &exchange.GetEcoTipsResponse{
		Tips:  tips,
		Total: total,
		Page:  req.Page,
		Limit: req.Limit,
	}, nil
}
