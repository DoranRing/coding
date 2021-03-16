package datastruct

import "errors"

// 线性表
type List interface {

	// Size 获取线性表元素的数量
	Size() int

	// Empty 检查线性表元素为空
	Empty() bool

	// Clear 清空线性表
	Clear()

	// Contains 检查线性表包含元素
	Contains(val int) bool

	// Add 在线性表的尾部添加元素
	Add(val int)

	// Remove 删除元素
	Remove(idx int) bool

	// Set 设置线性表指定索引的值
	Set(idx, val int) error

	// Get 获取线性表指定索引的值
	Get(idx int) (int, error)
}

// ArrayList 顺序线性表
type ArrayList struct {
	idx int
	arr []int
}

func NewArrayList() *ArrayList {
	return &ArrayList{
		idx: 0,
		arr: make([]int, 10),
	}
}

func (a *ArrayList) Size() int {
	return a.idx
}

func (a *ArrayList) Empty() bool {
	return a.idx == 0
}

func (a *ArrayList) Clear() {
	a.idx = 0
}

func (a *ArrayList) Contains(val int) bool {
	for i := 0; i < a.idx; i++ {
		if a.arr[i] == val {
			return true
		}
	}
	return false
}

func (a *ArrayList) Add(val int) {
	a.grow()
	a.arr[a.idx] = val
	a.idx += 1
}

func (a *ArrayList) Remove(idx int) bool {
	if idx >= a.idx {
		return false
	}
	for i := idx; i < a.idx-1; i++ {
		a.arr[i] = a.arr[i+1]
	}
	a.idx -= 1
	return true
}

func (a *ArrayList) Set(idx, val int) error {
	if idx >= a.idx {
		return errors.New("outbound index")
	}
	a.arr[idx] = val
	return nil
}

func (a *ArrayList) Get(idx int) (int, error) {
	if idx >= a.idx {
		return 0, errors.New("outbound index")
	}
	return a.arr[idx], nil
}

func (a *ArrayList) grow() {
	if a.idx >= len(a.arr) {
		newArr := make([]int, len(a.arr)*2)
		copy(newArr, a.arr)
		a.arr = newArr
	}
}

// LinkedListNode 单向链式线性表的节点
type LinkedListNode struct {
	val  int
	next *LinkedListNode
}

// NewLinkedListNode 创建链式线性表节点
func NewLinkedListNode(val int) *LinkedListNode {
	return &LinkedListNode{val: val}
}

// GenLinkedListNodeByLastInsert 使用尾插方式根据输入数组创建链式线性表节点
func GenLinkedListNodeByLastInsert(nums []int) *LinkedListNode {
	var head, cur *LinkedListNode
	for _, num := range nums {
		node := NewLinkedListNode(num)
		if cur == nil {
			head = node
			cur = node
		} else {
			cur.next = node
			cur = cur.next
		}
	}
	return head
}

// GenLinkedListNodeByHeadInsert 使用头插方式根据输入数组创建链式线性表节点
func GenLinkedListNodeByHeadInsert(nums []int) *LinkedListNode {
	var head *LinkedListNode
	for _, num := range nums {
		node := NewLinkedListNode(num)
		node.next = head
		head = node
	}
	return head
}

func (n *LinkedListNode) DepthEqual(head *LinkedListNode) bool {
	if n == nil && head == nil {
		return true
	}
	if n != nil && head == nil {
		return false
	}
	if n == nil && head != nil {
		return false
	}
	cur, otherCur := n, head
	for cur != nil && otherCur != nil {
		if cur.val != otherCur.val {
			return false
		}
		cur, otherCur = cur.next, otherCur.next
	}
	return cur == nil && otherCur == nil
}

// LinkedList 单向链式线性表
type LinkedList struct {
	len  int
	head *LinkedListNode
	last *LinkedListNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Size() int {
	return l.len
}

func (l *LinkedList) Empty() bool {
	return l.len == 0
}

func (l *LinkedList) Clear() {
	l.len, l.head, l.last = 0, nil, nil
}

func (l *LinkedList) Contains(val int) bool {
	cur := l.head
	for i := 1; i < l.len; i++ {
		if cur.val == val {
			return true
		}
		cur = cur.next
	}
	return false
}

func (l *LinkedList) Add(val int) {
	n := &LinkedListNode{val: val}
	if l.last != nil {
		l.last.next = n
		l.last = l.last.next
	} else {
		l.head = n
		l.last = n
	}
	l.len++
}

func (l *LinkedList) Remove(idx int) bool {
	if idx >= l.len {
		return false
	}
	if idx == 0 {
		l.head = l.head.next
	} else {
		cur := l.head
		for i := 0; i < idx-1; i++ {
			cur = cur.next
		}
		cur.next = cur.next.next
		if idx == l.len-1 {
			l.last = cur
		}
	}
	l.len--
	return true
}

func (l *LinkedList) Set(idx, val int) error {
	if idx >= l.len {
		return errors.New("outbound index")
	}
	cur := l.head
	for i := 0; i < idx; i++ {
		cur = cur.next
	}
	cur.val = val
	return nil
}

func (l *LinkedList) Get(idx int) (int, error) {
	if idx >= l.len {
		return 0, errors.New("outbound index")
	}
	cur := l.head
	for i := 0; i < idx; i++ {
		cur = cur.next
	}
	return cur.val, nil
}

// DoubleLinkedListNode 双向链式线性表的节点
type DoubleLinkedListNode struct {
	val  int
	pre  *DoubleLinkedListNode
	next *DoubleLinkedListNode
}

// NewDoubleLinkedList 创建双向链式线性表的节点
func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{}
}

