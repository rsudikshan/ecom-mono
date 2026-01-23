package service

import (
	"context"
	"ecom-mono-go/domain/repository"
	"ecom-mono-go/domain/types"
)

type UserService interface {
	SaveUser(ctx context.Context, user *types.User) (*types.User, error)
}

type userService struct {
	userRepo repository.UserRepo
}

func NewUserService(userRepo repository.UserRepo) UserService{
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) SaveUser(ctx context.Context, user *types.User) (*types.User, error){
	return s.userRepo.SaveUser(ctx, user)
}