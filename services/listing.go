package services

import (
	"context"
	"tech-challenge-rent-and-buy/models"
	"tech-challenge-rent-and-buy/repositories"

	"gitlab.com/argi.garnadi/go-common/log"
	"gitlab.com/argi.garnadi/go-common/otel"
)

type ListingService interface {
	GetListing(ctx context.Context, req models.ListingRequest) (res models.ListingResponse, err error)
	AddListing(ctx context.Context, req models.AddListingRequest) (res models.AddListingResponse, err error)
}

type listingService struct {
	listingRepo repositories.ListingRepository
}

func NewListingService(listingRepo repositories.ListingRepository) ListingService {
	return &listingService{listingRepo: listingRepo}
}

const (
	spanGetListing = "services.listing.GetListing"
	spanAddListing = "services.listing.AddListing"
)

func (l *listingService) GetListing(ctx context.Context, req models.ListingRequest) (res models.ListingResponse, err error) {
	ctx, span := otel.Trace(ctx, spanGetListing)
	defer span.End()

	data, err := l.listingRepo.GetListing(ctx, req)
	if err != nil {
		log.Error(ctx, err)
		res.Result = false
		return
	}

	res.Result = true
	res.Listings = data
	return
}

func (l *listingService) AddListing(ctx context.Context, req models.AddListingRequest) (res models.AddListingResponse, err error) {
	ctx, span := otel.Trace(ctx, spanAddListing)
	defer span.End()

	payload := models.Listing{
		UserId:      req.UserId,
		Price:       req.Price,
		ListingType: req.ListingType,
	}

	data, err := l.listingRepo.AddedListing(ctx, payload)
	if err != nil {
		log.Error(ctx, err)
		res.Result = false
		return
	}

	res.Result = true
	res.Listing = data
	return
}
