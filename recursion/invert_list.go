package recursion

type Node struct {
	Val  int
	Next *Node
}

func InvertList(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	p := InvertList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
