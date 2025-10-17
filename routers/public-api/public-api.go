package public_api

import (
	"context"
	"tech-challenge-rent-and-buy/controllers"

	"github.com/gin-gonic/gin"
)

func PublicApiRoute(ctx context.Context, group *gin.RouterGroup, publicApiCtrl controllers.PublicApiController) {
	group.GET("/listings", publicApiCtrl.GetListing)
	group.POST("/users", publicApiCtrl.AddUser)
	group.POST("/listings", publicApiCtrl.AddListing)
}
