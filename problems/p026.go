package p026

import (
	"fmt"
)

func findCycleLength(d int) int {
	// Track a list of remainders we have seen.
	rs := make([]int, 0)

	// Assume we are calculating 1/d, for d > 1.
	r := 10
	for i := 1; i < 10000; i++ {
		x := r / d
		r = r - (d * x)
		r *= 10

		// Abort if no remainder.
		if r == 0 {
			return 0
		}
		// Check the existing remainders to see if we've repeated.
		for j, prev := range rs {
			if prev == r {
				// The size is the difference in positions.
				return len(rs) - j
			}
		}
		// No repeat yet. Append and continue to loop.
		rs = append(rs, r)
	}
	fmt.Printf("overflow\n")
	return 0
}

func main() {
	answer := 0
	longest := 0
	for d := 2; d < 1000; d++ {
		cycle := findCycleLength(d)
		if cycle > longest {
			longest = cycle
			answer = d
		}
	}
	fmt.Printf("1/%v has cycle length %v\n", answer, longest)
}
