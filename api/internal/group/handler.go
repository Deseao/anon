package group

type GroupHandler struct {
	Groups []*Group
}

func (gh *GroupHandler) Create() string {
	newGroup := NewGroup()
	gh.Groups = append(gh.Groups, newGroup)
	return newGroup.Code
}
