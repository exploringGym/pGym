package recursion

import (
	"testing"
)

func TestInvertList(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "Test 01",
			args: args{
				head: &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3, Next: &Node{Val: 4}}}},
			},
			want: &Node{Val: 4, Next: &Node{Val: 3, Next: &Node{Val: 2, Next: &Node{Val: 1}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := InvertList(tt.args.head)
			for got != nil {
				if got.Val != tt.want.Val {
					t.Errorf("InvertList() = %v, want %v", got, tt.want)
				}
				got = got.Next
				tt.want = tt.want.Next
			}
		})
	}
}
