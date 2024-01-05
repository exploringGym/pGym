package p2951

import (
	"reflect"
	"testing"
)

func Test_findPeaks(t *testing.T) {
	type args struct {
		mountain []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 01",
			args: args{
				mountain: []int{2, 4, 4},
			},
			want: []int{},
		},
		{
			name: "Test 02",
			args: args{
				mountain: []int{1, 4, 3, 8, 5},
			},
			want: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPeaks(tt.args.mountain); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findPeaks() = %v, want %v", got, tt.want)
			}
		})
	}
}
