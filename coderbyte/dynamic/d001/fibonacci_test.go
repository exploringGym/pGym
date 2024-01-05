package d001

import "testing"

func TestFibonacci(t *testing.T) {
	type args struct {
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
				target: 6,
			},
			want: 8,
		},
		{
			name: "Test 02",
			args: args{
				target: 7,
			},
			want: 13,
		},
		{
			name: "Test 03",
			args: args{
				target: 8,
			},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibonacci(tt.args.target); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
