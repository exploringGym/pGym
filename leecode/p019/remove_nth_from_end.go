package p019

// https://leetcode.com/problems/remove-nth-node-from-end-of-list

/*
Given the head of a linked list, remove the nth node from the end of the list
and return its head.

Example 1:

Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]

Example 2:

Input: head = [1], n = 1
Output: []

Example 3:

Input: head = [1,2], n = 1
Output: [1]

Constraints:

    The number of nodes in the list is sz.
    1 <= sz <= 30
    0 <= Node.val <= 100
    1 <= n <= sz



Follow up: Could you do this in one pass?

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n == 0 {
		return head
	}

	links := []*ListNode{}

	c := head
	size := 0

	for c != nil {
		links = append(links, c)
		c = c.Next
		size++
	}

	pos := size - n
	if pos == 0 {
		head = head.Next
	} else if pos == size-1 {
		l := links[pos-1]
		l.Next = nil
	} else {
		l := links[pos-1]
		r := links[pos+1]

		l.Next = r
	}

	return head
}

func removeNthFromEndBetter(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	current := dummy
	prev := dummy

	for current != nil {
		if n < 0 {
			prev = prev.Next
		}
		current = current.Next
		n--
	}

	prev.Next = prev.Next.Next
	return dummy.Next
}
