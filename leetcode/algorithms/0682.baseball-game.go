package algorithms

import "strconv"

// time:O(n), space:O(n)
//执行用时: 0 ms 内存消耗: 6.5 MB
func calPointsByStack(ops []string) int {
	stack := NewArrayStack()
	for _, item := range ops {
		if item == "C" {
			stack.Pop()
		} else if item == "D" {
			top, _ := stack.Pop()
			stack.Push(top)
			stack.Push(top * 2)
		} else if item == "+" {
			a, _ := stack.Pop()
			b, _ := stack.Pop()
			stack.Push(b)
			stack.Push(a)
			stack.Push(a + b)
		} else {
			if num, err := strconv.Atoi(item); err == nil {
				stack.Push(num)
			}
		}
	}

	sum := 0
	for stack.Len() > 0 {
		top, _ := stack.Pop()
		sum += top
	}
	return sum
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

func (a *ArrayStack) Len() int {
	return a.idx + 1
}

func (a *ArrayStack) Pop() (int, bool) {
	if a.idx < 0 {
		return 0, false
	}
	if a.idx <= len(a.arr)/4 {
		a.resize(len(a.arr))
	}
	val := a.arr[a.idx]
	a.idx--
	return val, true
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
