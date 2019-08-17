package input

import (
	"fmt"
	"strings"

	"github.com/ggalihpp/galactic-market/roman"
	"github.com/ggalihpp/galactic-market/stored"
)

func getMetalsValue(msg string) string {
	var rom, material string
	var romanVal, materialVal float64

	x := strings.Split(msg, "HOW MANY CREDITS IS")
	xx := strings.TrimSpace(x[1])

	xxx := strings.Split(xx, " ")

	if xxx[len(xxx)-1] != "?" {
		return haveNoIdea()
	}

	material = xxx[len(xxx)-2]

	materialVal, found := stored.MetalsValue[material]
	if !found {
		return haveNoIdea()
	}

	romans := xxx[0 : len(xxx)-2]

	for i := range romans {
		if val, ok := stored.RulesRoman.CheckValue(romans[i]); ok {
			rom += val
		} else {
			if romans[i] != "?" {
				return haveNoIdea()
			}
		}
	}

	rv, err := roman.Rtoi(rom)
	if err != nil {
		return haveNoIdea()
	}

	romanVal = float64(rv)

	return fmt.Sprintf("%v %v is %v Credits", strings.ToLower(strings.Join(romans, " ")), strings.Title(strings.ToLower(material)), romanVal*materialVal)
}
