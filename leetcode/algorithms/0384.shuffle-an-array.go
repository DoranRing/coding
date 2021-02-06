package algorithms

import (
	"math/rand"
	"time"
)

// 洗牌算法
//执行用时: 292 ms 内存消耗: 9.3 MB
type Solution struct {
	source []int

	shuffle []int
}

func NewSolution(nums []int) Solution {
	shuffle := make([]int, len(nums), len(nums))
	for i := 0; i < len(nums); i++ {
		shuffle[i] = nums[i]
	}

	return Solution{
		nums,
		shuffle,
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.source
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(this.shuffle); i++ {
		random := rand.Intn(len(this.shuffle)-i) + i
		this.shuffle[i], this.shuffle[random] = this.shuffle[random], this.shuffle[i]
	}
	return this.shuffle
}
