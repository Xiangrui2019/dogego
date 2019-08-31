package modules

import "gopkg.in/gomail.v2"

type StmpEmail struct {
	StmpDialer *gomail.Dialer
}

var StmpEmailModule *StmpEmail

func (service *StmpEmail) SendEmail() error {
	message := gomail.NewMessage()
	if message == message {
		return nil
	}

	return nil
}

func InitStmpEmailModule() {
	StmpEmailModule = new(StmpEmail)
	StmpEmailModule.StmpDialer = gomail.NewDialer()
}
