package repository

import (
	"context"
	"ecom-mono-go/domain/types"
	"ecom-mono-go/utils/apperror"
	"fmt"
	"net/http"
	"strings"
	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *types.User)(*types.User, error)
	GetUser(ctx context.Context, id types.ID) (*types.User, error)
	UpdateUser(ctx context.Context, user *types.User) (*types.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo{
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) CreateUser(ctx context.Context, user *types.User)(*types.User, error){
	err:=r.db.WithContext(ctx).
	Create(user).
	Error
	if err!=nil {
		if IsDuplicateError(err){
			switch{
				case strings.Contains(err.Error(), "users_email_key"):
					return nil, apperror.New(http.StatusBadRequest, fmt.Errorf("the user with %s email is already registered:",user.Email))
				
				default :
					return nil, err
			}
		}
		return nil,err
	}
	return user,nil
}

func (r *userRepo) UpdateUser(ctx context.Context, user *types.User)(*types.User, error){
	if err:=r.db.WithContext(ctx).
	Model(types.User{}).
	Where("id = ?", user.ID).
	Updates(map[string]any{
		"id":user.ID,
		"username":user.Username,
		"password":user.Password,
		"email_verified":user.EmailVerified,
		"password_reset_at":user.PasswordResetAt,
	}).Error; err!=nil {
		return nil,err
	}
	return user,nil
}

func (r *userRepo) GetUser(ctx context.Context, ID types.ID)(*types.User, error){
	var result types.User
	if err:=r.db.WithContext(ctx).Where("id = ?", ID).Find(&result).Error; err!=nil{
		return nil,err
	}
	return &result,nil
}


