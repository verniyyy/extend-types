package types

import (
	"testing"
)

func TestBool_Xor(t *testing.T) {
	tests := []struct {
		name  string
		bool_ Bool
		arg   bool
		want  bool
	}{
		{
			name:  "true, true -> false",
			bool_: Bool(true),
			arg:   true,
			want:  false,
		},
		{
			name:  "true, false -> true",
			bool_: Bool(true),
			arg:   false,
			want:  true,
		},
		{
			name:  "false, true -> true",
			bool_: Bool(false),
			arg:   true,
			want:  true,
		},
		{
			name:  "false, false -> false",
			bool_: Bool(false),
			arg:   false,
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bool_.XOr(tt.arg); got != tt.want {
				t.Errorf("Bool.XOr() = %v, want %v", got, tt.want)
			}
		})
	}
}
