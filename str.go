package types

import (
	"fmt"
)

type Str struct {
	basic[string]
}

func NewStr[T ~string](s T) *Str {
	return &Str{
		newBasic(string(s)),
	}
}

func fromStringer(s fmt.Stringer) *Str {
	return &Str{
		newBasic(s.String()),
	}
}

// func (s *Str) Chars() Chars {
// 	return lo.Map(
// 		strings.Split(s.Read(), ""),
// 		func(c string, _ int) *Char {
// 			return NewChar(c)
// 		},
// 	)
// }
