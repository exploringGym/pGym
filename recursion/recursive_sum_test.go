package recursion

import "testing"

func TestRecursiveSum(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 01",
			args: args{
				input: 10,
			},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RecursiveSum(tt.args.input); got != tt.want {
				t.Errorf("RecursiveSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
