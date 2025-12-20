package find_middle

type Node struct {
	Val  int
	Next *Node
}

func FindMiddle(head *Node) *Node {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

	}
	return slow
}
