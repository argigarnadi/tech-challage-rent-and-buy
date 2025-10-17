package main

import (
	"context"
	"tech-challenge-rent-and-buy/controllers"
	"tech-challenge-rent-and-buy/models"
	"tech-challenge-rent-and-buy/repositories"
	"tech-challenge-rent-and-buy/routers"
	"tech-challenge-rent-and-buy/services"

	"gitlab.com/argi.garnadi/go-common/app"
	"gitlab.com/argi.garnadi/go-common/database"
	"gitlab.com/argi.garnadi/go-common/log"
)

func main() {
	ctx := context.Background()
	// Get db client config
	dbId := "99-test"
	gSql, _, err := database.Manager.DB(ctx, dbId)
	if err != nil {
		log.Fatalf(ctx, err, "failed initialize db:%s", dbId)
	}

	gSql.AutoMigrate(&models.Listing{}, &models.User{})

	// repository declare interface
	listingRepo := repositories.NewListingRepository(gSql)
	userRepo := repositories.NewUserRepository(gSql)

	// service declare interface
	listingServ := services.NewListingService(listingRepo)
	userService := services.NewUserService(userRepo)
	publicApiService := services.NewPublicApiService(userRepo, listingRepo)

	// handler declare interface
	listingCtrl := controllers.NewListingController(listingServ)
	userCtrl := controllers.NewUserController(userService)
	publicApiCtrl := controllers.NewPublicApiController(publicApiService, userService, listingServ)

	routers.Init(ctx, listingCtrl, userCtrl, publicApiCtrl)
	app.Run()
}
