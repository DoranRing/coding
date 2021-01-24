package algorithms

// 贪心
// time: O(n) space: O(1)
// 执行用时: 4 ms 内存消耗: 3 MB
func maxProfitByGreedy(prices []int) int {
	sum := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			sum += prices[i+1] - prices[i]
		}
	}
	return sum
}
