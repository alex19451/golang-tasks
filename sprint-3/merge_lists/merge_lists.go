package mergelists

type Node struct {
	Val  int
	Next *Node
}

func MergeLists(list1 *Node, list2 *Node) *Node {
	dummy_list := &Node{}
	current := dummy_list

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return dummy_list.Next
}
