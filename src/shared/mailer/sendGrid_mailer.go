package mailer

import (
	"github.com/jahs/clinic-backend/src/application/config"
	"github.com/matcornic/hermes/v2"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"net/http"
	"strings"
)

type SendGridMailer struct{}

func NewSendGridMailer() *SendGridMailer {
	return &SendGridMailer{}
}

func (p *SendGridMailer) Send(message Message) (*EmailResponse, error) {

	h := hermes.Hermes{
		Product: hermes.Product{
			Name: "Clinic Go",
			Link: "http://a1d5af7.online-server.cloud/login",
		},
	}
	email := hermes.Email{
		Body: hermes.Body{
			Name: message.To,
			Intros: []string{
				message.Intro,
			},
			Outros: []string{
				p.replaceParams(message.Params, message.Body),
			},
			Signature: "Siempre agradecidos",
		},
	}

	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		return nil, err
	}

	fromMail := mail.NewEmail("Clinic Go", config.SENDGRID_FROM)

	toMail := mail.NewEmail("Creación de usuario", message.To)

	sendGridMessage := mail.NewSingleEmail(fromMail, message.Subject, toMail, emailBody, emailBody)

	client := sendgrid.NewSendClient(config.SENDGRID_KEY)

	_, err = client.Send(sendGridMessage)

	return &EmailResponse{
		Status:   http.StatusOK,
		RespBody: "Success",
	}, nil
}

func (p *SendGridMailer) replaceParams(params map[string]string, body string) string {
	for param, value := range params {
		body = strings.ReplaceAll(body, param, value)
	}
	return body
}
