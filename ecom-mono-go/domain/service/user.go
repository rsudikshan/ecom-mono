package service

import (
	"context"
	"ecom-mono-go/domain/repository"
	"ecom-mono-go/domain/types"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(ctx context.Context, user *types.User) (*types.User, error)
	GetUser(ctx context.Context, ID types.ID) (*types.User, error)
	GetUserByEmail(ctx context.Context, email string) (*types.User, error)
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
	hash,_ := bcrypt.GenerateFromPassword([]byte(user.Password),12)
	user.Password = string(hash)
	return s.userRepo.CreateUser(ctx, user)
}

func (s *userService) UpdateUser(ctx context.Context, user *types.User) (*types.User, error){
	return s.userRepo.UpdateUser(ctx, user)
}

func (s *userService) GetUser(ctx context.Context,ID types.ID) (*types.User, error){
	return s.userRepo.GetUser(ctx, ID)
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	return s.userRepo.GetUserByEmail(ctx, email)
}