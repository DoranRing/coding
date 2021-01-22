package datastruct

import "errors"

type Queue interface {

	// Top 队列首部的元素
	Top() (int, error)

	// Dequeue 出队
	Dequeue() (int, error)

	// Enqueue 入队
	Enqueue(val int)
}

type LinkedQueue struct {
	head *node
	last *node
}

func (l *LinkedQueue) Top() (int, error) {
	if l.head == nil {
		return 0, errors.New("queue empty")
	}
	return l.head.val, nil
}

func (l *LinkedQueue) Dequeue() (int, error) {
	if l.head == nil {
		return 0, errors.New("queue empty")
	}
	val := l.head.val
	l.head = l.head.next
	return val, nil
}

func (l *LinkedQueue) Enqueue(val int) {
	n := &node{val: val}
	if l.head == nil {
		l.head = n
		l.last = n
	} else {
		l.last.next = n
		l.last = n
	}
}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{}
}
