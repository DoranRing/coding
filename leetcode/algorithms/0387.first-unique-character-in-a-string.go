package algorithms

// 将a-z字符转化成0-25的下标，
// time:O(n), space:O(1)
// 执行用时: 8 ms 内存消耗: 5.2 MB
func firstUniqCharByArray(s string) int {
	arr := make([]int, 'z'-'a'+1)
	for _, c := range s {
		count := arr[c-'a']
		arr[c-'a'] = count + 1
	}

	for idx, c := range s {
		if arr[c-'a'] == 1 {
			return idx
		}
	}
	return -1
}

// 通过哈希表记录字符出现情况
// time:O(n), space:O(1)
// 执行用时: 52 ms 内存消耗: 5.3 MB
func firstUniqCharByHash(s string) int {
	hash := make(map[int32]int)
	for _, c := range s {
		count := hash[c]
		hash[c] = count + 1
	}

	for idx, c := range s {
		if hash[c] == 1 {
			return idx
		}
	}
	return -1
}
