package maybe

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func TestMaybeUnmarshalJSON(t *testing.T) {
	type T struct {
		Foo  int     `json:"foo"`
		Bar  string  `json:"bar"`
		Buzz float64 `json:"buzz"`
	}
	tests := []struct {
		name     string
		jsonFile string
		want     Maybe[T]
	}{
		{
			name:     "some",
			jsonFile: "testdata/sample1.json",
			want: Some(T{
				Foo:  1,
				Bar:  "bar",
				Buzz: 5.57,
			}),
		},
		{
			name:     "none",
			jsonFile: "testdata/sample2.json",
			want:     None[T](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := os.Open(tt.jsonFile)
			if err != nil {
				t.Fatalf("cannot read file: %s", tt.jsonFile)
			}
			defer file.Close()

			var decoded Maybe[T]
			if err := json.NewDecoder(file).Decode(&decoded); err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(decoded, tt.want) {
				t.Errorf("got %#v, want %#v", decoded, tt.want)
			}
		})
	}
}
