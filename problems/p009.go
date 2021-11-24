package p009

import "fmt"

func main() {
	for a := 1; a < 1000; a++ {
		for b := a; a+b < 1000; b++ {
			c := 1000 - a - b
			if a*a+b*b == c*c {
				fmt.Printf("a = %v, b = %v, c = %v\n", a, b, c)
				fmt.Printf("abc = %v\n", a*b*c)
				return
			}
		}
	}
	fmt.Printf("Not found\n")
}
