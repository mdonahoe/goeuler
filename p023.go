package p023

import "fmt"
import "math"

func summedDivisors(n int) int {
	limit := int(math.Sqrt(float64(n)))
	sum := 1
	for i := 2; i <= limit; i++ {
		if i*i == n {
			sum += i
		} else if n%i == 0 {
			sum += i
			sum += n / i
		}
	}
	return sum
}

func main() {
	n := 28123
	abundant := make([]int, 0)
	for i := 1; i < n; i++ {
		s := summedDivisors(i)
		if s > i {
			abundant = append(abundant, i)
		}
	}
	m := make(map[int]bool)
	for _, a := range abundant {
		for _, b := range abundant {
			m[a+b] = true
		}
	}
	sum := 0
	for i := 0; i < n; i++ {
		if !m[i] {
			sum += i
		}
	}
	fmt.Printf("%v\n", sum)
}

