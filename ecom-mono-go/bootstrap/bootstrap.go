package bootstrap

import (
	"ecom-mono-go/api/auth"
	"ecom-mono-go/api/base"
	"ecom-mono-go/domain/repository"
	"ecom-mono-go/domain/service"
	"ecom-mono-go/infrastructure"
)

func Run() {
	env := infrastructure.NewEnv()
	db := infrastructure.NewDb(env)
	router := infrastructure.NewAppRouter(env)
	
	userRepo := repository.NewUserRepo(db)

	userService := service.NewUserService(userRepo)

	baseHandler := base.NewBaseHandler()
	authHandler := auth.NewAuthHandler(baseHandler,userService)

	auth.NewAuthRoutes(authHandler,router.RG).Setup()

	router.Start()
}