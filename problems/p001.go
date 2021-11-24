package p001

import "fmt"

func main() {
	total := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}
	fmt.Printf("answer = %v\n", total)
}
