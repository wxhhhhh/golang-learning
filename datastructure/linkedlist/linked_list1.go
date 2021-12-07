package linkedlist

// linked list with dummy head
type LinkedListA struct {
	Head *Node
}

func NewLinkedLisA() LinkedList {
	return &LinkedListA{
		Head: &Node{
			Val:  0,
			Next: nil,
		},
	}
}

func (l *LinkedListA) findNodeAtIndex(index int) *Node {
	// invalid input
	if index < -1 {
		return nil
	}
	// corner case, find the dummy head. (real head's prev node)
	if index == -1 {
		return l.Head
	}

	cur := l.Head
	cnt := 0

	for ; index > 0 && cur.Next != nil; index-- {
		cur = cur.Next
		cnt++
	}

	return cur.Next
}

func (l *LinkedListA) findLastNode() *Node {
	cur := l.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	return cur
}

func (l *LinkedListA) Get(index int) int {
	node := l.findNodeAtIndex(index)
	if node != nil {
		return node.Val
	}
	return -1
}

func (l *LinkedListA) AddAtHead(val int) {
	l.AddAtIndex(0, val)
}

func (l *LinkedListA) AddAtTail(val int) {
	newNode := &Node{Val: val}
	node := l.findLastNode()
	node.Next = newNode
}

func (l *LinkedListA) AddAtIndex(index, val int) {
	node := l.findNodeAtIndex(index - 1)

	if node == nil {
		return
	}

	newNode := &Node{Val: val}
	// insert the node
	tmpNode := node.Next
	node.Next = newNode
	newNode.Next = tmpNode

}

func (l *LinkedListA) DeleteAtIndex(index int) {
	node := l.findNodeAtIndex(index - 1)

	// delete the node
	if node != nil && node.Next != nil {
		node.Next = node.Next.Next
	}
}
