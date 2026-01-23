package repository

import (
	"context"
	"ecom-mono-go/domain/types"

	"gorm.io/gorm"
)

type UserRepo interface {
	SaveUser(ctx context.Context, user *types.User)(*types.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo{
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) SaveUser(ctx context.Context, user *types.User)(*types.User, error){
	if err:=r.db.Create(user).Error; err!=nil{
		return nil,err
	}

	return user,nil
}