package p004

import (
	"fmt"
	"strconv"
)

func isPalindrome(n int) bool {
	// convert to string
	s := strconv.Itoa(n)

	// compare chars from the outside in.
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	// they all match!
	return true
}

func main() {
	largest := 0
	for i := 1; i <= 999; i++ {
		for j := 1; j <= 999; j++ {
			n := i * j
			if n > largest && isPalindrome(n) {
				largest = n
			}
		}
	}
	fmt.Printf("%v\n", largest)
}
