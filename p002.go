package p002

import "fmt"

func main() {
	a := 1
	b := 2
	sum := 2
	for i := 2; ; i++ {
		tmp := a
		a = b
		b = b + tmp
		if b > 4000000 {
			break
		}
		if b%2 == 0 {
			sum += b
		}
	}
	fmt.Printf("%v\n", sum)
}
