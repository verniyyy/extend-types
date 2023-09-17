package types

import (
	"fmt"
	"strings"
	"unsafe"
)

type Str struct {
	basic[string]
}

func NewStr[T ~string](s T) Str {
	return Str{
		newBasic(string(s)),
	}
}

func NewStrFromBytes(data []byte) Str {
	s := *(*string)(unsafe.Pointer(&data))
	return Str{
		newBasic(s),
	}
}

func fromStringer(s fmt.Stringer) Str {
	return Str{
		newBasic(s.String()),
	}
}

func (s Str) Size() int {
	return len(s.value())
}

func (s Str) Bytes() []byte {
	return unsafe.Slice(unsafe.StringData(s.value()), s.Size())
}

func (s Str) Split(sep string) []Str {
	strList := strings.Split(s.value(), sep)
	result := make([]Str, len(strList))
	for i, item := range strList {
		result[i] = NewStr(item)
	}
	return result
}

func (s Str) Include(sub string) bool {
	return strings.Contains(s.value(), sub)
}
