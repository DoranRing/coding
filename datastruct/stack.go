package datastruct

import "errors"

// 栈
type Stack interface {

	// Top 获取栈顶元素
	Top() (int, error)

	// Pop 出栈
	Pop() (int, error)

	// Push 入栈
	Push(val int)
}

type ArrayStack struct {
	idx int
	arr []int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		-1,
		make([]int, 10),
	}
}

func (a *ArrayStack) Top() (int, error) {
	if a.idx < 0 {
		return 0, errors.New("stack empty")
	}
	return a.arr[a.idx], nil
}

func (a *ArrayStack) Pop() (int, error) {
	if a.idx < 0 {
		return 0, errors.New("stack empty")
	}
	if a.idx <= len(a.arr)/4 {
		a.resize(len(a.arr))
	}
	val := a.arr[a.idx]
	a.idx--
	return val, nil
}

func (a *ArrayStack) Push(val int) {
	if a.idx+1 >= len(a.arr) {
		a.resize(len(a.arr) * 2)
	}
	a.arr[a.idx+1] = val
	a.idx++
}

func (a *ArrayStack) resize(size int) {
	newArr := make([]int, size)
	copy(newArr, a.arr)
	a.arr = newArr
}

type LinkedStack struct {
	head *LinkedListNode
}

func NewLinkedStack() *LinkedStack {
	return &LinkedStack{}
}

func (l *LinkedStack) Top() (int, error) {
	if l.head == nil {
		return 0, errors.New("stack empty")
	}
	return l.head.val, nil
}

func (l *LinkedStack) Pop() (int, error) {
	if l.head == nil {
		return 0, errors.New("stack empty")
	}
	val := l.head.val
	l.head = l.head.next
	return val, nil
}

func (l *LinkedStack) Push(val int) {
	n := &LinkedListNode{val: val}
	if l.head == nil {
		l.head = n
	} else {
		n.next = l.head
		l.head = n
	}
}
