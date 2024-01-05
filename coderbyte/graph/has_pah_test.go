package graph

import (
	"testing"
)

func Test_hasPathDFS(t *testing.T) {
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
			//  /-------------\
			//  g ---- f ---- i
			//  |             |
			//  h             k
			name: "test 01",
			args: args{
				graph: map[string][]string{
					"f": {"g", "i"},
					"g": {"h"},
					"h": {},
					"i": {"g", "k"},
					"j": {"i"},
					"k": {},
				},
				src: "f",
				dst: "k",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathDFS(tt.args.graph, tt.args.src, tt.args.dst); got != tt.want {
				t.Errorf("hasPathDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasPathDFSRec(t *testing.T) {
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
			//  /-------------\
			//  g ---- f ---- i
			//  |             |
			//  h             k
			name: "test 01",
			args: args{
				graph: map[string][]string{
					"f": {"g", "i"},
					"g": {"h"},
					"h": {},
					"i": {"g", "k"},
					"j": {"i"},
					"k": {},
				},
				src: "f",
				dst: "k",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathDFSRec(tt.args.graph, tt.args.src, tt.args.dst); got != tt.want {
				t.Errorf("hasPathDFSRec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasPathBST(t *testing.T) {
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
			//  /-------------\
			//  g ---- f ---- i
			//  |             |
			//  h             k
			name: "test 01",
			args: args{
				graph: map[string][]string{
					"f": {"g", "i"},
					"g": {"h"},
					"h": {},
					"i": {"g", "k"},
					"j": {"i"},
					"k": {},
				},
				src: "f",
				dst: "k",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathBST(tt.args.graph, tt.args.src, tt.args.dst); got != tt.want {
				t.Errorf("hasPathBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
