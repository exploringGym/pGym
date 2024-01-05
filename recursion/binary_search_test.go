package recursion

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		array  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 01",
			args: args{
				array:  []int{-1, 0, 1, 2, 3, 4, 7, 9, 10, 20},
				target: 0,
			},
			want: 1,
		},
		{
			name: "Test 02",
			args: args{
				array:  []int{-1, 0, 1, 2, 3, 4, 7, 9, 10, 20},
				target: 20,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.array, tt.args.target); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
