package p019

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Test 01",
			args: args{
				head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
				n:    2,
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}},
		},
		{
			name: "Test 02",
			args: args{
				head: &ListNode{Val: 1},
				n:    1,
			},
			want: &ListNode{},
		},
		{
			name: "Test 03",
			args: args{
				head: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
				n:    1,
			},
			want: &ListNode{Val: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				c := got
				e := tt.want
				if e != nil {
					for c != nil {
						if c.Val != e.Val {
							t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
						}
						c = c.Next
						e = e.Next
					}
				} else {
					if c != nil {
						t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
					}
				}
			}
		})
	}
}
