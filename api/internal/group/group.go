package group

import "github.com/Deseao/anon/api/internal/code"

type Group struct {
	Code string
}

func NewGroup() *Group {
	return &Group{Code: code.GenRandCode(code.CODE_LEN)}
}
