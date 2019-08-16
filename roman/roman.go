package roman

import (
	"fmt"
	"strings"
)

// Rtoi - Convert Roman numeral to numeric
func Rtoi(r string) (res int, err error) {

	for _, v := range r {
		var lastNum int

		switch strings.ToUpper(string(v)) {
		case "I":
			res = checkNum(1, lastNum, res)
			lastNum = 1
		case "V":
			res = checkNum(5, lastNum, res)
			lastNum = 5
		case "X":
			res = checkNum(10, lastNum, res)
			lastNum = 10
		case "L":
			res = checkNum(50, lastNum, res)
			lastNum = 50
		case "C":
			res = checkNum(100, lastNum, res)
			lastNum = 100
		case "D":
			res = checkNum(500, lastNum, res)
			lastNum = 500
		case "M":
			res = checkNum(1000, lastNum, res)
			lastNum = 1000
		default:
			err = fmt.Errorf("THEHELL!!!??")
			return
		}
	}

	return
}

func checkNum(n, lastNum, num int) int {
	if lastNum > n {
		return num - n
	}

	return num + n
}
