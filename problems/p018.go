package p018

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	s := `
75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23
`

	// Parse
	triangle := make([][]int, 0)
	maxsums := make([][]float64, 0)
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		nums := strings.Split(line, " ")
		triangle = append(triangle, make([]int, len(nums)))
		maxsums = append(maxsums, make([]float64, len(nums)))
		for j, num := range nums {
			x, err := strconv.Atoi(num)
			if err == nil {
				triangle[len(triangle)-1][j] = x
				maxsums[len(maxsums)-1][j] = 0
			}
		}
	}

	// Confirm parse
	for row := 0; row < len(triangle); row++ {
		for col := 0; col < len(triangle[row]); col++ {
			fmt.Printf("%v ", triangle[row][col])
		}
		fmt.Printf("\n")
	}

	// Look at every row
	for i := 0; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			a := 0.0
			b := 0.0
			if i > 0 {
				if j < len(maxsums[i-1]) {
					a = maxsums[i-1][j]
				}
				if j > 0 {
					b = maxsums[i-1][j-1]
				}
			}
			maxsums[i][j] = math.Max(a, b) + float64(triangle[i][j])
		}
	}

	max := 0.0
	final := maxsums[len(maxsums)-1]
	for i := 0; i < len(final); i++ {
		if final[i] > max {
			max = final[i]
		}
	}

	fmt.Printf("%v\n", max)
}
