package auth

import (
	//"ecom-mono-go/api/base"
	"ecom-mono-go/api/base"
	"ecom-mono-go/api/dtos"
	"ecom-mono-go/domain/service"
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
	var registerParam dtos.UserRegisterParams
	err := ctx.ShouldBindJSON(&registerParam)

	if err!=nil {
		h.HandleError(ctx,err)
		return
	}

	
}