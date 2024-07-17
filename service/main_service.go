package services

import (
	"database/sql"
	user "item_service/genproto/authentication_service"
	eco "item_service/genproto/item_service"
	repo "item_service/storege/postgres"
)

type MainService interface {
	ItemService() ItemService
	SwapService() SwapService
	RecyclingCenterService() RecyclingCenterService
	UserRatingService() UserRatingService
	ItemCategoryService() ItemCategoryService
	StatisticsService() StatisticsService
	UserActivityService() UserActivityService
	EcoChallengeService() EcoChallengeService
	EcoTipService() EcoTipService
}

type mainServiceImpl struct {
	eco.UnimplementedEcoExchangeServiceServer
	authClient             user.EcoServiceClient
	itemService            ItemService
	swapService            SwapService
	recyclingCenterService RecyclingCenterService
	userRatingService      UserRatingService
	itemCategoryService    ItemCategoryService
	statisticsService      StatisticsService
	userActivityService    UserActivityService
	ecoChallengeService    EcoChallengeService
	ecoTipService          EcoTipService
}

func NewMainService(db *sql.DB, userClient user.EcoServiceClient) *mainServiceImpl {
	return &mainServiceImpl{
		authClient:             userClient,
		itemService:            NewItemService(repo.NewItemRepo(db)),
		swapService:            NewSwapService(repo.NewSwapRepo(db)),
		recyclingCenterService: NewRecyclingCenterService(repo.NewRecyclingCenterRepository(db)),
		userRatingService:      NewUserRatingService(repo.NewRatingRepo(db)),
		itemCategoryService:    NewItemCategoryService(repo.NewItemCategoryRepository(db)),
		statisticsService:      NewStatisticsService(repo.NewStatisticsRepository(db)),
		userActivityService:    NewUserActivityService(repo.NewUserActivityRepository(db)),
		ecoChallengeService:    NewEcoChallengeService(repo.NewEcoChallengeRepository(db)),
		ecoTipService:          NewEcoTipService(repo.NewEcoTipRepository(db)),
	}
}

func (s *mainServiceImpl) ItemService() ItemService {
	return s.itemService
}

func (s *mainServiceImpl) SwapService() SwapService {
	return s.swapService
}

func (s *mainServiceImpl) RecyclingCenterService() RecyclingCenterService {
	return s.recyclingCenterService
}

func (s *mainServiceImpl) UserRatingService() UserRatingService {
	return s.userRatingService
}

func (s *mainServiceImpl) ItemCategoryService() ItemCategoryService {
	return s.itemCategoryService
}

func (s *mainServiceImpl) StatisticsService() StatisticsService {
	return s.statisticsService
}

func (s *mainServiceImpl) UserActivityService() UserActivityService {
	return s.userActivityService
}

func (s *mainServiceImpl) EcoChallengeService() EcoChallengeService {
	return s.ecoChallengeService
}

func (s *mainServiceImpl) EcoTipService() EcoTipService {
	return s.ecoTipService
}
