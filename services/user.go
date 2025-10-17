package services

import (
	"context"
	"tech-challenge-rent-and-buy/models"
	"tech-challenge-rent-and-buy/repositories"

	"gitlab.com/argi.garnadi/go-common/log"
	"gitlab.com/argi.garnadi/go-common/otel"
)

type UserService interface {
	GetUser(ctx context.Context, req models.UsersRequest) (res models.UsersResponse, err error)
	AddUser(ctx context.Context, req models.AddUserRequest) (res models.AddUserResponse, err error)
	GetUserById(ctx context.Context, id int) (res models.AddUserResponse, err error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

const (
	spanGetUsers    = "app.service.GetUser"
	spanAddUser     = "app.service.AddUser"
	spanGetUserById = "app.service.AddUser"
)

func (u *userService) GetUser(ctx context.Context, req models.UsersRequest) (res models.UsersResponse, err error) {
	ctx, span := otel.Trace(ctx, spanGetUsers)
	defer span.End()

	data, err := u.userRepo.GetListUsers(ctx, req)
	if err != nil {
		log.Error(ctx, err)
		res.Result = false
		return
	}

	res.Result = true
	res.Users = data
	return
}

func (u *userService) AddUser(ctx context.Context, req models.AddUserRequest) (res models.AddUserResponse, err error) {
	ctx, span := otel.Trace(ctx, spanAddUser)
	defer span.End()

	payload := models.User{
		Name: req.Name,
	}

	data, err := u.userRepo.AddUser(ctx, payload)
	if err != nil {
		log.Error(ctx, err)
		res.Result = false
		return
	}

	res.Result = true
	res.User = data
	return
}

func (u *userService) GetUserById(ctx context.Context, id int) (res models.AddUserResponse, err error) {
	ctx, span := otel.Trace(ctx, spanGetUserById)
	defer span.End()

	data, err := u.userRepo.GetUserById(ctx, id)
	if err != nil {
		log.Error(ctx, err)
		res.Result = false
		return
	}

	res.Result = true
	res.User = data
	log.Info(ctx, "user found")
	return
}
