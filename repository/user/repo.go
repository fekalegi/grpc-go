package user

import (
	"context"
	"gorm.io/gorm"
	"grpc-go/domain"
)

type RepoUser struct {
	db *gorm.DB
}

type RepoUserContract interface {
	CreateUser(ctx context.Context, request *domain.User) error
	GetListUser(ctx context.Context, limit, offset int) ([]domain.User, int64, error)
	GetUser(ctx context.Context, id int) (*domain.User, error)
	UpdateUser(ctx context.Context, request *domain.User) error
	DeleteUser(ctx context.Context, request int) error
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
}

func NewUserRepo(db *gorm.DB) *RepoUser {
	return &RepoUser{
		db: db,
	}
}

func (r *RepoUser) CreateUser(ctx context.Context, request *domain.User) error {
	return r.db.Create(&request).Error
}

func (r *RepoUser) GetListUser(ctx context.Context, limit, offset int) ([]domain.User, int64, error) {
	var res []domain.User
	var count int64
	err := r.db.Offset(offset).Limit(limit).Find(&res).Offset(-1).Limit(-1).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	return res, count, nil
}

func (r *RepoUser) GetUser(ctx context.Context, request int) (*domain.User, error) {
	var res *domain.User

	err := r.db.Where("id = ?", request).First(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *RepoUser) UpdateUser(ctx context.Context, request *domain.User) error {
	err := r.db.Where("email = ?", request.Email).Updates(&request).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *RepoUser) DeleteUser(ctx context.Context, request int) error {
	return r.db.Where("id = ?", request).Delete(&domain.User{}).Error
}

func (r *RepoUser) GetUserByEmail(ctx context.Context, request string) (*domain.User, error) {
	var res *domain.User

	err := r.db.Where("email = ?", request).First(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}
