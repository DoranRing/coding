package algorithms

type ListNode struct {
	Val int

	Next *ListNode
}

func genListNode(input []int) *ListNode {
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

func listNodeEqual(input1, input2 *ListNode) bool {
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

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	min := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return min
}
