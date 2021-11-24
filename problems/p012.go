package p012

import (
	"fmt"
	"math"
)

func divisors(n int) int {
	ds := 0

	// Take the sqrt so we check vastly fewer numbers.
	limit := int(math.Sqrt(float64(n)))

	for i := 1; i <= limit; i++ {
		if i*i == n {
			// squares only count once
			ds += 1
		} else if n%i == 0 {
			// other divisors have a mirror image above the sqrt
			ds += 2
		}
	}
	return ds
}

func main() {
	s := 0
	for i := 1; ; i++ {
		s += i
		ds := divisors(s)
		if ds > 500 {
			fmt.Printf("Done! %v has %v divisors\n", s, ds)
			return
		}
	}
}
