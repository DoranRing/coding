package algorithms

// time:O(26*n),space:O(26)
// 执行用时: 0 ms 内存消耗: 2.6 MB
func checkInclusionBySlidingWindow(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	sourceBucket := make([]int, 26, 26)
	targetBucket := make([]int, 26, 26)
	for _, c := range s1 {
		targetBucket[c-'a']++
	}

	// 初始化窗口大小
	l, r := -1, -1
	for r < len(s1)-2 {
		r++
		sourceBucket[s2[r]-'a']++
	}
	// 滑动窗口
	for r < len(s2)-1 {
		r++
		sourceBucket[s2[r]-'a']++
		if l >= 0 {
			sourceBucket[s2[l]-'a']--
		}
		l++
		if bucketEqual(sourceBucket, targetBucket) {
			return true
		}
	}

	return false
}

func bucketEqual(sourceBucket, targetBucket []int) bool {
	for i := 0; i < len(sourceBucket); i++ {
		if sourceBucket[i] != targetBucket[i] {
			return false
		}
	}
	return true
}
