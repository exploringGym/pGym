package graph

import "testing"

func TestFindShortestPath(t *testing.T) {
	type args struct {
		edges [][]string
		start string
		end   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 01",
			args: args{
				edges: [][]string{
					{"w", "x"},
					{"x", "y"},
					{"z", "y"},
					{"z", "v"},
					{"w", "v"},
				},
				start: "w",
				end:   "z",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindShortestPath(tt.args.edges, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("FindShortestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
