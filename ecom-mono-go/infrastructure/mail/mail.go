package mail

import (
	"context"
	"ecom-mono-go/infrastructure"
	"strconv"
	"gopkg.in/gomail.v2"
)

type EmailParams struct {
	To	 	   []string
	Body 	   string
}

type MailSender interface {
	SendEmail(ctx context.Context, params *EmailParams ) error
}

type mailSender struct {
	env *infrastructure.Env
	dialer *gomail.Dialer
}

func NewMailSender(env *infrastructure.Env) MailSender{
	i,_ := strconv.Atoi((env.EMAIL_PORT))
	return &mailSender{
		env: env,
		dialer: gomail.NewDialer(env.EMAIL_SENDER_URL,i, env.EMAIL_SENDER, env.EMAIL_APP_PASSWORD),
	}
}

func (m *mailSender) SendEmail(ctx context.Context, params *EmailParams) error {
	message := gomail.NewMessage()
	message.SetHeader("From", m.env.EMAIL_SENDER)
	message.SetHeader("To", params.To...)
	message.SetBody("text/html", params.Body)

	return m.dialer.DialAndSend(message)
}