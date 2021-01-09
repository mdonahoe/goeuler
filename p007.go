package p007

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n%2 == 0 {
		return false
	}
	root := int(math.Sqrt(float64(n)))
	for i := 3; i <= root; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	count := 1
	for n := 3; ; n += 2 {
		if isPrime(n) {
			count += 1
			fmt.Printf("%v -> %v\n", count, n)
			if count == 10001 {
				return
			}
		}
	}
}
