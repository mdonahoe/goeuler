package p016

import (
	"fmt"
	"math/big"
)

func main() {
	var i, e = big.NewInt(2), big.NewInt(1000)
	i.Exp(i, e, nil)
	s := i.String()
	sum := 0
	for i := 0; i < len(s); i++ {
		d := int(s[i] - '0')
		sum += d
	}
	fmt.Printf("%v\n", sum)
}
