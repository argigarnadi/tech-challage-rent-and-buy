package user

import (
	"context"
	"tech-challenge-rent-and-buy/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(ctx context.Context, group *gin.RouterGroup, userCtrl controllers.UserController) {
	group.GET("/", userCtrl.GetUser)
	group.GET("/:id", userCtrl.GetUserById)
	group.POST("/", userCtrl.AddUser)
}
