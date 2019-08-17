package input

import (
	"fmt"
	"strings"

	"github.com/ggalihpp/galactic-market/roman"
	"github.com/ggalihpp/galactic-market/stored"
)

func getRomansVal(msg string) string {
	var rom string

	x := strings.Split(msg, "HOW MUCH IS")
	xx := strings.TrimSpace(x[1])
	romans := strings.Split(strings.ToLower(xx), " ")

	if romans[len(romans)-1] != "?" {
		return haveNoIdea()
	}

	for i := range romans {
		if val, ok := stored.RulesRoman.CheckValue(strings.ToUpper(romans[i])); ok {
			rom += val
		} else {
			if romans[i] != "?" {
				return haveNoIdea()
			}
		}
	}

	val, err := roman.Rtoi(rom)
	if err != nil {
		return haveNoIdea()
	}

	return fmt.Sprintf("%v is %v", strings.Join(romans[0:len(romans)-1], " "), val)
}
