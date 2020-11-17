package mailer

import (
	"net/http"
)

type FakeMailer struct{}

func NewFakeMailer() *FakeMailer {
	return &FakeMailer{}
}

func (p *FakeMailer) Send(to string, subject string, body string) (*EmailResponse, error) {
	return &EmailResponse{
		Status:   http.StatusOK,
		RespBody: "Success, email was sent",
	}, nil
}
