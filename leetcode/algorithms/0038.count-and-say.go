package algorithms

import (
	"strconv"
	"strings"
)

// 执行用时: 0 ms 内存消耗: 6.2 MB
func countAndSay(n int) string {
	res := "1"
	for i := 1; i < n; i++ {
		res = gen(res)
	}
	return res
}

func gen(s string) string {
	res := make([]string, 0)
	for i := 0; i < len(s); {
		j := i
		for ; j < len(s); j++ {
			if s[i] != s[j] {
				break
			}
		}
		res = append(res, strconv.Itoa(j-i), string(s[i]))
		i = j
	}

	return strings.Join(res, "")
}
