package p010

import "fmt"

// Find the sum of all the primes below two million.
func main() {
	// Create a sieve
	n := 2000000
	nums := make([]bool, n)
	psum := 0
	pcount := 0
	for index := 2; index < n; index++ {
		if !nums[index] {
			// This is a prime number. Increment and sum.
			pcount++
			psum += index
		}
		for x := 2 * index; x < n; x = x + index {
			// Mark the multiples of index as non-prime
			nums[x] = true
		}
	}
	fmt.Printf("The sum of the %v primes under %v is %v\n", pcount, n, psum)
}
