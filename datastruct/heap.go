package datastruct

import (
	"container/heap"
	"log"
)

// Heap 堆
type Heap interface {

	// Insert 添加一个元素
	Insert(num int)

	// Top 获取堆顶的元素
	Top() int

	// Pop 弹出堆顶的元素
	Pop() int
}

// NonSortedArrayHeap 无序的连续线性表堆
type NonSortedArrayHeap struct {

	// arr 容器
	arr []int

	// length 长度
	length int
}

func NewNonSortedArrayHeap(length int) *NonSortedArrayHeap {
	return &NonSortedArrayHeap{length: length}
}

func (n *NonSortedArrayHeap) Insert(num int) {
	if len(n.arr) >= n.length {
		// 空间已满时，替换最小的元素
		min := 0
		for i := 1; i < len(n.arr); i++ {
			if n.arr[min] > n.arr[i] {
				min = i
			}
		}
		if n.arr[min] < num {
			n.arr[min] = num
		}
	} else {
		n.arr = append(n.arr, num)
	}
}

func (n *NonSortedArrayHeap) Top() int {
	if len(n.arr) <= 0 {
		log.Fatalf("heap is empty")
	}
	max := 0
	for i := 1; i < len(n.arr); i++ {
		if n.arr[max] < n.arr[i] {
			max = i
		}
	}
	return n.arr[max]
}

func (n *NonSortedArrayHeap) Pop() int {
	if len(n.arr) <= 0 {
		log.Fatalf("heap is empty")
	}
	max := 0
	for i := 1; i < len(n.arr); i++ {
		if n.arr[max] < n.arr[i] {
			max = i
		}
	}
	val := n.arr[max]
	n.arr = append(n.arr[:max], n.arr[max+1:]...)
	return val
}

// SortedArrayHeap 有序的连续线性表堆
type SortedArrayHeap struct {

	// arr 容器
	arr []int

	// length 长度
	length int
}

func NewSortedArrayHeap(length int) *SortedArrayHeap {
	return &SortedArrayHeap{length: length}
}

func (s *SortedArrayHeap) Insert(num int) {
	if len(s.arr) == 0 {
		s.arr = append(s.arr, num)
		return
	}

	if len(s.arr) >= s.length {
		// 空间已满,从前向后插入排序
		if s.arr[0] <= num {
			s.arr[0] = num
		}
		for i := 0; i < len(s.arr)-1; i++ {
			if s.arr[i] <= s.arr[i+1] {
				break
			}
			s.arr[i], s.arr[i+1] = s.arr[i+1], s.arr[i]
		}
	} else {
		// 空间不满，从后向前插入排序
		s.arr = append(s.arr, num)
		for i := len(s.arr) - 1; i > 0; i-- {
			if s.arr[i] >= s.arr[i-1] {
				break
			}
			s.arr[i], s.arr[i-1] = s.arr[i-1], s.arr[i]
		}
	}
}

func (s *SortedArrayHeap) Top() int {
	if len(s.arr) <= 0 {
		log.Fatalf("heap is empty")
	}
	return s.arr[len(s.arr)-1]
}

func (s *SortedArrayHeap) Pop() int {
	if len(s.arr) <= 0 {
		log.Fatalf("heap is empty")
	}
	val := s.arr[len(s.arr)-1]
	s.arr = append(s.arr[:len(s.arr)-1])
	return val
}

// MaxBinaryArrayHeap 二叉堆
type MaxBinaryArrayHeap struct {

	// 容器
	slice []int
}

func NewMaxBinaryHeap() *MaxBinaryArrayHeap {
	return &MaxBinaryArrayHeap{
		slice: []int{0},
	}
}

func (b *MaxBinaryArrayHeap) Insert(num int) {
	b.slice = append(b.slice, num)
	b.up(len(b.slice) - 1)
}

func (b *MaxBinaryArrayHeap) Top() int {
	if len(b.slice) <= 1 {
		log.Fatalf("heap is empty")
	}
	return b.slice[1]
}

// Pop 删除最大值
// 1.存储顶点元素
// 2.尾部元素设置为顶点元素，然后将顶点元素一直下沉
func (b *MaxBinaryArrayHeap) Pop() int {
	if len(b.slice) <= 1 {
		log.Fatalf("heap is empty")
	}
	val := b.slice[1]
	b.slice[1] = b.slice[len(b.slice)-1]
	b.slice = append(b.slice[:len(b.slice)-1])
	b.down(1)
	return val
}

// down 节点下沉
// 1.获取两个子节点最大的节点A
// 2.当前节点N的值小于最大子节点A则与之交换，然后继续下层最大子节点A
// 3.当前节点N的值大于等于子节点A则不再继续
func (b *MaxBinaryArrayHeap) down(k int) {
	for k*2 < len(b.slice) {
		max := k * 2
		if max+1 < len(b.slice) && b.slice[max] < b.slice[max+1] {
			max++
		}
		if b.slice[k] >= b.slice[max] {
			break
		}
		b.slice[k], b.slice[max] = b.slice[max], b.slice[k]
		k = max
	}
}

// up 节点上浮
// 1.当前节点N大于父节点P，则交换节点N-P，然后继续上浮节点P
// 2.当前节点N小于或者登录父节点P则不再继续
func (b *MaxBinaryArrayHeap) up(k int) {
	for k > 1 && b.slice[(k/2)] < b.slice[k] {
		b.slice[(k/2)], b.slice[k] = b.slice[k], b.slice[(k/2)]
		k /= 2
	}
}

// KthLargest kth大的元素
// 使用最小堆，添加元素之后如果空间已满则删除堆顶的最小值
type KthLargest struct {
	arr []int

	k int
}

func NewKthLargest(k int) *KthLargest {
	return &KthLargest{k: k}
}

func (kl *KthLargest) Len() int {
	return len(kl.arr)
}

func (kl *KthLargest) Less(i, j int) bool {
	return kl.arr[i] < kl.arr[j]
}

func (kl *KthLargest) Swap(i, j int) {
	kl.arr[i], kl.arr[j] = kl.arr[j], kl.arr[i]
}

// Push implement heap.Interface
func (kl *KthLargest) Push(x interface{}) {
	kl.arr = append(kl.arr, x.(int))
}

// Pop implement heap.Interface
func (kl *KthLargest) Pop() interface{} {
	val := kl.arr[kl.Len()-1]
	kl.arr = kl.arr[:kl.Len()-1]
	return val
}

// Insert 添加元素
func (kl *KthLargest) Insert(val int) int {
	if kl.Len() < kl.k || kl.arr[0] < val {
		heap.Push(kl, val)
		if kl.Len() > kl.k {
			kl.Delete()
		}
	}
	return kl.arr[0]
}

// InsertList 批量添加元素
func (kl *KthLargest) InsertList(arr []int) {
	for _, val := range arr {
		kl.Insert(val)
	}
}

// Delete 删除顶端的元素，即最小的值
func (kl *KthLargest) Delete() int {
	return heap.Pop(kl).(int)
}
