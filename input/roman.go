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
