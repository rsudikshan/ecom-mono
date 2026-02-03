package auth

import (
	"ecom-mono-go/api/base"
	"ecom-mono-go/api/dtos"
	"ecom-mono-go/api/middleware"
	"ecom-mono-go/domain/service"
	"ecom-mono-go/domain/types"
	"ecom-mono-go/infrastructure"
	"ecom-mono-go/utils/apperror"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Signup(ctx *gin.Context)
	Signin(ctx *gin.Context)
	GetEmailVerificationLink(ctx *gin.Context)
	VerifyEmail(ctx *gin.Context)
}

type authHandler struct {
	userService service.UserService
	authService service.AuthService
	am middleware.AuthMiddleware
	env *infrastructure.Env
	*base.Handler
}

func NewAuthHandler(h *base.Handler, 
	userService service.UserService, 
	authService service.AuthService, 
	am middleware.AuthMiddleware,
	env *infrastructure.Env,) AuthHandler {
	return &authHandler{
		userService: userService,
		authService: authService,
		Handler: h,
		am:am,
		env: env,
	}
}

func (h *authHandler) Signup(ctx *gin.Context) {
	var registerParam dtos.UserSignUpParams
	err := ctx.ShouldBindJSON(&registerParam)

	if err!=nil {
		h.HandleError(ctx,err)
		return
	}

	id := types.NewID()
	user := &types.User{
		ID: id,
		Username: registerParam.Username,
		Password: registerParam.Password,
		Email: registerParam.Email,
		Base: &types.Base{
			CreatedByID : id,
			UpdatedByID	: id,
		},
	}

	m, err := h.userService.CreateUser(ctx, user)

	if err!=nil {
		h.HandleError(ctx, err)
		return
	}


	err = h.authService.SendEmailVerificationToken(ctx, m.ID, m.Email)

	if err!=nil {
		h.HandleError(ctx, err)
		return
	}

	h.JSON(ctx, http.StatusOK, "Registration succsessful. Email Verification link has been sent to your email. Procced to verify email.")
}

func (h *authHandler) Signin(ctx *gin.Context) {
	var loginParam dtos.UserLoginParams
	err := ctx.ShouldBindJSON(&loginParam)
	if err!=nil {
		h.HandleError(ctx, err)
		return
	}

	user,err := h.userService.GetUserByEmail(ctx, loginParam.Email)

	if err!=nil {
		h.HandleError(ctx, fmt.Errorf("user not found"))
		return
	}

	if !user.EmailVerified {
		h.HandleError(ctx, fmt.Errorf("email is not verified, please verify the email to be able to login"))
		return
	}
	
	var r dtos.UserLoginResponse
	token, err := h.authService.LoginUser(ctx, user, loginParam.Password)
	
	if err!=nil {
		h.HandleError(ctx, apperror.New(http.StatusUnauthorized, err))
		return
	}

	r.Token = token
	h.JSON(ctx, http.StatusOK, r)
}

func (h *authHandler) GetEmailVerificationLink(ctx *gin.Context) {
	
}

func(h *authHandler) VerifyEmail(ctx *gin.Context) {
	emailVerificationToken := ctx.Query("token")
	userID, err := h.am.HandleEmailVerification(emailVerificationToken)
	if err!=nil {
		h.HandleError(ctx,err)
		return
	}

	user,err := h.userService.GetUser(ctx, *userID)

	if err!=nil {
		h.HandleError(ctx,err)
		return
	}

	user.EmailVerified = true
	_, err = h.userService.UpdateUser(ctx, user)

	if err!=nil {
		h.HandleError(ctx,err)
		return
	}

	h.JSON(ctx,http.StatusOK, "user verified successfully, procced to login")
}