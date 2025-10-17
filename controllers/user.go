package controllers

import (
	"errors"
	"strconv"
	"tech-challenge-rent-and-buy/models"
	"tech-challenge-rent-and-buy/services"

	"github.com/gin-gonic/gin"
	"gitlab.com/argi.garnadi/go-common/constant/integer"
	"gitlab.com/argi.garnadi/go-common/constant/str"
	"gitlab.com/argi.garnadi/go-common/log"
	"gitlab.com/argi.garnadi/go-common/response"
	"gitlab.com/argi.garnadi/go-common/transport/restserver"
)

type UserController interface {
	GetUser(c *gin.Context)
	AddUser(c *gin.Context)
	GetUserById(c *gin.Context)
}

type userController struct {
	userServ services.UserService
}

func NewUserController(userServ services.UserService) UserController {
	return &userController{userServ: userServ}
}

func (u *userController) GetUser(c *gin.Context) {
	ctx := c.Request.Context()
	var req models.UsersRequest
	if err := c.BindQuery(&req); err != nil {
		log.Error(ctx, err)
		restserver.SetResponse(c.Writer, response.InvalidArgument)
		return
	}

	res, err := u.userServ.GetUser(ctx, req)
	if err != nil {
		log.Error(ctx, err)
		restserver.SetResponse(c.Writer, response.GeneralError)
		return
	}

	if len(res.Users) == integer.Zero {
		log.Info(ctx, "no user found")
		restserver.SetResponse(c.Writer, response.DataNotFound)
		return
	}

	restserver.SetResponse(c.Writer, response.Success)
	c.JSON(response.Success.HttpStatusCode(), res)
}

func (u *userController) AddUser(c *gin.Context) {
	ctx := c.Request.Context()
	var req models.AddUserRequest
	if err := c.BindQuery(&req); err != nil {
		log.Error(ctx, err)
		restserver.SetResponse(c.Writer, response.InvalidArgument)
		return
	}

	res, err := u.userServ.AddUser(ctx, req)
	if err != nil {
		log.Error(ctx, err)
		restserver.SetResponse(c.Writer, response.GeneralError)
		return
	}

	restserver.SetResponse(c.Writer, response.Success)
	c.JSON(response.Success.HttpStatusCode(), res)
}

func (u *userController) GetUserById(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	if id == str.Empty {
		err := errors.New("id is empty")
		log.Error(ctx, err)
		restserver.SetResponse(c.Writer, response.InvalidArgument)
		return
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Error(ctx, err)
		restserver.SetResponse(c.Writer, response.InvalidArgument)
		return
	}

	res, err := u.userServ.GetUserById(ctx, intId)
	if err != nil {
		log.Error(ctx, err)
		restserver.SetResponse(c.Writer, response.GeneralError)
		return
	}

	restserver.SetResponse(c.Writer, response.Success)
	c.JSON(response.Success.HttpStatusCode(), res)
}
