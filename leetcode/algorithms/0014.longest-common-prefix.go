package algorithms

//执行用时: 0 ms 内存消耗: 2.3 MB
func longestCommonPrefix(strs []string) string {
	for idx := 0; idx <= 200; idx++ {
		var val uint8
		var init = false
		for _, str := range strs {
			if len(str) <= idx || (init && str[idx] != val) {
				if idx > 0 {
					return str[0:idx]
				} else {
					return ""
				}
			}

			if !init {
				val = str[idx]
				init = true
			}
		}
	}

	return ""
}
