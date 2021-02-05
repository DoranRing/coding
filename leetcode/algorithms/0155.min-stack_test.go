package algorithms

import (
	"fmt"
	"testing"
)

func Test_MinStack(t *testing.T) {
	minStack := Constructor()
	minStack.Push(85)
	minStack.Push(-99)
	minStack.Push(67)
	minStack.Pop()
	minStack.Pop()
	min := minStack.GetMin()
	fmt.Printf("min:%d", min)
}
