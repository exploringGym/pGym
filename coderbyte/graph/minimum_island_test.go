package graph

import "testing"

func TestFindMinimumSizeIsland(t *testing.T) {
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
			want: 2,
		},
		{
			name: "Test 02",
			args: args{
				grid: [][]string{
					{"L", "W", "W", "L", "W"},
					{"L", "W", "W", "L", "L"},
					{"W", "L", "W", "L", "W"},
					{"W", "W", "W", "W", "W"},
					{"W", "W", "L", "L", "L"},
				},
			},
			want: 1,
		},
		{
			name: "Test 03",
			args: args{
				grid: [][]string{
					{"L", "L", "L"},
					{"L", "L", "L"},
					{"L", "L", "L"},
				},
			},
			want: 9,
		},
		{
			name: "Test 04",
			args: args{
				grid: [][]string{
					{"W", "W", "W"},
					{"W", "W", "W"},
					{"W", "W", "W"},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMinimumSizeIsland(tt.args.grid); got != tt.want {
				t.Errorf("FindMinimumSizeIsland() = %v, want %v", got, tt.want)
			}
		})
	}
}
