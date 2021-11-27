package p042

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func isTriangle(x int) bool {
	n := 0
	for {
		t := ((n + 1) * n) / 2
		if t == x {
			return true
		}
		if t > x {
			return false
		}
		n += 1
	}
	return false
}

func main() {
	filename := "p042_words.txt"

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	words := strings.Split(string(data), ",")

	count := 0
	for _, word := range words {
		s := 0
		for _, c := range word {
			if c == '"' {
				continue
			}
			s += int(c-'A') + 1
		}
		if isTriangle(s) {
			count += 1
		}
	}

	fmt.Println(count)
}
