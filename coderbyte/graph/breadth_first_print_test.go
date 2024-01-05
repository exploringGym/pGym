package graph

import (
	"reflect"
	"testing"
)

func TestBreadthFirstPrint(t *testing.T) {
	type args struct {
		graph  map[string][]string
		source string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			// c ---- a ---- b
			// |             |
			// e             d ---- g ---- h ---- j
			//               |      |
			//               f      i
			name: "Test 0",
			args: args{
				graph: map[string][]string{
					"a": {"b", "c"},
					"b": {"d"},
					"c": {"e"},
					"d": {"f", "g"},
					"e": {},
					"f": {},
					"g": {"i", "h"},
					"h": {"j"},
					"i": {},
					"j": {},
				},
				source: "a",
			},
			want: []string{"a", "b", "c", "d", "e", "f", "g", "i", "h", "j"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BreadthFirstPrint(tt.args.graph, tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BreadthFirstPrint() = %v, want %v", got, tt.want)
			}
		})
	}
}
