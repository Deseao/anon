package participant

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
}
