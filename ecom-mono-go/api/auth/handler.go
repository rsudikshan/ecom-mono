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
	GetEmailVerificationLink(ctx *gin.Context)
}

type authHandler struct {
	userService service.UserService
	authService service.AuthService
	*base.Handler
}

func NewAuthHandler(h *base.Handler, userService service.UserService, authService service.AuthService) AuthHandler {
	return &authHandler{
		userService: userService,
		authService: authService,
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

	m, err := h.userService.CreateUser(ctx, user)

	if err!=nil {
		h.HandleError(ctx, err)
		return
	}


	err = h.authService.SendEmailVerificationToken(ctx, m.Email)

	if err!=nil {
		h.HandleError(ctx, err)
		return
	}

	h.JSON(ctx, http.StatusOK, "Registration succsessful. Email Verification link has been sent to your email. Procced to verify email.")
}

func (h *authHandler) GetEmailVerificationLink(ctx *gin.Context) {
	
}