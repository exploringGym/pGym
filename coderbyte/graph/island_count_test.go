package graph

import "testing"

func TestIslandCount(t *testing.T) {
	type args struct {
		grid [][]string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 01",
			args: args{
				grid: [][]string{
					{"W", "L", "W", "W", "W"},
					{"W", "L", "W", "W", "W"},
					{"W", "W", "W", "L", "W"},
					{"W", "W", "L", "L", "W"},
					{"L", "W", "W", "L", "L"},
					{"L", "L", "W", "W", "W"},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IslandCount(tt.args.grid); got != tt.want {
				t.Errorf("IslandCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
