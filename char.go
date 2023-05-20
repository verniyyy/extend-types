package types

import (
	"strings"

	"github.com/samber/lo"
)

type Chars []*Char

func (cs Chars) Str() *Str {
	strings_ := lo.Map(
		cs,
		func(c *Char, _ int) string {
			return c.Get()
		},
	)
	return NewStr(strings.Join(strings_, ""))
}

type Char struct {
	basic[string]
}

func NewChar(char string) *Char {
	return &Char{
		newBasic(headOfString(char)),
	}
}

func (c *Char) Set(char string) {
	c.basic.Set(headOfString(char))
}

func headOfString(str string) string {
	return str[0:1]
}
