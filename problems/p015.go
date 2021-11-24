package p015

import "fmt"

func main() {
	n := 20
	var lattice = make([][]int, n+1)
	for i := 0; i < len(lattice); i++ {
		lattice[i] = make([]int, n+1)
	}
	lattice[n][n] = 1
	for i := n; i >= 0; i-- {
		for j := n; j >= 0; j-- {
			if j != n {
				lattice[i][j] += lattice[i][j+1]
			}
			if i != n {
				lattice[i][j] += lattice[i+1][j]
			}
		}
	}

	fmt.Printf("%v paths for a %vx%v grid\n", lattice[0][0], n, n)
}
