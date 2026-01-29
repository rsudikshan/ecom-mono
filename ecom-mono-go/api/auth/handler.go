package auth

import (
	//"ecom-mono-go/api/base"
	"ecom-mono-go/api/base"
	"ecom-mono-go/api/dtos"
	"ecom-mono-go/domain/service"
	"ecom-mono-go/domain/types"
	"net/http"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Signup(ctx *gin.Context)
}

type authHandler struct {
	userService service.UserService
	*base.Handler
}

func NewAuthHandler(h *base.Handler, userService service.UserService) AuthHandler {
	return &authHandler{
		userService: userService,
	}
}

func (h *authHandler) Signup(ctx *gin.Context) {
	var registerParam dtos.UserSignUpParams
	err := ctx.ShouldBindJSON(&registerParam)

	if err!=nil {
		h.HandleError(ctx,err)
		return
	}

	user := &types.User{
		ID: types.NewID(),
		Username: registerParam.Username,
		Password: registerParam.Password,
		Email: registerParam.Email,
	}

	//user,err := h.userService.CreateUser(ctx, user)
	_,err = h.userService.CreateUser(ctx, user)

	if err!=nil {
		h.HandleError(ctx, err)
		return
	}

	h.JSON(ctx, http.StatusOK, "Registration succsseful. Email Verification link has been sent to your email. Procced to verify email.")
}