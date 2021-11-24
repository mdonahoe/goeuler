package p021

import (
	"fmt"
	"math"
)

func divisors(n int) int {
	sum := 1
	stop := int(math.Sqrt(float64(n)))
	for i := 2; i <= stop; i++ {
		if n%i == 0 {
			d := n / i
			sum += d + i
		}
	}
	return sum
}

func main() {
	n := 10000
	sum := 0
	for a := 2; a < n; a++ {
		b := divisors(a)
		c := divisors(b)
		if a == c && a != b {
			fmt.Printf("d(%v) == d(%v)\n", a, b)
			sum += a
		}
	}
	fmt.Printf("%v\n", sum)
}
