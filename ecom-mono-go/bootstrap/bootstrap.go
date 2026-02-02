package bootstrap

import (
	"ecom-mono-go/api/auth"
	"ecom-mono-go/api/base"
	"ecom-mono-go/api/middleware"
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

	authMiddleware := middleware.NewAuthMiddleware(env, baseHandler, )

	authHandler := auth.NewAuthHandler(baseHandler, userService, authService, authMiddleware, env)

	auth.NewAuthRoutes(authHandler,router.RG, authMiddleware).Setup()

	router.Start()
}