package algorithms

// Table 符号表接口
type Table interface {
	Set(key string, value string)

	Get(key string) (string, bool)
}

type TableEntry struct {
	key string
	val string
}

// StringListNode string链表节点
type StringListNode struct {
	TableEntry
	next *StringListNode
}

// NonSortedListTable 无序的链式线性表实现的符号表
type NonSortedListTable struct {
	head *StringListNode
	last *StringListNode
}

func (n *NonSortedListTable) Set(key string, value string) {
	node := &StringListNode{TableEntry: TableEntry{key: key, val: value}}
	if n.head == nil {
		n.head = node
		n.last = node
	} else {
		existNode, ok := n.getNode(key)
		if ok {
			// set
			existNode.val = value
		} else {
			// append
			n.last.next = node
			n.last = n.last.next
		}
	}
}

func (n *NonSortedListTable) Get(key string) (string, bool) {
	if node, ok := n.getNode(key); ok {
		return node.val, true
	}
	return "", false
}

// getNode 获取匹配的节点
func (n *NonSortedListTable) getNode(key string) (*StringListNode, bool) {
	cur := n.head
	for cur != nil {
		if cur.key == key {
			return cur, true
		}
		cur = cur.next
	}
	return nil, false
}

// SortedArrayTable 排序的连续线性表实现的符号表
type SortedArrayTable struct {
	arr []TableEntry
}

func (s *SortedArrayTable) Set(key string, value string) {
	if match, ok := s.getMatch(key); ok {
		// set
		match.val = value
	} else {
		// append
		entry := TableEntry{key: key, val: value}
		s.arr = append(s.arr, entry)
		// sort
		s.sort()
	}
}

func (s *SortedArrayTable) Get(key string) (string, bool) {
	if entry, ok := s.getMatch(key); ok {
		return entry.val, ok
	}
	return "", false
}

// getMatch 获取匹配的元素
// 已排序数据使用二分查找
func (s SortedArrayTable) getMatch(key string) (*TableEntry, bool) {
	l, r := 0, len(s.arr)-1
	for l <= r {
		mid := (l + r) / 2
		if s.arr[mid].key == key {
			return &s.arr[mid], true
		} else if s.arr[mid].key < key {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return nil, false
}

// sort 排序元素
// 已排序数据使用插入排序
func (s *SortedArrayTable) sort() {
	for i := len(s.arr) - 1; i > 0; i-- {
		if s.arr[i].key >= s.arr[i-1].key {
			break
		}
		s.arr[i], s.arr[i-1] = s.arr[i-1], s.arr[i]
	}
}

// BinaryTreeTableNode 二叉查找树实现的符号表节点
type BinaryTreeTableNode struct {
	TableEntry
	left  *BinaryTreeTableNode
	right *BinaryTreeTableNode
}

func NewBinaryTreeTableNode(key, value string) *BinaryTreeTableNode {
	return &BinaryTreeTableNode{TableEntry: TableEntry{key: key, val: value}}
}

// BinaryTreeTable 二叉查找树实现的符号表
type BinaryTreeTable struct {
	root *BinaryTreeTableNode
}

func (b *BinaryTreeTable) Set(key, value string) {
	if b.root == nil {
		b.root = NewBinaryTreeTableNode(key, value)
	} else {
		b.set(key, value, b.root)
	}
}

func (b *BinaryTreeTable) set(key, value string, node *BinaryTreeTableNode) {
	if node.key == key {
		node.val = value
	} else if node.key > key {
		if node.left == nil {
			node.left = NewBinaryTreeTableNode(key, value)
		} else {
			b.set(key, value, node.left)
		}
	} else {
		if node.right == nil {
			node.right = NewBinaryTreeTableNode(key, value)
		} else {
			b.set(key, value, node.right)
		}
	}
}

func (b *BinaryTreeTable) Get(key string) (string, bool) {
	return b.get(key, b.root)
}

func (b *BinaryTreeTable) get(key string, node *BinaryTreeTableNode) (string, bool) {
	if node == nil {
		return "", false
	}
	if node.key == key {
		return node.val, true
	} else if node.key > key {
		return b.get(key, node.left)
	} else {
		return b.get(key, node.right)
	}
}

// BinarySearchByBase 二分查找
// 匹配多个时选取的值不固定
func BinarySearchByBase(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

// BinarySearchByRangeFirst 二分查找
// 查询符合条件的第一个元素
func BinarySearchByRangeFirst(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)
	for l < r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if l != len(nums) && nums[l] == target {
		return l
	}
	return -1
}

// CollisionPointer 碰撞指针
// 集合中两个元素的关联性，这里找出两个元素和等于target
func CollisionPointer(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		} else if sum < target {
			l++
		} else {
			r--
		}
	}
	return []int{}
}

// SlowFastPoint 快慢指针
// 集合中两个元素的关联性，这里找出尾部节点是否指向了首部节点
func SlowFastPoint(head *ListNode) bool {
	var fast, slow *ListNode
	for head != nil {
		// fast+1
		if fast == nil {
			fast = head
		} else {
			fast = fast.Next
		}

		// fast+1,slow+1
		if fast != nil && fast.Next != nil {
			fast = fast.Next
			if slow == nil {
				slow = head
			} else {
				slow = slow.Next
			}
		} else {
			// next is nil
			return false
		}

		// equal
		if slow == fast {
			return true
		}
	}

	return false
}

type ListNode struct {
	Val int

	Next *ListNode
}

// SlidingWindow 滑动窗口
// 检查一个子集符合要求
func SlidingWindow(target int, nums []int) int {
	// init window
	l, r := 0, -1
	min, sum := len(nums)+1, 0

	// sliding window
	for l < len(nums) {
		if r < len(nums)-1 && sum < target {
			r++
			sum += nums[r]
		} else {
			sum -= nums[l]
			l++
		}
		if sum >= target {
			window := r - l + 1
			if window == 1 {
				return 1
			} else {
				if min > window {
					min = window
				}
			}
		}
	}

	if min == len(nums)+1 {
		return 0
	}
	return min
}
