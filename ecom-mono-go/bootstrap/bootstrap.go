package bootstrap

import (
	"ecom-mono-go/api/auth"
	"ecom-mono-go/api/base"
	"ecom-mono-go/domain/repository"
	"ecom-mono-go/domain/service"
	"ecom-mono-go/infrastructure"
	"ecom-mono-go/infrastructure/mail"
)

func Run() {
	env := infrastructure.NewEnv()
	db := infrastructure.NewDb(env)
	infrastructure.Migrate(db)
	router := infrastructure.NewAppRouter(env)
	mailSender := mail.NewMailSender(env)
	
	userRepo := repository.NewUserRepo(db)

	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(env, mailSender)

	baseHandler := base.NewBaseHandler()
	authHandler := auth.NewAuthHandler(baseHandler, userService, authService)

	auth.NewAuthRoutes(authHandler,router.RG).Setup()

	router.Start()
}