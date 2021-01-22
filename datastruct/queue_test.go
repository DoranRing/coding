package datastruct

import (
	"testing"
)

func TestNewLinkedQueue(t *testing.T) {
	queue := NewLinkedQueue()
	testQueue(t, queue)
}

func testQueue(t *testing.T, queue Queue) {
	queue.Enqueue(10)
	top, err := queue.Top()
	if err != nil {
		t.Errorf("error: %t\n", err)
	}
	if top != 10 {
		t.Errorf("actual %d, want %d\n", top, 10)
	}
	top, err = queue.Dequeue()
	if err != nil {
		t.Errorf("error: %t\n", err)
	}
	if top != 10 {
		t.Errorf("actual %d, want %d\n", top, 10)
	}
	_, err = queue.Dequeue()
	if err == nil {
		t.Errorf("actual nil, want error")
	}

}
