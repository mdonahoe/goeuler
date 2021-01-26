package p017

import "fmt"

func teens(n int) string {
	switch n {
	case 10:
		return "ten"
	case 11:
		return "eleven"
	case 12:
		return "twelve"
	case 13:
		return "thirteen"
	case 14:
		return "fourteen"
	case 15:
		return "fifteen"
	case 16:
		return "sixteen"
	case 17:
		return "seventeen"
	case 18:
		return "eighteen"
	case 19:
		return "nineteen"
	}
	return ""
}

func letters(n int) string {
	if n == 1000 {
		return "onethousand"
	}
	if n > 99 {
		if n%100 == 0 {
			return letters(n/100) + "hundred"
		} else {
			return letters(n/100) + "hundredand" + letters(n%100)
		}
	}
	if n > 9 {
		ones := letters(n % 10)
		switch n / 10 {
		case 0:
			return ones
		case 1:
			return teens(n)
		case 2:
			return "twenty" + ones
		case 3:
			return "thirty" + ones
		case 4:
			return "forty" + ones
		case 5:
			return "fifty" + ones
		case 6:
			return "sixty" + ones
		case 7:
			return "seventy" + ones
		case 8:
			return "eighty" + ones
		case 9:
			return "ninety" + ones
		}
	}
	switch n {
	case 0:
		return ""
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	default:
		return "oops"
	}
	return ""
}

func main() {
	count := 0
	for i := 1; i <= 1000; i++ {
		ls := letters(i)
		count += len(ls)
		fmt.Printf("%v\n", ls)
	}
	fmt.Printf("%v\n", count)
}
