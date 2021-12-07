package linkedlist

// the interface of the LinkedList
type LinkedList interface {
	Get(int) int
	AddAtHead(int)
	AddAtTail(int)
	AddAtIndex(int, int)
	DeleteAtIndex(int)
}
