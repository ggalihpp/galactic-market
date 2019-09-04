package input

import (
	"fmt"
	"strings"

	"github.com/ggalihpp/galactic-market/roman"
	"github.com/ggalihpp/galactic-market/stored"
)

func addCustomRomanRules(input []string) string {
	if string(input[1]) == "IS" {
		//// Store Based Roman
		v := input[2]
		if _, ok := roman.Dict[v]; ok {
			stored.RulesRoman[v] = input[0]
			return fmt.Sprintf("Rules roman %v has been updated to %v", v, input[0])
		}
	}

	return haveNoIdea(fmt.Errorf("Unrecognize input"))
}

func getRomansVal(msg string) string {
	var rom string

	x := strings.Split(msg, "HOW MUCH IS")
	xx := strings.TrimSpace(x[1])
	romans := strings.Split(strings.ToLower(xx), " ")

	if romans[len(romans)-1] != "?" {
		return haveNoIdea(fmt.Errorf("Last index not a question mark"))
	}

	for i := range romans {
		if val, ok := stored.RulesRoman.CheckValue(strings.ToUpper(romans[i])); ok {
			rom += val
		} else {
			if romans[i] != "?" {
				return haveNoIdea(fmt.Errorf("Roman %v not found", romans[i]))
			}
		}
	}

	val, err := roman.Rtoi(rom)
	if err != nil {
		return haveNoIdea(err)
	}

	return fmt.Sprintf("%v is %v", strings.Join(romans[0:len(romans)-1], " "), val)
}

func comparison(msg string) string {
	var found bool
	var left, right string

	var leftMtrl float64 = 1
	var rightMtrl float64 = 1

	var ti, isSmaller bool
	msgArr := strings.Split(msg, " ")
	msgArr = msgArr[1:len(msgArr)]

	for i := range msgArr {
		if msgArr[i] == "CREDITS" {
			continue
		}

		if msgArr[i] == "LARGER" || msgArr[i] == "HIGHER" {
			isSmaller = false
			continue
		} else if msgArr[i] == "LOWER" || msgArr[i] == "SMALLER" {
			isSmaller = true
			continue
		}

		if msgArr[i] == "THAN" {
			ti = true
			continue
		}

		//// FROM LEFT TO RIGHT
		if val, ok := stored.RulesRoman.CheckValue(msgArr[i]); ok {
			if ti {
				right += val
			} else {
				left += val
			}
		} else {
			if msgArr[i] != "?" {
				if !ti {
					leftMtrl, found = stored.MaterialsValue[msgArr[i]]
					if !found {
						return haveNoIdea(fmt.Errorf("Material value for %v not found", msgArr[i]))
					}
					continue
				} else {
					rightMtrl, found = stored.MaterialsValue[msgArr[i]]
					if !found {
						return haveNoIdea(fmt.Errorf("Material value for %v not found", msgArr[i]))
					}
					continue
				}
			}
		}
	}

	leftValRom, err := roman.Rtoi(left)
	if err != nil {
		return haveNoIdea(err)
	}

	rightValRom, err := roman.Rtoi(right)
	if err != nil {
		return haveNoIdea(err)
	}

	leftMtrl = leftMtrl * float64(leftValRom)
	rightMtrl = rightMtrl * float64(rightValRom)

	if (leftMtrl < rightMtrl && isSmaller) || (leftMtrl > rightMtrl && !isSmaller) {
		return "yes"
	} else if leftMtrl == rightMtrl {
		return "same value"
	} else {
		return "no"
	}

}
