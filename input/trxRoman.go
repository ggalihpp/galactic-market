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

	xxx := strings.Split(xx, " ")

	for i := range xxx {
		if val, ok := stored.RulesRoman.CheckValue(xxx[i]); ok {
			rom += val
		} else {
			if xxx[i] != "?" {
				return haveNoIdea()
			}
		}
	}

	val, err := roman.Rtoi(rom)
	if err != nil {
		return haveNoIdea()
	}

	return fmt.Sprintf("%v is %v", strings.Join(xxx, " "), val)
}
