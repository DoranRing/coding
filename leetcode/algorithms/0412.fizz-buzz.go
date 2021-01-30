package algorithms

import "strconv"

//执行用时: 4 ms 内存消耗: 3.4 MB
func fizzBuzz(n int) []string {
	res := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		var item string
		if i%3 == 0 && i%5 == 0 {
			item = "FizzBuzz"
		} else if i%3 == 0 {
			item = "Fizz"
		} else if i%5 == 0 {
			item = "Buzz"
		} else {
			item = strconv.Itoa(i)
		}
		res = append(res, item)
	}
	return res
}
