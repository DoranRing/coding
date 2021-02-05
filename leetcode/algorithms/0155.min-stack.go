package algorithms

//执行用时: 16 ms 内存消耗: 9.5 MB
type MinStack struct {
	list []int

	minIdx int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{[]int{}, -1}
}

func (this *MinStack) Push(x int) {
	this.list = append(this.list, x)
	if this.minIdx == -1 || x < this.list[this.minIdx] {
		this.minIdx = len(this.list) - 1
	}
}

func (this *MinStack) Pop() {
	reset := this.minIdx == len(this.list)-1
	this.list = this.list[:len(this.list)-1]
	if reset {
		this.resetMin()
	}
}

func (this *MinStack) Top() int {
	return this.list[len(this.list)-1]
}

func (this *MinStack) GetMin() int {
	return this.list[this.minIdx]
}

func (this *MinStack) resetMin() {
	idx := 0
	for i := 0; i < len(this.list); i++ {
		if this.list[idx] > this.list[i] {
			idx = i
		}
	}
	this.minIdx = idx
}
