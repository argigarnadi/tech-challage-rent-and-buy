package repositories

import (
	"context"
	"errors"
	"tech-challenge-rent-and-buy/models"

	"gitlab.com/argi.garnadi/go-common/log"
	"gitlab.com/argi.garnadi/go-common/otel"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetListUsers(ctx context.Context, req models.UsersRequest) (data []models.User, err error)
	AddUser(ctx context.Context, req models.User) (data models.User, err error)
	GetUserById(ctx context.Context, id int) (data models.User, err error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

const (
	spanGetListUsersRepo = "app.repositories.GetListUser"
	spanAddUser          = "app.repositories.AddUser"
	spanGetUserById      = "app.repositories.GetUserById"
)

func (u *userRepository) GetListUsers(ctx context.Context, req models.UsersRequest) (data []models.User, err error) {
	ctx, span := otel.Trace(ctx, spanGetListUsersRepo)
	defer span.End()

	err = u.db.WithContext(ctx).Find(&data).Order("created_at desc").Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error(ctx, err)
		return
	}

	if err == gorm.ErrRecordNotFound {
		err = nil
		log.Info(ctx, "no user found")
		return
	}

	return
}

func (u *userRepository) AddUser(ctx context.Context, req models.User) (data models.User, err error) {
	ctx, span := otel.Trace(ctx, spanAddUser)
	defer span.End()

	q := u.db.WithContext(ctx).Create(&req)
	err = q.Error
	if err != nil {
		log.Error(ctx, err)
		return
	}

	q.Scan(&data)
	return
}

func (u *userRepository) GetUserById(ctx context.Context, id int) (data models.User, err error) {
	ctx, span := otel.Trace(ctx, spanGetUserById)
	defer span.End()

	q := u.db.WithContext(ctx).Where("id = ?", id).First(&data)
	err = q.Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error(ctx, err)
		return
	}

	if err == gorm.ErrRecordNotFound {
		err = nil
		log.Info(ctx, "no user found")
		return
	}
	log.Info(ctx, "user found")
	return
}
