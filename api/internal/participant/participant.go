package participant

import (
	"bytes"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var netClient = &http.Client{
	Timeout: time.Second * 10,
}
var (
	Email_api_key      string = ""
	From_address              = ""
	Twilio_sid         string = ""
	Twilio_api_key     string = ""
	Twilio_account_id  string = ""
	Twilio_from_number string = ""
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
	p.sendText()
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

func (p *Participant) sendText() error {
	data := url.Values{}
	data.Add("To", p.Phone)
	data.Add("From", Twilio_from_number)
	data.Add("Body", p.Message)
	fmt.Println("Sending this data: ", data)
	req, err := http.NewRequest("POST", "https://api.twilio.com/2010-04-01/Accounts/"+Twilio_account_id+"/SMS/Messages.json", bytes.NewBufferString(data.Encode()))
	req.SetBasicAuth(Twilio_sid, Twilio_api_key)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	resp, err := netClient.Do(req)
	if err != nil {
		fmt.Println("Error sending to Twilio: ", err)
		return err
	} else {
		defer resp.Body.Close()
		bts, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("Response Body: ", string(bts))
		return nil
	}
}
