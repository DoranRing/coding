package algorithms

import (
	"sort"
)

// 排序两个字符串，然后依次比较每个字符是否相等
// time: O(n + logN) space:O(n)
//执行用时: 12 ms 内存消耗: 3.6 MB
func isAnagramBySort(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sRune := []rune(s)
	sort.Slice(sRune, func(i, j int) bool {
		return sRune[i] < sRune[j]
	})
	tRune := []rune(t)
	sort.Slice(tRune, func(i, j int) bool {
		return tRune[i] < tRune[j]
	})
	for i := 0; i < len(sRune); i++ {
		if sRune[i] != tRune[i] {
			return false
		}
	}

	return true
}

// 先用哈希表记录第一个字符串的字符情况，然后依次与第二个字符串的字符做比较
// time:O(n), space:O(n)
//执行用时: 8 ms 内存消耗: 2.8 MB
func isAnagramByHash(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hash := make(map[rune]int)
	for _, val := range s {
		if count, ok := hash[val]; ok {
			hash[val] = count + 1
		} else {
			hash[val] = 1
		}
	}

	for _, val := range t {
		if count, ok := hash[val]; ok {
			if count <= 0 {
				return false
			}
			hash[val] = count - 1
		} else {
			return false
		}
	}
	return true
}

// 借助元素限定在小写字符，使用a-z共26个元素的数组记录第一个字符串每个字符出现的次数，然后依次与第二个字符串的字符做比较
// time:O(n), space:O(1)
// 执行用时: 0 ms, 内存消耗: 2.7 MB
func isAnagramByArray(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hash := make([]int, 26)
	for _, val := range s {
		count := hash[val-'a']
		hash[val-'a'] = count + 1
	}

	for _, val := range t {
		count := hash[val-'a']
		if count == 0 {
			return false
		}
		hash[val-'a'] = count - 1
	}
	return true
}
