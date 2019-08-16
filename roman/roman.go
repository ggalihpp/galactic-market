package roman

import (
	"fmt"
	"strings"
)

// Dict - Contains roman numeral value
var Dict = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// Rtoi - Convert Roman numeral to numeric
func Rtoi(r string) (res int, err error) {
	var lastNum int

	r = strings.TrimSuffix(r, "\n") // remove new line
	r = reverse(r)                  // reverse the string

	err = checkRomRepeated(r) // check any repeated char
	if err != nil {
		return
	}

	for _, v := range r {
		if val, ok := Dict[strings.ToUpper(string(v))]; ok {
			res = checkNum(val, lastNum, res)
			lastNum = val
		} else {
			err = fmt.Errorf("WRONG VALUE")
			return
		}
	}

	return
}

func checkRomRepeated(r string) error {
	r = strings.ToUpper(r)

	switch {
	case strings.Contains(r, "IIII"):
		return fmt.Errorf("I Repeated more than 3 times")
	case strings.Contains(r, "XXXX"):
		return fmt.Errorf("X Repeated more than 3 times")
	case strings.Contains(r, "CCCC"):
		return fmt.Errorf("C Repeated more than 3 times")
	case strings.Contains(r, "MMMM"):
		return fmt.Errorf("M Repeated more than 3 times")
	case strings.Contains(r, "DD"):
		return fmt.Errorf("D Repeated")
	case strings.Contains(r, "LL"):
		return fmt.Errorf("L Repeated")
	case strings.Contains(r, "VV"):
		return fmt.Errorf("V Repeated")
	default:
		return nil
	}

}

func checkNum(n, lastNum, num int) int {
	if lastNum > n {
		return num - n
	}

	return num + n
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
