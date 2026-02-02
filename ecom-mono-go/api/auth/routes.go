package auth

import (
	"ecom-mono-go/api/middleware"
	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	handler AuthHandler
	rg      *gin.RouterGroup
	am 		middleware.AuthMiddleware
}

func NewAuthRoutes(handler AuthHandler, rg *gin.RouterGroup, am middleware.AuthMiddleware) *AuthRoutes {
	return &AuthRoutes{
		handler: handler,
		rg: 	 rg.Group("auth"),	
		am: am,
	}
}

func (r *AuthRoutes) Setup() {
	r.rg.POST("/register", r.handler.Signup)
	r.rg.POST("/email-verify", r.handler.VerifyEmail)
	// TODO
	// reset-password
}