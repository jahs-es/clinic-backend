package mailer

type EmailResponse struct {
	Status int
	RespBody string
}

type Message struct {
	To string
	Subject string
	Intro string
	Body string
	Params map[string]string
}

type IMailer interface {
	Send(message Message) (*EmailResponse, error)
}




