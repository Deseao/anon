package group

import (
	"errors"
)

type GroupHandler struct {
	Groups []*Group
}

func (gh *GroupHandler) Create() string {
	newGroup := NewGroup()
	gh.Groups = append(gh.Groups, newGroup)
	return newGroup.Code
}

func (gh *GroupHandler) GroupExists(groupCode string) bool {
	for _, v := range gh.Groups {
		if v.Code == groupCode {
			return true
		}
	}
	return false
}

func (gh *GroupHandler) getGroupByCode(groupCode string) *Group {
	for _, v := range gh.Groups {
		if v.Code == groupCode {
			return v
		}
	}
	return nil
}

func (gh *GroupHandler) AddParticipant(groupCode string, email string, phone string) error {
	groupToAddTo := gh.getGroupByCode(groupCode)
	if groupToAddTo != nil {
		groupToAddTo.AddParticipant(email, phone)
		return nil
	}
	return errors.New("Group Not Found For Code " + groupCode)
}

func (gh *GroupHandler) SendMessage(groupCode string, message string) error {
	groupToSendTo := gh.getGroupByCode(groupCode)
	if groupToSendTo != nil {
		groupToSendTo.SendMessage(message)
		return nil
	}
	return errors.New("Group Not Found For Code " + groupCode)
}