// DoubleLinkedList 双向链式线性表
type DoubleLinkedList struct {
	len  int
	head *DoubleLinkedListNode
	last *DoubleLinkedListNode
}

func (d *DoubleLinkedList) Size() int {
	return d.len
}

func (d *DoubleLinkedList) Empty() bool {
	return d.len == 0
}

func (d *DoubleLinkedList) Clear() {
	d.len, d.head, d.last = 0, nil, nil
}

func (d *DoubleLinkedList) Contains(val int) bool {
	cur := d.head
	for i := 1; i < d.len; i++ {
		if cur.val == val {
			return true
		}
		cur = cur.next
	}
	return false
}

func (d *DoubleLinkedList) Add(val int) {
	n := &DoubleLinkedListNode{val: val}
	if d.last == nil {
		d.head, d.last = n, n
	} else {
		d.last.next = n
		n.pre = d.last
		d.last = n
	}
	d.len++
}

func (d *DoubleLinkedList) Remove(idx int) bool {
	if idx >= d.len {
		return false
	}
	match := d.match(idx)
	match.pre.next = match.next
	match.next.pre = match.pre
	d.len--
	return true
}

func (d *DoubleLinkedList) Set(idx, val int) error {
	if idx >= d.len {
		return errors.New("outbound index")
	}
	match := d.match(idx)
	match.val = val
	return nil
}

func (d *DoubleLinkedList) Get(idx int) (int, error) {
	if idx >= d.len {
		return 0, errors.New("outbound index")
	}
	return d.match(idx).val, nil
}

func (d *DoubleLinkedList) match(idx int) *DoubleLinkedListNode {
	half := idx >= d.len/2
	// 大于一半，从尾部开始
	if half {
		cur := d.last
		for i := d.len - 1; i > idx; i-- {
			cur = cur.pre
		}
		return cur
	}
	// 小于一半，从头部开始
	cur := d.head
	for i := 0; i < idx; i++ {
		cur = cur.next
	}
	return cur
}

type LinkedListReverser interface {

	// 翻转整个链表
	Reverse(head *LinkedListNode) *LinkedListNode

	// 翻转区间
	ReverseRange(head *LinkedListNode, m, n int) *LinkedListNode
}

// RecursionLinkedListReverser 基于递归的单链表翻转
type RecursionLinkedListReverser struct {
}

func (r RecursionLinkedListReverser) Reverse(head *LinkedListNode) *LinkedListNode {
	if head == nil || head.next == nil {
		return head
	}
	newHead := r.Reverse(head.next)
	head.next.next = head
	head.next = nil
	return newHead
}

func (r RecursionLinkedListReverser) ReverseRange(head *LinkedListNode, m, n int) *LinkedListNode {
	pre := &LinkedListNode{next: head}
	for i := 0; i < m; i++ {
		pre = pre.next
	}
	cur := pre.next
	rangeHead, rangeNext := r.reverseRange(cur, n-m)
	cur.next = rangeNext
	pre.next = rangeHead

	if m == 0 {
		return rangeHead
	}
	return head
}

func (r RecursionLinkedListReverser) reverseRange(node *LinkedListNode, n int) (*LinkedListNode, *LinkedListNode) {
	if n == 0 {
		return node, node.next
	}
	newHead, rangeNext := r.reverseRange(node.next, n-1)
	node.next.next = node
	node.next = nil
	return newHead, rangeNext
}

// HeadInsertLinkedListReverser 头插法的单链表翻转
type HeadInsertLinkedListReverser struct {
}

func (h HeadInsertLinkedListReverser) Reverse(head *LinkedListNode) *LinkedListNode {
	var newHead *LinkedListNode
	for head != nil {
		next := head.next
		// move new
		head.next = newHead
		newHead = head
		head = next
	}

	return newHead
}

func (h HeadInsertLinkedListReverser) ReverseRange(head *LinkedListNode, m, n int) *LinkedListNode {
	if m == 0 {
		return h.reverseRange(head, n-m)
	}

	pre := head
	for i := 1; i < m; i++ {
		pre = pre.next
	}
	pre.next = h.reverseRange(pre.next, n-m)
	return head
}

func (h *HeadInsertLinkedListReverser) reverseRange(node *LinkedListNode, n int) *LinkedListNode {
	var head, last *LinkedListNode = nil, node
	for i := 0; i <= n; i++ {
		next := node.next
		node.next = head
		head = node
		node = next
	}
	// link last next
	last.next = node
	return head
}
