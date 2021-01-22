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

// LinkedList 单向链式线性表
type LinkedList struct {
	len  int
	head *node
	last *node
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
	n := &node{val: val}
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

type node struct {
	val  int
	next *node
}

// DoubleLinkedList 双向链式线性表
type DoubleLinkedList struct {
	len  int
	head *doubleNode
	last *doubleNode
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
	n := &doubleNode{val: val}
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

func (d *DoubleLinkedList) match(idx int) *doubleNode {
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

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{}
}

type doubleNode struct {
	val  int
	pre  *doubleNode
	next *doubleNode
}
