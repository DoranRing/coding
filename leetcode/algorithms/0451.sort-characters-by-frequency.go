package algorithms

import "container/heap"

//time:O(2n*logN),space:O(2n)
//执行用时: 8 ms 内存消耗: 5.2 MB
func frequencySortByHeap(s string) string {
	// 统计次数
	hash := make(map[rune]int)
	runeCount := 0
	for _, runeVal := range s {
		runeCount++
		count := 0
		if exist, ok := hash[runeVal]; ok {
			count = exist
		}
		count++
		hash[runeVal] = count
	}

	// 堆排序
	maxHeap := &MaxFrequencySortHeap{}
	heap.Init(maxHeap)
	for k, v := range hash {
		entry := &FrequencySortEntry{key: k, count: v}
		heap.Push(maxHeap, entry)
	}

	// 输出结果
	runeSlice := make([]rune, 0, runeCount)
	for maxHeap.Len() > 0 {
		top := heap.Pop(maxHeap).(*FrequencySortEntry)
		for i := 0; i < top.count; i++ {
			runeSlice = append(runeSlice, top.key)
		}
	}

	return string(runeSlice)
}

type FrequencySortEntry struct {
	key   rune
	count int
}

type MaxFrequencySortHeap struct {
	arr []*FrequencySortEntry
}

func (m *MaxFrequencySortHeap) Len() int {
	return len(m.arr)
}

func (m *MaxFrequencySortHeap) Less(i, j int) bool {
	return m.arr[i].count > m.arr[j].count
}

func (m *MaxFrequencySortHeap) Swap(i, j int) {
	m.arr[i], m.arr[j] = m.arr[j], m.arr[i]
}

func (m *MaxFrequencySortHeap) Push(x interface{}) {
	m.arr = append(m.arr, x.(*FrequencySortEntry))
}

func (m *MaxFrequencySortHeap) Pop() interface{} {
	entry := m.arr[len(m.arr)-1]
	m.arr = m.arr[:len(m.arr)-1]
	return entry
}
