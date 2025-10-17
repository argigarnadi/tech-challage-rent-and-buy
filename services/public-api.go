package services

import (
	"context"
	"tech-challenge-rent-and-buy/models"
	"tech-challenge-rent-and-buy/repositories"

	"gitlab.com/argi.garnadi/go-common/log"
	"gitlab.com/argi.garnadi/go-common/otel"
)

type PublicApiService interface {
	GetListing(ctx context.Context, req models.ListingRequest) (res models.PublicApiResponse, err error)
}

type publicApiService struct {
	userRepo    repositories.UserRepository
	listingRepo repositories.ListingRepository
}

func NewPublicApiService(
	userRepo repositories.UserRepository,
	listingRepo repositories.ListingRepository,
) PublicApiService {
	return &publicApiService{
		userRepo:    userRepo,
		listingRepo: listingRepo,
	}
}

const (
	spanPublicApiGetListing = "app.public-api.GetListing"
)

func (p *publicApiService) GetListing(ctx context.Context, req models.ListingRequest) (res models.PublicApiResponse, err error) {
	ctx, span := otel.Trace(ctx, spanPublicApiGetListing)
	defer span.End()

	// get data listing
	data, err := p.listingRepo.GetListing(ctx, req)
	if err != nil {
		log.Error(ctx, err)
		res.Result = false
		return
	}

	var listings []models.Listings
	for _, v := range data {
		userData, errUserData := p.userRepo.GetUserById(ctx, v.UserId)
		if errUserData != nil {
			log.Error(ctx, errUserData)
			res.Result = false
			return
		}

		listings = append(listings, models.Listings{
			Id:          v.Id,
			ListingType: v.ListingType,
			Price:       v.Price,
			CreatedAt:   v.CreatedAt,
			UpdateAt:    v.UpdateAt,
			User:        userData,
		})
	}

	res.Result = true
	res.Listings = listings
	return
}
