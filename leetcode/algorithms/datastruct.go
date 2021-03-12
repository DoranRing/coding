package algorithms

type ListNode struct {
	Val int

	Next *ListNode
}

func GenListNode(input []int) *ListNode {
	var head, cur *ListNode
	for _, val := range input {
		if cur == nil {
			head = &ListNode{Val: val}
			cur = head
		} else {
			cur.Next = &ListNode{Val: val}
			cur = cur.Next
		}
	}

	return head
}

func ListNodeEqual(input1, input2 *ListNode) bool {
	if input1 == nil && input2 == nil {
		return true
	}

	if input1 == nil || input2 == nil {
		return false
	}

	for input1 != nil {
		if input1.Val != input2.Val {
			return false
		}
		if input1.Next == nil && input2.Next == nil {
			return true
		}
		if input1.Next == nil || input2.Next == nil {
			return false
		}
		input1 = input1.Next
		input2 = input2.Next
	}

	return true
}

type TreeNode struct {
	Val int

	Left *TreeNode

	Right *TreeNode
}

// MinIntHeap 最小堆
type MinIntHeap []int

func (h MinIntHeap) Len() int {
	return len(h)
}

func (h MinIntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinIntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinIntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinIntHeap) Pop() interface{} {
	min := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return min
}

// MaxIntHeap 最大堆
type MaxIntHeap []int

func (h MaxIntHeap) Len() int {
	return len(h)
}

func (h MaxIntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxIntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxIntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxIntHeap) Pop() interface{} {
	min := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return min
}

func (h *MaxIntHeap) Top() int {
	return (*h)[0]
}
