package linkedlist

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	a := NewLinkedLisA()
	a.AddAtIndex(3, 1)
	a.AddAtIndex(0, 2)
	a.AddAtIndex(0, 3)
	a.DeleteAtIndex(1)
	t.Logf("%d", a.Get(0))
}
