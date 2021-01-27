package p019

import (
	"fmt"
	"time"
)

func main() {
	sundays := 0
	for year := 1901; year <= 2000; year++ {
		for month := 1; month <= 12; month++ {
			date := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
			if date.Weekday() == time.Sunday {
				sundays++
			}
		}
	}
	fmt.Printf("%v Sundays\n", sundays)
}
