package auth

import "github.com/gin-gonic/gin"

type AuthRoutes struct {
	handler AuthHandler
	rg      *gin.RouterGroup
}

func NewAuthRoutes(handler AuthHandler, rg *gin.RouterGroup) *AuthRoutes {
	return &AuthRoutes{
		handler: handler,
		rg: 	 rg.Group("auth"),	
	}
}

func (r *AuthRoutes) Setup() {
	r.rg.POST("/register", r.handler.Signup)
}