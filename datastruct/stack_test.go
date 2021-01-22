package datastruct

import (
	"testing"
)

func TestNewArrayStack(t *testing.T) {
	stack := NewArrayStack()
	testStack(t, stack)
}

func TestNewLinkedStack(t *testing.T) {
	stack := NewLinkedStack()
	testStack(t, stack)
}

func testStack(t *testing.T, stack Stack) {
	stack.Push(10)
	top, err := stack.Top()
	if err != nil {
		t.Errorf("error: %t\n", err)
	}
	if top != 10 {
		t.Errorf("actual %d, want %d\n", top, 10)
	}
	top, err = stack.Pop()
	if err != nil {
		t.Errorf("error: %t\n", err)
	}
	if top != 10 {
		t.Errorf("actual %d, want %d\n", top, 10)
	}
}
