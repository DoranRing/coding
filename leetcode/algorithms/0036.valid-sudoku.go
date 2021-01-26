package algorithms

// 使用哈希表记录已出现元素的坐标，后面的元素检查相等元素的坐标，不符合数独要求的返回结果
// time:O(n), space:O(n)
// 执行用时: 4 ms 内存消耗: 2.9 MB
func isValidSudokuByHash(board [][]byte) bool {
	hash := make(map[byte][][]byte)
	for row := byte(0); row < byte(len(board)); row++ {
		for col := byte(0); col < byte(len(board[row])); col++ {
			val := board[row][col]
			if val == '.' {
				continue
			}
			var points [][]byte
			if existPoints, ok := hash[val]; ok {
				if !valid(existPoints, row, col) {
					return false
				}
				points = existPoints
			}
			points = append(points, []byte{col, row})
			hash[val] = points
		}
	}

	return true
}

func valid(points [][]byte, row, col byte) bool {
	for i := 0; i < len(points); i++ {
		if points[i][0] == col {
			return false
		}
		if points[i][1] == row {
			return false
		}
		// 3*3
		if col/3 == points[i][0]/3 && row/3 == points[i][1]/3 {
			return false
		}
	}
	return true
}
