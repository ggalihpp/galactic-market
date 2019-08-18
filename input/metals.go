package input

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ggalihpp/galactic-market/roman"
	"github.com/ggalihpp/galactic-market/stored"
)

func insertMetals(input string) string {
	var err error
	var mtrl, rom string
	var crdts, mtrlIdx, qty int

	msgs := strings.Split(input, " ")

	for i := range msgs {
		if msgs[i] == "IS" {
			mtrlIdx = i - 1
			mtrl = msgs[i-1]
			crdts, err = strconv.Atoi(msgs[i+1])
			if err != nil {
				return haveNoIdea(err)
			}
		}
	}

	rmns := msgs[0:mtrlIdx]

	for _, v := range rmns {
		if val, ok := stored.RulesRoman.CheckValue(v); ok {
			rom += val
		} else {
			return haveNoIdea(fmt.Errorf("Roman %v not found", v))
		}
	}

	qty, err = roman.Rtoi(rom)
	if err != nil {
		return haveNoIdea(err)
	}

	//// evaluate the value
	value := float64(crdts) / float64(qty)
	stored.MetalsValue[mtrl] = value

	return fmt.Sprintf("%v value updated:: %v", strings.Title(mtrl), value)
}

func getMetalsValue(msg string) string {
	var rom, material string
	var romanVal, materialVal float64

	x := strings.Split(msg, "HOW MANY CREDITS IS")
	xx := strings.TrimSpace(x[1])

	xxx := strings.Split(xx, " ")

	if xxx[len(xxx)-1] != "?" {
		return haveNoIdea(fmt.Errorf("Last index not a question mark"))
	}

	material = xxx[len(xxx)-2]

	materialVal, found := stored.MetalsValue[material]
	if !found {
		return haveNoIdea(fmt.Errorf("Material value for %v not found", material))
	}

	romans := xxx[0 : len(xxx)-2]

	for i := range romans {
		if val, ok := stored.RulesRoman.CheckValue(romans[i]); ok {
			rom += val
		} else {
			if romans[i] != "?" {
				return haveNoIdea(fmt.Errorf("Roman %v not found", romans[i]))
			}
		}
	}

	rv, err := roman.Rtoi(rom)
	if err != nil {
		return haveNoIdea(err)
	}

	romanVal = float64(rv)

	return fmt.Sprintf("%v %v is %v Credits", strings.ToLower(strings.Join(romans, " ")), strings.Title(strings.ToLower(material)), romanVal*materialVal)
}
