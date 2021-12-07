package linkedlist

type LinkedListB struct {
	Head *Node
	Size int
}

func NewLinkedListB() LinkedList {
	return &LinkedListB{
		Head: &Node{},
		Size: 0,
	}
}

func (l *LinkedListB) findNodeAtIndex(index int) *Node {
	if index < -1 || index >= l.Size {
		return nil
	}

	if index == -1 {
		return l.Head
	}

	cur := l.Head

	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Next
}

func (l *LinkedListB) Get(index int) int {
	node := l.findNodeAtIndex(index)
	if node != nil {
		return node.Val
	}
	return -1
}

func (l *LinkedListB) AddAtHead(val int) {
	l.AddAtIndex(0, val)
}

func (l *LinkedListB) AddAtTail(val int) {
	l.AddAtIndex(l.Size, val)
}

func (l *LinkedListB) AddAtIndex(index int, val int) {
	prevNode := l.findNodeAtIndex(index - 1)
	if prevNode == nil {
		return
	}
	newNode := &Node{Val: val}

	tmpNode := prevNode.Next
	prevNode.Next = newNode
	newNode.Next = tmpNode
	l.Size += 1
}

func (l *LinkedListB) DeleteAtIndex(index int) {
	prevNode := l.findNodeAtIndex(index - 1)
	if prevNode == nil {
		return
	}

	if prevNode.Next != nil {
		prevNode.Next = prevNode.Next.Next
		l.Size -= 1
	}
}
