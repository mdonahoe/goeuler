package p003
// TODO(matt): figure out how to run all these problems...

import (
	"fmt"
	"math"
)

func firstFactor(n int) int {
	if n%2 == 0 {
		return 2
	}
	root := int(math.Sqrt(float64(n)))
	for i := 3; i <= root; i += 2 {
		if n%i == 0 {
			return i
		}
	}
	return n
}

func largestPrimeFactor(n int) int {
	for p := 1; p != n; p = firstFactor(n) {
		fmt.Printf("n=%v, p=%v\n", n, p)
		n = n / p
	}
	return n
}

func main() {
	x := largestPrimeFactor(600851475143)
	fmt.Printf("largest = %v\n", x)
}
