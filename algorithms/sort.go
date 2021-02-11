package algorithms

import (
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
	for i := 1; i < len(nums); i++ {
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

// MergeSorter 归并排序(稳定排序)
// time:O(n*logN), space:O(n)
func MergeSorter(nums []int) {
	if len(nums) <= 1 {
		return
	}
	temp := make([]int, len(nums), len(nums))
	mergeSorter(nums, temp, 0, len(nums)-1)
}

func mergeSorter(nums, temp []int, left, right int) {
	if left == right {
		// 一个元素
		return
	} else if left+1 == right {
		// 排序两个相邻元素
		merge(nums, temp, left, right, right)
	} else {
		mid := (left + right) / 2
		// 分治前半部分
		mergeSorter(nums, temp, left, mid)
		// 分支后半部分
		mergeSorter(nums, temp, mid+1, right)
		// 归并前、后半部分
		merge(nums, temp, left, mid+1, right)
	}
}

func merge(nums, temp []int, l1, l2, r2 int) {
	idx := l1
	for i, j := l1, l2; i < l2 || j <= r2; idx++ {
		if i < l2 && j <= r2 {
			if nums[i] < nums[j] {
				temp[idx] = nums[i]
				i++
			} else {
				temp[idx] = nums[j]
				j++
			}
		} else if i < l2 {
			temp[idx] = nums[i]
			i++
		} else if j <= r2 {
			temp[idx] = nums[j]
			j++
		}
	}
	for i := l1; i <= r2; i++ {
		nums[i] = temp[i]
	}
}

// QuickSorter 快速排序
// time:O(n*longN), space:O(1)
func QuickSorter(nums []int) {
	if len(nums) <= 1 {
		return
	}
	quickSorter(nums, 0, len(nums)-1)
}

func quickSorter(nums []int, left, right int) {
	if left >= right {
		return
	}

	refVal, refIdx := nums[left], left
	for i, j := 0, len(nums)-1; i <= j; {
		// 小于值向前移动
		for j >= refIdx && nums[j] >= refVal {
			j--
		}
		if j >= refIdx {
			nums[refIdx] = nums[j]
			refIdx = j
		}
		// 大于值向后移动
		for i <= refIdx && nums[i] <= refVal {
			i++
		}
		if i <= refIdx {
			nums[refIdx] = nums[i]
			refIdx = i
		}
	}

	nums[refIdx] = refVal
	quickSorter(nums, left, refIdx-1)
	quickSorter(nums, refIdx+1, right)
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
	FilePath string

	// 线性表的长度
	len int

	// 打开的文件
	f *os.File
}

func NewHardList(filePath string) *HardList {
	list := &HardList{FilePath: filePath}
	list.Open()
	return list
}

func DefaultHardList() *HardList {
	return NewHardList("hard_data")
}

// Open 打开磁盘文件
func (h *HardList) Open() {
	f, err := os.Create(h.FilePath)
	if err != nil {
		log.Fatalf("open file error: %s\n", err.Error())
	}
	h.f = f
}

// Close 关闭线性表
func (h *HardList) Close() {
	if h.f == nil {
		return
	}
	err := h.f.Close()
	if err != nil {
		log.Fatalf("close file error: %s\n", err.Error())
	}
}

// Len 线性表长度
func (h *HardList) Len() int {
	return h.len
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

// Set 设置指定索引的数据
func (h *HardList) Set(idx int, num int32) {
	if idx >= h.len {
		log.Fatalf("outbound index, idx: %d, len: %d\n", idx, h.len)
	}
	data := h.intMapByte(num)
	h.writeDate(data, int64(idx*4))
}

// Append 线性表尾部添加元素
func (h *HardList) Append(num int32) {
	data := h.intMapByte(num)
	h.writeDate(data, int64(h.len*4))
	h.len++
}

// AppendList 线性表尾部批量添加元素
func (h *HardList) AppendList(nums []int32) {
	data := make([]byte, 0, 4*len(nums))
	for _, num := range nums {
		data = append(
			data,
			byte(num>>24),     // 0-7 bit
			byte(num<<8>>24),  // 8-15 bit
			byte(num<<16>>24), // 16-23 bit
			byte(num))         // 24-31bit
	}

	h.writeDate(data, int64(h.len*4))
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

// intMapByte int转换成byte[]
func (h *HardList) intMapByte(num int32) []byte {
	data := make([]byte, 0)
	data = append(
		data,
		byte(num>>24),     // 0-7 bit
		byte(num<<8>>24),  // 8-15 bit
		byte(num<<16>>24), // 16-23 bit
		byte(num))         // 24-31bit

	return data
}

// byteMapInt byte[]转换成int
func (h *HardList) byteMapInt(nums []byte, idx int) int32 {
	return int32(nums[idx])<<24 |
		int32(nums[idx+1])<<16 |
		int32(nums[idx+2])<<8 |
		int32(nums[idx+3])
}

// MemorySorter 内存内排序
func MemorySorter(nums []int32) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// HardSorter 磁盘内排序
func HardSorter(nums []int32) {
	hardList := DefaultHardList()
	defer hardList.Close()
	hardList.AppendList(nums)
	hardSorter(hardList)
}

// HardSorterValid 测试磁盘内排序
func HardSorterValid(nums []int32) []int32 {
	hardList := DefaultHardList()
	defer hardList.Close()
	hardList.AppendList(nums)
	hardSorter(hardList)
	return hardList.All()
}

// hardSorter 磁盘内排序
func hardSorter(hardList *HardList) {
	for i := 0; i < hardList.len; i++ {
		for j := 0; j < hardList.len-i-1; j++ {
			if hardList.Get(j) > hardList.Get(j+1) {
				cur := hardList.Get(j)
				next := hardList.Get(j + 1)
				hardList.Set(j, next)
				hardList.Set(j+1, cur)
			}
		}
	}
}
