package list

import "testing"

func Test_indexError_Error(t *testing.T) {
	type fields struct {
		size  uint
		index int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ok",
			fields: fields{
				size:  3,
				index: 5,
			},
			want: "index 5 outside of list bounds: -3...2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := indexError{
				size:  tt.fields.size,
				index: tt.fields.index,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("indexError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
