package participant

type Participant struct {
	Email string
	Phone string
}

func NewParticipant(email string, phone string) *Participant {
	return &Participant{Email: email, Phone: phone}
}
