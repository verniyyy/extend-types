package list

import (
	"reflect"
	"testing"
)

func TestEnqueue(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name      string
		state     arrayList[string]
		args      args
		wantState arrayList[string]
	}{
		{
			name:  "Enqueue",
			state: arrayList[string]{"first", "second"},
			args: args{
				v: "third",
			},
			wantState: arrayList[string]{"first", "second", "third"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.state
			if q.Enqueue(tt.args.v); !reflect.DeepEqual(q, tt.wantState) {
				t.Errorf("queue[string].Enqueue() then %v, want %v", q, tt.wantState)
			}
		})
	}
}

func TestDequeue(t *testing.T) {
	tests := []struct {
		name      string
		state     arrayList[string]
		want      string
		wantState arrayList[string]
	}{
		{
			name:      "Dequeue",
			state:     arrayList[string]{"first", "second", "third"},
			want:      "first",
			wantState: arrayList[string]{"second", "third"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.state
			got := q.Dequeue()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queue[string].Dequeue() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(q, tt.wantState) {
				t.Errorf("queue[string].Dequeue() then %v, want %v", q, tt.wantState)
			}
		})
	}
}
