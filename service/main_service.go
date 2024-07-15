package services

import (
	"database/sql"
	"item_service/genproto/authentication_service"

	repository "item_service/storege/postgres"
)

type MainService interface {
	ItemService() ItemService
	SwapService() SwapService
	RecyclingCenterService() RecyclingCenterService
	RatingService() RatingService
}

type mainServiceImpl struct {
	authClient             authentication_service.EcoServiceClient
	itemService            ItemService
	swapService            SwapService
	recyclingCenterService RecyclingCenterService
	ratingService          RatingService
}

func NewMainService(db *sql.DB, authClient authentication_service.EcoServiceClient) MainService {
	return &mainServiceImpl{
		authClient:             authClient,
		itemService:            NewItemService(repository.NewItemRepo(db)),
		swapService:            NewSwapService(repository.NewSwapRepo(db)),
		recyclingCenterService: NewRecyclingCenterService(repository.NewRecyclingCenterRepo(db)),
		ratingService:          NewRatingService(repository.NewRatingRepo(db)),
	}
}

func (ms *mainServiceImpl) ItemService() ItemService {
	return ms.itemService
}

func (ms *mainServiceImpl) SwapService() SwapService {
	return ms.swapService
}

func (ms *mainServiceImpl) RecyclingCenterService() RecyclingCenterService {
	return ms.recyclingCenterService
}

func (ms *mainServiceImpl) RatingService() RatingService {
	return ms.ratingService
}
