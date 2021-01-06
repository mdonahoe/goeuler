package p005

import (
	"fmt"
)

func main() {
	x := 0
	for i := 1; ; i++ {
		success := true
		for j := 2; j < 20; j++ {
			if i%j > 0 {
				success = false
				break
			}
		}
		if success {
			x = i
			break
		}
	}
	fmt.Printf("%v\n", x)
}
