package recursion

import "testing"

func TestIsPalindrome(t *testing.T) {
	type args struct {
		i string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 01",
			args: args{
				i: "kayak",
			},
			want: true,
		},
		{
			name: "Test 02",
			args: args{
				i: "abcd",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.i); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
