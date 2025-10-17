package controllers

import (
	"tech-challenge-rent-and-buy/models"
	"tech-challenge-rent-and-buy/services"

	"github.com/gin-gonic/gin"
	"gitlab.com/argi.garnadi/go-common/constant/integer"
	"gitlab.com/argi.garnadi/go-common/log"
	"gitlab.com/argi.garnadi/go-common/response"
	"gitlab.com/argi.garnadi/go-common/transport/restserver"
)

type ListingController interface {
	GetListing(c *gin.Context)
	AddListing(c *gin.Context)
}

type listingController struct {
	listingServ services.ListingService
}

func NewListingController(listingServ services.ListingService) ListingController {
	return &listingController{listingServ: listingServ}
}

func (l *listingController) GetListing(c *gin.Context) {
	ctx := c.Request.Context()
	var req models.ListingRequest
	if err := c.BindQuery(&req); err != nil {
		log.Error(ctx, err)
		restserver.SetResponse(c.Writer, response.InvalidArgument)
		return
	}

	if req.PageSize == integer.Zero {
		req.PageSize = 10
	}

	if req.PageNum == integer.Zero {
		req.PageNum = 1
	}

	res, err := l.listingServ.GetListing(ctx, req)
	if err != nil {
		log.Error(ctx, err)
		restserver.SetResponse(c.Writer, response.GeneralError)
		return
	}

	if len(res.Listings) == integer.Zero {
		log.Info(ctx, "no listing found")
		restserver.SetResponse(c.Writer, response.DataNotFound)
		return
	}

	restserver.SetResponse(c.Writer, response.Success)
	c.JSON(response.Success.HttpStatusCode(), res)
}

func (l *listingController) AddListing(c *gin.Context) {
	ctx := c.Request.Context()
	var req models.AddListingRequest
	if err := c.BindQuery(&req); err != nil {
		log.Error(ctx, err)
		restserver.SetResponse(c.Writer, response.InvalidArgument)
		return
	}

	res, err := l.listingServ.AddListing(ctx, req)
	if err != nil {
		log.Error(ctx, err)
		restserver.SetResponse(c.Writer, response.GeneralError)
		return
	}
	restserver.SetResponse(c.Writer, response.Success)
	c.JSON(response.Success.HttpStatusCode(), res)
}
