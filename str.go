package types

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

type Str struct {
	basic[string]
}

func NewStr(s string) *Str {
	return &Str{
		newBasic(s),
	}
}

func fromStringer(s fmt.Stringer) *Str {
	return &Str{
		newBasic(s.String()),
	}
}

func (s *Str) Chars() Chars {
	return lo.Map(
		strings.Split(s.Get(), ""),
		func(c string, _ int) *Char {
			return NewChar(c)
		},
	)
}
