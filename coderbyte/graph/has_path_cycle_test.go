package graph

import (
	"reflect"
	"testing"
)

func TestHasPathCycleRec(t *testing.T) {
	type args struct {
		graph map[string][]string
		src   string
		dst   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test 01",
			args: args{
				graph: map[string][]string{
					"i": {"j", "k"},
					"j": {"k", "i"},
					"k": {"i", "j", "l", "m"},
					"l": {"k"},
					"m": {"k"},
					"o": {"n"},
					"n": {"o"},
				},
				src: "j",
				dst: "m",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasPathCycleRec(tt.args.graph, tt.args.src, tt.args.dst); got != tt.want {
				t.Errorf("HasPathCycleRec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildGraph(t *testing.T) {
	type args struct {
		edges [][]string
	}
	tests := []struct {
		name string
		args args
		want map[string][]string
	}{
		{
			name: "Test 01",
			args: args{
				edges: [][]string{
					{"i", "j"},
					{"k", "i"},
					{"m", "k"},
					{"k", "l"},
					{"o", "n"},
				},
			},
			want: map[string][]string{
				"i": {"j", "k"},
				"j": {"i"},
				"k": {"i", "m", "l"},
				"l": {"k"},
				"m": {"k"},
				"n": {"o"},
				"o": {"n"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildGraph(tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}
