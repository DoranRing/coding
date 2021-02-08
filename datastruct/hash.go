package datastruct

// HashEntry 哈希表节点
type HashEntry struct {

	// key 哈希表key
	key string

	// value 哈希表value
	value string
}

// HashEntryLinkedList 哈希表节点的链表
type HashEntryLinkedList interface {

	// add 添加哈希表节点
	add(entry *HashEntry)

	// get 获取哈希表节点
	get(key string) (*HashEntry, bool)

	// delete 删除哈希表节点
	delete(key string)
}

// HashEntryLinkNode 哈希表节点的链表的节点
type HashEntryLinkNode struct {
	val *HashEntry

	next *HashEntryLinkNode
}

type DefaultHashEntryLinkedList struct {
	head *HashEntryLinkNode

	last *HashEntryLinkNode
}

func (d *DefaultHashEntryLinkedList) add(entry *HashEntry) {
	node := &HashEntryLinkNode{entry, nil}
	if d.last == nil {
		d.head = node
		d.last = node
	} else {
		d.last.next = node
		d.last = node
	}
}

func (d *DefaultHashEntryLinkedList) get(key string) (*HashEntry, bool) {
	cur := d.head
	for cur != nil {
		if cur.val.key == key {
			return cur.val, true
		}
		cur = cur.next
	}
	return nil, false
}

func (d *DefaultHashEntryLinkedList) delete(key string) {
	if d.head == nil {
		return
	}

	if d.head.val.key == key {
		d.head = d.head.next
		return
	}
	cur := d.head
	for cur != nil && cur.next != nil {
		if cur.next.val.key == key {
			cur.next = cur.next.next
			break
		}
		cur = cur.next
	}
}

// Hash 哈希表ADT
type Hash interface {

	// Put 存放键值对
	Put(string, string)

	// Get 获取key对应的value
	Get(string) (string, bool)

	// Delete 删除key的值
	Delete(string)

	// Contains 检查哈希表内包含key
	Contains(string) bool
}

// SeparateChainHash 分离链接法实现的哈希表
type SeparateChainHash struct {
	size int32

	arr []HashEntryLinkedList
}

func NewSeparateChainHash(size int32) *SeparateChainHash {
	arr := make([]HashEntryLinkedList, size, size)
	return &SeparateChainHash{
		size: size,
		arr:  arr,
	}
}

func (h *SeparateChainHash) Put(key string, value string) {
	hashVal := h.hash(key)
	list := h.arr[hashVal]
	if list == nil {
		list = &DefaultHashEntryLinkedList{}
	}
	entry := &HashEntry{key, value}
	list.add(entry)
	h.arr[hashVal] = list
}

func (h *SeparateChainHash) Get(key string) (string, bool) {
	hashVal := h.hash(key)
	list := h.arr[hashVal]
	if list != nil {
		if entry, ok := list.get(key); ok {
			return entry.value, ok
		}
	}
	return "", false
}

func (h *SeparateChainHash) Delete(key string) {
	hashVal := h.hash(key)
	list := h.arr[hashVal]
	if list != nil {
		list.delete(key)
	}
}

func (h *SeparateChainHash) Contains(key string) bool {
	_, ok := h.Get(key)
	return ok
}
func (h *SeparateChainHash) hash(key string) int32 {
	var hashVal int32
	for _, c := range key {
		hashVal = 37*hashVal + c
	}

	hashVal %= h.size
	if hashVal < 0 {
		hashVal += h.size
	}
	return hashVal
}
