package p028

import (
	"fmt"
)

func main() {
	answer := 1
	n := 1001
	// The levels of the spiral are odd in length
	for i := 3; i <= n; i += 2 {
		// Each level of the spirals adds this much to the total sum.
		answer += (4 * i * i) - (6 * i) + 6
	}
	fmt.Printf("%v\n", answer)
}
