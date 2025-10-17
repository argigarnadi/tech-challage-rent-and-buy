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

type PublicApiController interface {
	GetListing(c *gin.Context)
	AddUser(c *gin.Context)
	AddListing(c *gin.Context)
}

type publicApiController struct {
	publicApiServ services.PublicApiService
	userServ      services.UserService
	listingServ   services.ListingService
}

func NewPublicApiController(publicApiServ services.PublicApiService,
	userServ services.UserService,
	listingServ services.ListingService,
) PublicApiController {
	return &publicApiController{
		publicApiServ: publicApiServ,
		userServ:      userServ,
		listingServ:   listingServ,
	}
}

func (p *publicApiController) GetListing(c *gin.Context) {
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

	res, err := p.publicApiServ.GetListing(ctx, req)
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

func (p *publicApiController) AddUser(c *gin.Context) {
	ctx := c.Request.Context()
	var req models.AddUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(ctx, err)
		restserver.SetResponse(c.Writer, response.InvalidArgument)
		return
	}

	data, err := p.userServ.AddUser(ctx, req)
	if err != nil {
		log.Error(ctx, err)
		restserver.SetResponse(c.Writer, response.GeneralError)
		return
	}

	restserver.SetResponse(c.Writer, response.Success)
	c.JSON(response.Success.HttpStatusCode(), data)
}

func (p *publicApiController) AddListing(c *gin.Context) {
	ctx := c.Request.Context()
	var req models.AddListingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(ctx, err)
		restserver.SetResponse(c.Writer, response.InvalidArgument)
		return
	}

	data, err := p.listingServ.AddListing(ctx, req)
	if err != nil {
		log.Error(ctx, err)
		restserver.SetResponse(c.Writer, response.GeneralError)
		return
	}

	restserver.SetResponse(c.Writer, response.Success)
	c.JSON(response.Success.HttpStatusCode(), data)
}
