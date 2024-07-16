package services

import (
	"context"

	eco "item_service/genproto/item_service"
	repo "item_service/storege/postgres"
)

type EcoChallengeService interface {
	AddEcoChallenge(context.Context, *eco.AddEcoChallengeRequest) (*eco.AddEcoChallengeResponse, error)
	GetEcoChallenges(context.Context, *eco.ParticipateEcoChallengeRequest) (*eco.ParticipateEcoChallengeResponse, error)
}

type ecoChallengeServiceImpl struct {
	repo repo.EcoChallengeRepository
}

func NewEcoChallengeService(repo repo.EcoChallengeRepository) EcoChallengeService {
	return &ecoChallengeServiceImpl{repo: repo}
}

func (s *ecoChallengeServiceImpl) AddEcoChallenge(ctx context.Context, req *eco.AddEcoChallengeRequest) (*eco.AddEcoChallengeResponse, error) {
	challenge, err := s.repo.AddEcoChallenge(ctx, req)
	if err != nil {
		return nil, err
	}
	return challenge, nil
}

func (s *ecoChallengeServiceImpl) GetEcoChallenges(ctx context.Context, req *eco.ParticipateEcoChallengeRequest) (*eco.ParticipateEcoChallengeResponse, error) {
	challenges,  err := s.repo.ParticipateEcoChallenge(ctx, req)
	if err != nil {
		return nil, err
	}
	return challenges, nil
}
