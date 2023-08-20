package list

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

func TestNewArrayListFromSlice(t *testing.T) {
	hoges := []*Hoge{
		{Str: "hoge"},
		{Str: "fuga"},
	}
	l := NewArrayListFromSlice(hoges)
	l.Print()
}

func TestArrayList_Find(t *testing.T) {
	l := arrayList[*Hoge]{}
	l.Add(&Hoge{Str: "hoge"})
	l.Add(&Hoge{Str: "fuga"})

	found, _ := l.Find(func(v *Hoge) bool { return v.GetStr() == "hogee" })
	fmt.Printf("%+v\n", found)

	l2 := arrayList[Hoge]{}
	l2.Add(Hoge{Str: "hoge"})
	l2.Add(Hoge{Str: "fuga"})

	found2, _ := l2.Find(func(v Hoge) bool { return v.GetStr() == "hogee" })
	fmt.Printf("%+v\n", found2)
}

func TestArrayList_Each(t *testing.T) {
	l := arrayList[Hoge]{}
	l.Add(Hoge{Str: "hoge"})
	l.Add(Hoge{Str: "fuga"})

	e := l.Each(func(v *Hoge) { v.Str = "piyo" })
	e.Print()
}

func TestArrayList_EachWithIndex(t *testing.T) {
	l := arrayList[Hoge]{}
	l.Add(Hoge{Str: "hoge"})
	l.Add(Hoge{Str: "fuga"})

	l2 := NewArrayList[string](len(l))

	l.EachWithIndex(func(v *Hoge, i int) { l2.Overwrite(i, v.Str) })
	l2.Print()
}

func TestArrayList_Print(t *testing.T) {
	l := arrayList[*Hoge]{}
	l.Add(&Hoge{
		Str: "hoge",
		X:   3,
		Y:   7,
	})
	l.Add(&Hoge{
		Str: "fuga",
		X:   8,
		Y:   11,
	})

	l.Print()

	l2 := arrayList[string]{"aa", "bb", "cc"}
	l2.Print()
}

func TestArrayList_index(t *testing.T) {
	l := arrayList[*Hoge]{}
	l.Add(&Hoge{Str: "hoge"})
	l.Add(&Hoge{Str: "fuga"})

	l2 := l
	l2.Add(&Hoge{Str: "piyo"})

	type args struct {
		index int
	}
	tests := []struct {
		name  string
		state arrayList[*Hoge]
		args  args
		want  int
	}{
		// TODO: Add test cases.
		{
			name:  "ok",
			state: l,
			args: args{
				index: 0,
			},
			want: 0,
		},
		{
			name:  "ok",
			state: l,
			args: args{
				index: 1,
			},
			want: 1,
		},
		{
			name:  "ok",
			state: l,
			args: args{
				index: 2,
			},
			want: 0,
		},
		{
			name:  "ok",
			state: l,
			args: args{
				index: 3,
			},
			want: 1,
		},
		{
			name:  "ok",
			state: l,
			args: args{
				index: -1,
			},
			want: 1,
		},
		{
			name:  "ok",
			state: l,
			args: args{
				index: -2,
			},
			want: 0,
		},
		{
			name:  "ok",
			state: l,
			args: args{
				index: -3,
			},
			want: 1,
		},
		{
			name:  "ok",
			state: l,
			args: args{
				index: -4,
			},
			want: 0,
		},
		{
			name:  "got index is 0, if list size 3 and input 0",
			state: l2,
			args: args{
				index: 0,
			},
			want: 0,
		},
		{
			name:  "got index is 1, if list size 3 and input 1",
			state: l2,
			args: args{
				index: 1,
			},
			want: 1,
		},
		{
			name:  "got index is 2, if list size 3 and input 2",
			state: l2,
			args: args{
				index: 2,
			},
			want: 2,
		},
		{
			name:  "got index is 0, if list size 3 and input 3",
			state: l2,
			args: args{
				index: 3,
			},
			want: 0,
		},
		{
			name:  "got index is 2, if list size 3 and input -1",
			state: l2,
			args: args{
				index: -1,
			},
			want: 2,
		},
		{
			name:  "got index is 1, if list size 3 and input -2",
			state: l2,
			args: args{
				index: -2,
			},
			want: 1,
		},
		{
			name:  "got index is 0, if list size 3 and input -3",
			state: l2,
			args: args{
				index: -3,
			},
			want: 0,
		},
		{
			name:  "got index is 2, if list size 3 and input -4",
			state: l2,
			args: args{
				index: -4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.state.index(tt.args.index); got != tt.want {
				t.Errorf("arrayList.index() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestStr_Clone(t *testing.T) {
// 	type state arrayList[*Hoge]
// 	type args struct {
// 		str string
// 	}
// 	tests := []struct {
// 		name  string
// 		state state
// 		args  args
// 		want  string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			s := Str{
// 				basic: tt.fields.basic,
// 			}
// 			if got := s.Concat(tt.args.str); got != tt.want {
// 				t.Errorf("Str.Concat() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
