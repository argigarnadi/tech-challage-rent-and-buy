package listing

import (
	"context"
	"tech-challenge-rent-and-buy/controllers"

	"github.com/gin-gonic/gin"
)

func ListingRoute(ctx context.Context, group *gin.RouterGroup, listingCtrl controllers.ListingController) {
	group.GET("/", listingCtrl.GetListing)
	group.POST("/", listingCtrl.AddListing)
}
