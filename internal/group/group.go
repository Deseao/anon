package group

import (
	"github.com/Deseao/anon/internal/code"
	"github.com/Deseao/anon/internal/participant"
)

type Group struct {
	Code         string
	Participants []*participant.Participant
}

func NewGroup() *Group {
	return &Group{Code: code.GenRandCode(code.CODE_LEN)}
}

func (g *Group) AddParticipant(email string, phone string) {
	newParticipant := participant.NewParticipant(email, phone)
	g.Participants = append(g.Participants, newParticipant)
}

func (g *Group) SendMessage(message string) {
	for _, v := range g.Participants {
		v.SendMessage(message)
	}
}
