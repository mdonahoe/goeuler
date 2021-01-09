package p006

import (
	"fmt"
)

func main() {
	sum := 0
	squares := 0
	for i := 1; i <= 100; i++ {
		sum += i
		squares += i * i
	}
	diff := sum*sum - squares
	fmt.Printf("%v\n", diff)
}
