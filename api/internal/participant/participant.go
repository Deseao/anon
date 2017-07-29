package participant

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
)

var (
	Email_api_key string = ""
	From_address         = ""
)

type Participant struct {
	Email   string
	Phone   string
	Message string
}

func NewParticipant(email string, phone string) *Participant {
	return &Participant{Email: email, Phone: phone}
}

func (p *Participant) SendMessage(message string) {
	p.Message = message
	p.sendEmail()
}

func (p *Participant) sendEmail() error {
	from := mail.NewEmail("Anonymous", From_address)
	subject := "New Message from Anonymous"
	to := mail.NewEmail("Anonymous User", p.Email)
	message := mail.NewSingleEmail(from, subject, to, p.Message, p.Message)
	client := sendgrid.NewSendClient(Email_api_key)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return err
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
		return nil
	}
}
