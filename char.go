package types

// type Chars []*Char

// func (cs Chars) Str() *Str {
// 	strings_ := lo.Map(
// 		cs,
// 		func(c *Char, _ int) string {
// 			return string(c.Read())
// 		},
// 	)
// 	return NewStr(strings.Join(strings_, ""))
// }

type Char struct {
	basic[rune]
}

func NewChar(char string) *Char {
	for _, r := range char {
		return &Char{
			newBasic(r),
		}
	}
	return &Char{
		newBasic(rune(0)),
	}
}

func (c *Char) String() string {
	return string(c.Read())
}

func (c *Char) Byte() byte {
	return byte(c.Read())
}
