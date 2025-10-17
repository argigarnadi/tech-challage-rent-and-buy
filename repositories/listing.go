package repositories

import (
	"context"
	"errors"
	"fmt"
	"tech-challenge-rent-and-buy/models"

	"gitlab.com/argi.garnadi/go-common/constant/integer"
	"gitlab.com/argi.garnadi/go-common/constant/str"
	"gitlab.com/argi.garnadi/go-common/log"
	"gitlab.com/argi.garnadi/go-common/otel"
	"gorm.io/gorm"
)

type ListingRepository interface {
	GetListing(ctx context.Context, req models.ListingRequest) (data []models.Listing, err error)
	AddedListing(ctx context.Context, req models.Listing) (data models.Listing, err error)
}

type listingRepository struct {
	db *gorm.DB
}

func NewListingRepository(db *gorm.DB) ListingRepository {
	return &listingRepository{db: db}
}

const (
	spanGetListingRepo   = "repositories.listing.GetListing"
	spanAddedListingRepo = "repositories.listing.AddedListing"
)

func (l *listingRepository) GetListing(ctx context.Context, req models.ListingRequest) (data []models.Listing, err error) {
	ctx, span := otel.Trace(ctx, spanGetListingRepo)
	defer span.End()

	var (
		filters []string
		values  []any
	)

	filterCondition := str.Empty

	if req.UserId != integer.Zero {
		filterCondition += "where user_id = ?"
		filters = append(filters, filterCondition)
		values = append(values, req.UserId)
	}

	orderType := "order by created_at desc"

	query := fmt.Sprintf(`select * from listing %s %s LIMIT %d OFFSET %d`, filterCondition, orderType, req.PageSize, (req.PageNum-1)*req.PageSize)
	err = l.db.WithContext(ctx).Raw(query, values...).Scan(&data).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Errorf(ctx, err, "failed get listing")
		return
	}

	if err == gorm.ErrRecordNotFound {
		err = nil
		log.Infof(ctx, "no listing found")
		return
	}

	return
}

func (l *listingRepository) AddedListing(ctx context.Context, req models.Listing) (data models.Listing, err error) {
	ctx, span := otel.Trace(ctx, spanAddedListingRepo)
	defer span.End()

	q := l.db.WithContext(ctx).Create(&req)
	err = q.Error
	if err != nil {
		log.Error(ctx, err)
		return
	}

	q.Scan(&data)
	return
}
