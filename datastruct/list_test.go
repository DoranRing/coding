package datastruct

import "testing"

func TestArrayList_Size(t *testing.T) {
	list := NewArrayList()
	if list.Size() != 0 {
		t.Errorf("acutal %d, want %d\n", list.Size(), 0)
	}
	list.Add(1)
	if list.Size() == 0 {
		t.Errorf("actual %d, want %d\n", list.Size(), 1)
	}
}

func TestArrayList_Empty(t *testing.T) {
	list := NewArrayList()
	if !list.Empty() {
		t.Errorf("acutal %t, want %t\n", list.Empty(), true)
	}
	list.Add(1)
	if list.Empty() {
		t.Errorf("actual %t, want %t\n", list.Empty(), false)
	}
}

func TestArrayList_Clear(t *testing.T) {
	list := NewArrayList()
	list.Add(1)
	list.Clear()
	if !list.Empty() {
		t.Errorf("acutal %t, want %t\n", list.Empty(), true)
	}
}

func TestArrayList_Contains(t *testing.T) {
	list := NewArrayList()
	list.Add(1)
	if !list.Contains(1) {
		t.Errorf("acutal %t, want %t\n", list.Contains(1), true)
	}
	if list.Contains(2) {
		t.Errorf("acutal %t, want %t\n", list.Contains(2), false)
	}
}

func TestArrayList_Add(t *testing.T) {
	list := NewArrayList()
	list.Add(10)
	list.Add(20)
	val, err := list.Get(0)
	if err != nil {
		t.Errorf("error: %e\n", err)
	}
	if val != 10 {
		t.Errorf("acutal %d, want %d\n", val, 10)
	}
	val, err = list.Get(1)
	if err != nil {
		t.Errorf("error: %e\n", err)
	}
	if val != 20 {
		t.Errorf("acutal %d, want %d\n", val, 20)
	}
}

func TestArrayList_Remove(t *testing.T) {
	list := NewArrayList()
	list.Add(10)
	list.Add(20)
	list.Add(30)
	r := list.Remove(1)
	if !r {
		t.Errorf("actual %t, want %t\n", r, true)
	}
	val, err := list.Get(1)
	if err != nil {
		t.Errorf("error: %e\n", err)
	}
	if val != 30 {
		t.Errorf("actual %d, want %d\n", val, 30)
	}
}

func TestLinkedList_Add(t *testing.T) {
	list := NewLinkedList()
	list.Add(10)
	list.Add(20)
	list.Add(30)
	if val, _ := list.Get(0); val != 10 {
		t.Errorf("actual %d, want %d\n", val, 10)
	}
	if val, _ := list.Get(1); val != 20 {
		t.Errorf("actual %d, want %d\n", val, 20)
	}
	if val, _ := list.Get(2); val != 30 {
		t.Errorf("actual %d, want %d\n", val, 30)
	}

}

func TestLinkedList_Remove(t *testing.T) {
	list := NewLinkedList()
	list.Add(10)
	list.Add(20)
	list.Add(30)
	list.Remove(1)
	if list.Size() != 2 {
		t.Errorf("actual %d, want %d", list.Size(), 2)
	}
	if val, _ := list.Get(1); val != 30 {
		t.Errorf("actual %d, want %d\n", val, 30)
	}
	list.Remove(1)
	if list.Size() != 1 {
		t.Errorf("actual %d, want %d", list.Size(), 1)
	}
	if val, _ := list.Get(0); val != 10 {
		t.Errorf("actual %d, want %d\n", val, 10)
	}
}

func TestDoubleLinkedList_Add(t *testing.T) {
	list := NewDoubleLinkedList()
	list.Add(10)
	list.Add(20)
	list.Add(30)
	list.Add(40)
	if list.Size() != 4 {
		t.Errorf("actual %d, want %d", list.Size(), 4)
	}
}

func TestDoubleLinkedList_Set(t *testing.T) {
	list := NewDoubleLinkedList()
	list.Add(10)
	list.Add(20)
	err := list.Set(1, 25)
	if err != nil {
		t.Errorf("error: %e", err)
	}
	val, err := list.Get(1)
	if err != nil {
		t.Errorf("error: %e", err)
	}
	if val != 25 {
		t.Errorf("actual %d, want %d", val, 25)
	}
}

func TestDoubleLinkedList_Remove(t *testing.T) {
	list := NewDoubleLinkedList()
	list.Add(10)
	list.Add(20)
	list.Add(30)
	list.Add(40)

	remove := list.Remove(2)
	if !remove {
		t.Errorf("actual %t, want %t", remove, true)
	}
	val, err := list.Get(2)
	if err != nil {
		t.Errorf("error: %e", err)
	}
	if val != 40 {
		t.Errorf("actual %d, want %d", val, 40)
	}
}
