package p031

import (
	"fmt"
)

var coinValues []int

// Compute the number of ways to reach target from the total so far.
func numWays(total int, startIndex int, target int) int {
	if total > target {
		return 0
	}
	if total == target {
		return 1
	}
	num := 0
	for i := startIndex; i < len(coinValues); i++ {
		num += numWays(total+coinValues[i], i, target)
	}
	return num
}

func main() {
	coinValues = []int{1, 2, 5, 10, 20, 50, 100, 200}
	answer := numWays(0, 0, 200)
	fmt.Println(answer)
}
