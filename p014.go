package p014

import "fmt"

func collatz(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return 3*n + 1
}

func main() {
	longestLen := 0
	longestNum := 0
	for i := 1; i < 1000000; i++ {
		length := 1
		for x := i; x != 1; x = collatz(x) {
			length++
		}
		if length > longestLen {
			longestLen = length
			longestNum = i
			fmt.Printf("%v -> %v\n", i, length)
		}
	}
	fmt.Printf("starting num %v had a chain of length %v\n", longestNum, longestLen)
}
