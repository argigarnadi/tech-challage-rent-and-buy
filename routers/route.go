package routers

import (
	"context"
	"tech-challenge-rent-and-buy/controllers"
	"tech-challenge-rent-and-buy/routers/listing"
	public_api "tech-challenge-rent-and-buy/routers/public-api"
	"tech-challenge-rent-and-buy/routers/user"

	"gitlab.com/argi.garnadi/go-common/transport/restserver"
)

func Init(ctx context.Context,
	listingCtrl controllers.ListingController,
	userCtrl controllers.UserController,
	publicApiCtrl controllers.PublicApiController,
) {
	v1 := restserver.Router.Group("/v1")
	{
		listingApi := v1.Group("/listings")
		listing.ListingRoute(ctx, listingApi, listingCtrl)

		userApi := v1.Group("/users")
		user.UserRoute(ctx, userApi, userCtrl)

		publicApi := v1.Group("/public-api")
		public_api.PublicApiRoute(ctx, publicApi, publicApiCtrl)
	}
}
