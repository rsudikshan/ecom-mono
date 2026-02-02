package service

import (
	"context"
	"ecom-mono-go/domain/types"
	"ecom-mono-go/infrastructure"
	"ecom-mono-go/infrastructure/auth_utils"
	"ecom-mono-go/infrastructure/mail"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	SendEmailVerificationToken(ctx context.Context,id types.ID, recepient ...string) error
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


func (as *authService) SendEmailVerificationToken(ctx context.Context,id types.ID, recepient ...string) error{
	token,err:=
	auth_utils.CreateToken(
		as.env.AUTH_TOKEN_KEY_SECRET, 
		jwt.MapClaims{
			"id":id.String(),
			"exp":time.Hour*time.Duration(as.env.EMAIL_VERIFICATION_TOKEN_VALIDITY_TIME),
			"type":auth_utils.EMAIL_VERIFICATION_TOKEN,
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