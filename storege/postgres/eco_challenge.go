package postgres

import (
	"context"
	"database/sql"

	exchange "item_service/genproto/item_service"

	"github.com/google/uuid"
)

type EcoChallengeRepo struct {
	db *sql.DB
}

func NewEcoChallengeRepository(db *sql.DB) *EcoChallengeRepo {
	return &EcoChallengeRepo{db}
}

// Eco Challenges
func (r *EcoExchangeRepo) AddEcoChallenge(ctx context.Context, req *exchange.AddEcoChallengeRequest) (*exchange.AddEcoChallengeResponse, error) {
	query := `
		INSERT INTO eco_challenges (id, title, description, start_date, end_date, reward_points)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, title, description, start_date, end_date, reward_points, created_at
	`
	id := uuid.New().String()
	row := r.db.QueryRowContext(ctx, query, id, req.Title, req.Description, req.StartDate, req.EndDate, req.RewardPoints)

	var challenge exchange.EcoChallenge
	err := row.Scan(&challenge.Id, &challenge.Title, &challenge.Description, &challenge.StartDate, &challenge.EndDate, &challenge.RewardPoints, &challenge.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &exchange.AddEcoChallengeResponse{Challenge: &challenge}, nil
}

func (r *EcoExchangeRepo) ParticipateEcoChallenge(ctx context.Context, req *exchange.ParticipateEcoChallengeRequest) (*exchange.ParticipateEcoChallengeResponse, error) {
	query := `
		INSERT INTO eco_challenge_participants (id, challenge_id, user_id, status)
		VALUES ($1, $2, $3, $4)
		returing id, challenge_id, user_id, status, joined_at
	`
	id := uuid.New().String()
	row := r.db.QueryRowContext(ctx, query, req.ChallengeId, id, req.ChallengeId, req.UserId, req.Status)

	var challenge exchange.ParticipateEcoChallengeResponse
	err := row.Scan(&challenge.ChallengeId, &challenge.ChallengeId, &challenge.UserId, &challenge.Status, &challenge.JoinedAt)
	if err != nil {
		return nil, err
	}

	return &exchange.ParticipateEcoChallengeResponse{}, nil
}

func (r *EcoExchangeRepo) UpdateEcoChallengeProgress(ctx context.Context, req *exchange.UpdateEcoChallengeProgressRequest) (*exchange.UpdateEcoChallengeProgressResponse, error) {
	// Implement logic to update eco challenge progress
	// This is a placeholder function for demonstration purposes
	return nil, nil
}
