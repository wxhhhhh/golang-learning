package algorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

func findMidNode(head *ListNode) *ListNode {
	fastNode := head
	slowNode := head

	for fastNode.Next != nil && fastNode.Next.Next != nil {
		fastNode = fastNode.Next.Next
		slowNode = slowNode.Next
	}
	return slowNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev
}

func mergeList(l1, l2 *ListNode) {
	var l1Tmp, l2Tmp *ListNode
	for l1 != nil && l2 != nil {
		l1Tmp = l1.Next
		l2Tmp = l2.Next

		l1.Next = l2
		l1 = l1Tmp

		l2.Next = l1
		l2 = l2Tmp
	}
}

func reorderList(head *ListNode) {

	midNode := findMidNode(head)
	l1 := head
	l2 := midNode.Next
	midNode.Next = nil
	l2 = reverseList(l2)
	mergeList(l1, l2)
}
