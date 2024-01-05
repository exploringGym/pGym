package p2095

import (
	"reflect"
	"testing"
)

func Test_findMissingAndRepeatedValues(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 01",
			args: args{
				grid: [][]int{{1, 3}, {2, 2}},
			},
			want: []int{2, 4},
		},
		{
			name: "Test 02",
			args: args{
				grid: [][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}},
			},
			want: []int{9, 5},
		},
		{
			name: "Test 03",
			args: args{
				grid: [][]int{{2, 2}, {3, 4}},
			},
			want: []int{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMissingAndRepeatedValues(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMissingAndRepeatedValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
