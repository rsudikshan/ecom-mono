package service

import (
	"context"
	"ecom-mono-go/domain/repository"
	"ecom-mono-go/domain/types"
)

type UserService interface {
	CreateUser(ctx context.Context, user *types.User) (*types.User, error)
	GetUser(ctx context.Context, ID types.ID) (*types.User, error)
	UpdateUser(ctx context.Context, user *types.User) (*types.User, error)
}

type userService struct {
	userRepo repository.UserRepo
}

func NewUserService(userRepo repository.UserRepo) UserService{
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(ctx context.Context, user *types.User) (*types.User, error){
	return s.userRepo.CreateUser(ctx, user)
}

func (s *userService) UpdateUser(ctx context.Context, user *types.User) (*types.User, error){
	return s.userRepo.UpdateUser(ctx, user)
}

func (s *userService) GetUser(ctx context.Context,ID types.ID) (*types.User, error){
	return s.userRepo.GetUser(ctx, ID)
}

