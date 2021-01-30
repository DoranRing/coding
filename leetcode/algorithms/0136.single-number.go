package algorithms

// 利用相同数异或等于0的特性，将所有出现两次的数都转换为0，最后得到只出现一次的值

// time:O(n) space:O(1)
// 执行用时： 20 ms 内存消耗： 6.1 MB
func singleNumberByXOR(nums []int) int {
	res := 0
	for _, num := range nums {
		res = res ^ num
	}
	return res
}

// 用hash表记录元素的情况

// time: O(n) space: O(n)
// 执行用时: 28 ms 内存消耗: 6.8 MB
func singleNumberByHash(nums []int) int {
	hash := make(map[int]int)
	for _, v := range nums {
		if c, ok := hash[v]; ok {
			hash[v] = c + 1
		} else {
			hash[v] = 1
		}
	}

	for k, v := range hash {
		if v == 1 {
			return k
		}
	}
	return 0
}
