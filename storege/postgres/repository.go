package postgres

import (
	"context"

	exchange "item_service/genproto/item_service"
)

// Items interface
type ItemRepository interface {
	AddItem(ctx context.Context, req *exchange.AddItemRequest) (*exchange.AddItemResponse, error)
	UpdateItem(ctx context.Context, req *exchange.UpdateItemRequest) (*exchange.UpdateItemResponse, error)
	DeleteItem(ctx context.Context, req *exchange.DeleteItemRequest) (*exchange.DeleteItemResponse, error)
	GetItems(ctx context.Context, req *exchange.GetItemsRequest) (*exchange.GetItemsResponse, error)
	GetItem(ctx context.Context, req *exchange.GetItemRequest) (*exchange.GetItemResponse, error)
	SearchItems(ctx context.Context, req *exchange.SearchItemsRequest) (*exchange.SearchItemsResponse, error)
}

// Swaps interface
type SwapRepository interface {
	CreateSwap(ctx context.Context, req *exchange.CreateSwapRequest) (*exchange.CreateSwapResponse, error)
	UpdateSwapStatus(ctx context.Context, req *exchange.UpdateSwapStatusRequest) (*exchange.UpdateSwapStatusResponse, error)
	GetSwaps(ctx context.Context, req *exchange.GetSwapsRequest) (*exchange.GetSwapsResponse, error)
}

// Recycling Centers interface
type RecyclingCenterRepository interface {
	AddRecyclingCenter(ctx context.Context, req *exchange.AddRecyclingCenterRequest) (*exchange.AddRecyclingCenterResponse, error)
	GetRecyclingCenters(ctx context.Context, req *exchange.GetRecyclingCentersRequest) (*exchange.GetRecyclingCentersResponse, error)
	AddRecyclingSubmission(ctx context.Context, req *exchange.AddRecyclingSubmissionRequest) (*exchange.AddRecyclingSubmissionResponse, error)
}

// User Ratings interface
type UserRatingRepository interface {
	AddUserRating(ctx context.Context, req *exchange.AddUserRatingRequest) (*exchange.AddUserRatingResponse, error)
	GetUserRatings(ctx context.Context, req *exchange.GetUserRatingsRequest) (*exchange.GetUserRatingsResponse, error)
}

// Item Categories interface
type ItemCategoryRepository interface {
	AddItemCategory(ctx context.Context, req *exchange.AddItemCategoryRequest) (*exchange.AddItemCategoryResponse, error)
}

// Statistics interface
type StatisticsRepository interface {
	GetStatistics(ctx context.Context, req *exchange.GetStatisticsRequest) (*exchange.GetStatisticsResponse, error)
}

// User Activity interface
type UserActivityRepository interface {
	GetUserActivity(ctx context.Context, req *exchange.GetUserActivityRequest) (*exchange.GetUserActivityResponse, error)
}

// Eco Challenges interface
type EcoChallengeRepository interface {
	AddEcoChallenge(ctx context.Context, req *exchange.AddEcoChallengeRequest) (*exchange.AddEcoChallengeResponse, error)
	ParticipateEcoChallenge(ctx context.Context, req *exchange.ParticipateEcoChallengeRequest) (*exchange.ParticipateEcoChallengeResponse, error)
	UpdateEcoChallengeProgress(ctx context.Context, req *exchange.UpdateEcoChallengeProgressRequest) (*exchange.UpdateEcoChallengeProgressResponse, error)
}

// Eco Tips interface
type EcoTipRepository interface {
	AddEcoTip(ctx context.Context, req *exchange.AddEcoTipRequest) (*exchange.AddEcoTipResponse, error)
	GetEcoTips(ctx context.Context, req *exchange.GetEcoTipsRequest) (*exchange.GetEcoTipsResponse, error)
}
