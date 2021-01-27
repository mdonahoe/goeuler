package p020

import (
	"fmt"
	"math/big"
)

func fact(n *big.Int) *big.Int {
	one := big.NewInt(1)
	if n.Cmp(one) <= 0 {
		return one
	}
	a := big.NewInt(0)
	b := fact(a.Sub(n, one))
	return n.Mul(n, b)
}

// Find the sum of the digits in the number 100!
func main() {
	z := fact(big.NewInt(100))
	s := z.String()
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += int(s[i] - '0')
	}
	fmt.Printf("%v\n", sum)
}
