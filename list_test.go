package types

import (
	"fmt"
	"testing"
)

type Hoge struct {
	Str string
	X   int
	Y   int
}

func (h *Hoge) GetStr() string {
	return h.Str
}

func TestNewListFromSlice(t *testing.T) {
	hoges := []*Hoge{
		{Str: "hoge"},
		{Str: "fuga"},
	}
	l := NewListFromSlice(hoges)
	l.Print()
}

func TestList_Find(t *testing.T) {
	l := List[*Hoge]{}
	l.Append(&Hoge{Str: "hoge"})
	l.Append(&Hoge{Str: "fuga"})

	found, _ := l.Find(func(v *Hoge) bool { return v.GetStr() == "hogee" })
	fmt.Printf("%+v\n", found)

	l2 := List[Hoge]{}
	l2.Append(Hoge{Str: "hoge"})
	l2.Append(Hoge{Str: "fuga"})

	found2, _ := l2.Find(func(v Hoge) bool { return v.GetStr() == "hogee" })
	fmt.Printf("%+v\n", found2)
}

func TestList_Each(t *testing.T) {
	l := List[Hoge]{}
	l.Append(Hoge{Str: "hoge"})
	l.Append(Hoge{Str: "fuga"})

	e := l.Each(func(v *Hoge) { v.Str = "piyo" })
	e.Print()
}

func TestList_EachWithIndex(t *testing.T) {
	l := List[Hoge]{}
	l.Append(Hoge{Str: "hoge"})
	l.Append(Hoge{Str: "fuga"})

	l2 := NewList[string](len(l))

	l.EachWithIndex(func(v *Hoge, i int) { l2[i] = v.Str })
	l2.Print()
}

func TestList_Print(t *testing.T) {
	l := List[*Hoge]{}
	l.Append(&Hoge{
		Str: "hoge",
		X:   3,
		Y:   7,
	})
	l.Append(&Hoge{
		Str: "fuga",
		X:   8,
		Y:   11,
	})

	l.Print()

	l2 := List[string]{"aa", "bb", "cc"}
	l2.Print()
}
