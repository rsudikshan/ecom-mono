package service

import (
	"context"
	"ecom-mono-go/infrastructure"
	"ecom-mono-go/infrastructure/auth"
	"ecom-mono-go/infrastructure/mail"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	SendEmailVerificationToken(ctx context.Context, recepient ...string) error
}
type authService struct {
	env *infrastructure.Env
	mailSender mail.MailSender
}

func NewAuthService(env *infrastructure.Env, mailSender mail.MailSender) AuthService {
	return &authService{
		env: env,
		mailSender: mailSender,
	}
}


func (as *authService) SendEmailVerificationToken(ctx context.Context, recepient ...string) error{
	token,err:=
	auth.CreateToken(
		as.env.AUTH_TOKEN_KEY_SECRET, 
		jwt.MapClaims{
			"exp":time.Hour*time.Duration(as.env.EMAIL_VERIFICATION_TOKEN_VALIDITY_TIME),
		}, 
	)
	if err!=nil {
		return err
	}

	return as.mailSender.SendEmail(ctx, &mail.EmailParams{
		To: recepient,
		Body: fmt.Sprintf(`
			<div>
				<b>Hi please proceed to this url to verify your account %s </b>
			</div>
		`, token) ,
	})
}