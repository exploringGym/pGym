package graph

import (
	"reflect"
	"testing"
)

func TestDepthFirstPrint(t *testing.T) {
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
			want: []string{"a", "c", "e", "b", "d", "g", "h", "j", "i", "f"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DepthFirstPrint(tt.args.graph, tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DepthFirstPrint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDepthFirstPrintRec(t *testing.T) {
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
			want: []string{"a", "b", "d", "g", "h", "j", "i", "f", "c", "e"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DepthFirstPrintRec(tt.args.graph, tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DepthFirstPrintRec() = %v, want %v", got, tt.want)
			}
		})
	}
}
