package p025

import (
	"fmt"
	"math/big"
)

func indexOfFirstNDigitFib(n int) int {
    if n < 2 {
        return 1
    }
	a := big.NewInt(1)
	b := big.NewInt(1)
    index := 2
    for len(b.String()) < n {
        index += 1
        a.Add(a, b)
        a, b = b, a  // swap
    }
    return index
}

func main() {
	index := indexOfFirstNDigitFib(1000)
	fmt.Printf("%v\n", index)
}
