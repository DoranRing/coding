package algorithms

import (
	"bytes"
	"log"
	"os"
)

// BubbleSorter 冒泡排序(稳定排序)
// time:O(n2), space:O(1)
func BubbleSorter(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// SelectSorter 选择排序(稳定排序)
// time:O(n2), space:O(1)
func SelectSorter(nums []int) {
	for i := 0; i < len(nums); i++ {
		min := i
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
}

// InsertSorter 插入排序(稳定排序)
// time:O(n2), space:O(1)
func InsertSorter(nums []int) {
	insertSorter(nums, 0, len(nums)-1)
}

func insertSorter(nums []int, left, right int) {
	for i := left + 1; i <= right; i++ {
		for j := i; j > 0; j-- {
			if nums[j] >= nums[j-1] {
				break
			}
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}

// ShellSorter 希尔排序(不稳定排序)
// time:O(n1.3-2), space:O(1)
func ShellSorter(nums []int) {
	step := len(nums) / 2
	for step > 0 {
		// 每组分别排序
		for i := 0; i < step; i++ {
			// 插入排序
			for j := i; j < len(nums); j += step {
				for k := j; k >= step; k -= step {
					if nums[k] >= nums[k-step] {
						break
					}
					nums[k], nums[k-step] = nums[k-step], nums[k]
				}
			}
		}
		step /= 2
	}
}

// MergeSorterByRecursion 归并排序(稳定排序)
// 使用递归方式
// time:O(n*logN), space:O(n)
func MergeSorterByRecursion(nums []int) {
	if len(nums) <= 1 {
		return
	}
	temp := make([]int, len(nums), len(nums))
	mergeSorterByRecursion(nums, temp, 0, len(nums)-1)
}

func mergeSorterByRecursion(nums, temp []int, left, right int) {
	if left >= right {
		// 小于两个元素
		return
	}
	mid := (left + right) / 2
	// 分治前半部分
	mergeSorterByRecursion(nums, temp, left, mid)
	// 分治后半部分
	mergeSorterByRecursion(nums, temp, mid+1, right)
	// 归并前、后半部分
	merge(nums, temp, left, mid+1, right)
}

// MergeSorterByRecursion 归并排序(稳定排序)
// 使用迭代方式
// time:O(n*logN), space:O(n)
func MergeSorterByLoop(nums []int) {
	if len(nums) <= 1 {
		return
	}

	temp := make([]int, len(nums), len(nums))
	for step := 1; step < len(nums); step *= 2 {
		for i := 0; i+step < len(nums); i += step * 2 {
			r := i + step*2 - 1
			if r >= len(nums)-1 {
				r = len(nums) - 1
			}
			merge(nums, temp, i, i+step, r)
		}
	}
}

func merge(nums, temp []int, l1, l2, r2 int) {
	i, j := l1, l2
	for idx := l1; idx <= r2; idx++ {
		if i >= l2 {
			temp[idx] = nums[j]
			j++
		} else if j > r2 {
			temp[idx] = nums[i]
			i++
		} else if nums[i] < nums[j] {
			temp[idx] = nums[i]
			i++
		} else {
			temp[idx] = nums[j]
			j++
		}
	}
	for i := l1; i <= r2; i++ {
		nums[i] = temp[i]
	}
}

// QuickSorterByRecursion 快速排序
// time:O(n*longN), space:O(1)
func QuickSorterByRecursion(nums []int) {
	if len(nums) <= 1 {
		return
	}
	quickSorterByRecursion(nums, 0, len(nums)-1)
}

func quickSorterByRecursion(nums []int, left, right int) {
	if left+10 >= right {
		// 小数组切换到插入排序
		insertSorter(nums, left, right)
		return
	}
	idx := partition(nums, left, right)
	quickSorterByRecursion(nums, left, idx-1)
	quickSorterByRecursion(nums, idx+1, right)
}

func partition(nums []int, left, right int) int {
	i, j, val := left, right+1, nums[left]
	for true {
		i++
		for nums[i] <= val {
			if i == right {
				break
			}
			i++
		}
		j--
		for val <= nums[j] {
			if j == left {
				break
			}
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[left], nums[j] = nums[j], nums[left]
	return j
}

// BucketSorter 桶排序(稳定排序)
// time:O(m+n),space:O(m)
func BucketSorter(nums []int, max int) {
	bucket := make([]int, max+1, max+1)
	for _, num := range nums {
		bucket[num]++
	}

	idx := 0
	for i := 0; i < len(bucket); i++ {
		for bucket[i] > 0 {
			nums[idx] = i
			idx++
			bucket[i]--
		}
	}
}

// HardList 基于磁盘的线性表
type HardList struct {

	// 文件路径
	filePath string

	// 线性表的长度
	len int

	// 打开的文件
	f *os.File
}

func NewHardList(filePath string, len int) *HardList {
	list := &HardList{filePath: filePath}
	list.Open()
	list.init(len)
	return list
}

func DefaultHardList() *HardList {
	return NewHardList("hard_data", 0)
}

// Open 打开磁盘文件
func (h *HardList) Open() {
	h.f = h.openFile(h.filePath)
}

// openFile 打开磁盘文件
func (h *HardList) openFile(filePath string) *os.File {
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("open file error: %s\n", err.Error())
	}
	return f
}

// Len 线性表长度
func (h *HardList) Len() int {
	return h.len
}

func (h *HardList) init(len int) {
	h.len = len
	if len <= 0 {
		return
	}
	_, err := h.f.WriteAt([]byte{0}, int64(len*4-1))
	if err != nil {
		log.Fatalf("init error: %s\n", err.Error())
	}
}

func (h *HardList) Destroy() {
	err := h.f.Close()
	if err != nil {
		log.Fatalf("close file(%s) error: %s", h.filePath, err.Error())
	}
	err = os.Remove(h.filePath)
	if err != nil {
		log.Fatalf("remove file(%s) error: %s", h.filePath, err.Error())
	}
}

// Get 获取指定索引的数据
func (h *HardList) Get(idx int) int32 {
	if idx >= h.len {
		log.Fatalf("outbound index. len:%d,idx:%d\n", h.len, idx)
	}
	data := make([]byte, 4)
	_, err := h.f.ReadAt(data, int64(idx*4))
	if err != nil {
		log.Fatalf("read file error: %s\n", err.Error())
	}
	return h.byteMapInt(data, 0)
}

// gets 获取一块数据
func (h *HardList) gets(idx int, buf []byte) {
	if idx >= h.len {
		log.Fatalf("outbound index. len:%d,idx:%d\n", h.len, idx)
	}
	size, err := h.f.ReadAt(buf, int64(idx*4))
	if size == 0 && err != nil {
		log.Fatalf("read file error: %s\n", err.Error())
	}
}

// Set 设置指定索引的数据
func (h *HardList) Set(idx int, num int32) {
	if idx >= h.len {
		log.Fatalf("outbound index, idx: %d, len: %d\n", idx, h.len)
	}
	data := h.mapByte(num)
	h.writeDate(data, int64(idx*4))
}

// sets 设置一块数据
func (h *HardList) sets(idx int, buf []byte) {
	if idx >= h.len {
		log.Fatalf("outbound index. len:%d,idx:%d\n", h.len, idx)
	}
	h.writeDate(buf, int64(idx*4))
}

// Append 线性表尾部添加元素
func (h *HardList) Append(num int32) {
	data := h.mapByte(num)
	h.writeDate(data, int64(h.len*4))
	h.len++
}

// AppendList 线性表尾部批量添加元素
func (h *HardList) AppendList(nums []int32) {
	var buf bytes.Buffer
	for _, num := range nums {
		h.appendByte(&buf, num)
	}
	h.writeDate(buf.Bytes(), int64(h.len*4))
	h.len += len(nums)
}

// All 获取线性表的所有值
func (h *HardList) All() []int32 {
	data := make([]byte, 4*h.len)
	_, err := h.f.ReadAt(data, 0)
	if err != nil {
		log.Fatalf("")
	}
	res := make([]int32, h.len, h.len)
	for i := 0; i < len(data); i += 4 {
		res[i/4] = h.byteMapInt(data, i)
	}
	return res
}

// writeDate 指定偏移量位置设置值
func (h *HardList) writeDate(data []byte, offset int64) {
	write, err := h.f.WriteAt(data, offset)
	if err != nil {
		log.Fatalf("write error: %s\n", err.Error())
	}
	if write != len(data) {
		log.Fatalf("write error, data len: %d, write len: %d\n", len(data), write)
	}
}

// mapByte int转换成byte[]
func (h *HardList) mapByte(num int32) []byte {
	return []byte{
		byte(num >> 24),       // 0-7 bit
		byte(num << 8 >> 24),  // 8-15 bit
		byte(num << 16 >> 24), // 16-23 bit
		byte(num),             // 24-31 bit
	}
}

// appendByte int转换成byte后填充到尾部
func (h *HardList) appendByte(buf *bytes.Buffer, num int32) {
	buf.WriteByte(byte(num >> 24))       // 0-7 bit
	buf.WriteByte(byte(num << 8 >> 24))  // 8-15 bit
	buf.WriteByte(byte(num << 16 >> 24)) // 16-23 bit
	buf.WriteByte(byte(num))             // 24-31 bit
}

// byteMapInt byte[]转换成int
func (h *HardList) byteMapInt(nums []byte, idx int) int32 {
	// |0-7 bit|8-15 bit|16-23 bit|24-31 bit|
	return int32(nums[idx])<<24 |
		int32(nums[idx+1])<<16 |
		int32(nums[idx+2])<<8 |
		int32(nums[idx+3])
}

// BubbleSort 使用冒泡排序排序所有元素
func (h *HardList) BubbleSort() {
	for i := 0; i < h.len; i++ {
		for j := 0; j < h.len-i-1; j++ {
			if h.Get(j) > h.Get(j+1) {
				cur := h.Get(j)
				next := h.Get(j + 1)
				h.Set(j, next)
				h.Set(j+1, cur)
			}
		}
	}
}

// MergeSort 归并排序
// 磁盘具备随机读取功能，所有只需要1*n的额外空间
// 磁带不具备随机读取功能，所以需要多个额外空间
func (h *HardList) MergeSort() {
	if h.len <= 1 {
		return
	}
	tempFileName := h.filePath + ".temp"
	tempList := NewHardList(tempFileName, h.len)
	defer tempList.Destroy()
	h.mergeSort(tempList, 0, h.len-1)
}

func (h *HardList) mergeSort(tempList *HardList, left, right int) {
	if left == right {
		return
	} else if left+1 == right {
		h.merge(tempList, left, right, right)
	} else {
		mid := (left + right) / 2
		h.mergeSort(tempList, left, mid)
		h.mergeSort(tempList, mid+1, right)
		h.merge(tempList, left, mid+1, right)
	}
}

func (h *HardList) merge(tempList *HardList, aLeft, bLeft, bRight int) {
	aRight := bLeft - 1
	left, idx := aLeft, aLeft
	for aLeft <= aRight && bLeft <= bRight {
		aVal := h.Get(aLeft)
		bVal := h.Get(bLeft)
		if aVal < bVal {
			tempList.Set(idx, aVal)
			idx, aLeft = idx+1, aLeft+1
		} else {
			tempList.Set(idx, bVal)
			idx, bLeft = idx+1, bLeft+1
		}
	}
	for aLeft <= aRight {
		tempList.Set(idx, h.Get(aLeft))
		idx, aLeft = idx+1, aLeft+1
	}
	for bLeft <= bRight {
		tempList.Set(idx, h.Get(bLeft))
		idx, bLeft = idx+1, bLeft+1
	}

	// copy
	for i := left; i <= bRight; i++ {
		h.Set(i, tempList.Get(i))
	}

}

// MemorySorter 基于内存的冒泡排序
func MemorySorter(nums []int32) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// HardSorter 基于磁盘的冒泡排序
func HardSorter(nums []int32) {
	hardList := DefaultHardList()
	defer hardList.Destroy()
	hardList.AppendList(nums)
	hardList.BubbleSort()
}

// HardBubbleSorterValid 测试基于磁盘的冒泡排序
func HardBubbleSorterValid(nums []int32) []int32 {
	hardList := DefaultHardList()
	defer hardList.Destroy()
	hardList.AppendList(nums)
	hardList.BubbleSort()
	return hardList.All()
}

// HardBubbleSorterValid 基于磁盘的归并排序
func HardMergeSorter(nums []int32) {
	hardList := DefaultHardList()
	defer hardList.Destroy()
	hardList.AppendList(nums)
	hardList.MergeSort()
}

// HardBubbleSorterValid 测试基于磁盘的归并排序
func HardMergeSorterValid(nums []int32) []int32 {
	hardList := DefaultHardList()
	defer hardList.Destroy()
	hardList.AppendList(nums)
	hardList.MergeSort()
	return hardList.All()
}
