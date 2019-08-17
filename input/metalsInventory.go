package input

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ggalihpp/galactic-market/roman"
	"github.com/ggalihpp/galactic-market/stored"
)

func insertMetals(input string) string {
	switch {
	case strings.Contains(input, "SILVER IS"):
		var quantity int
		var rom string

		x := strings.Split(input, "SILVER IS")
		x[0] = strings.TrimSpace(x[0])
		x[1] = strings.TrimSpace(x[1])
		crRaw := strings.Split(x[1], " ")
		qtyRaw := strings.Split(x[0], " ")

		cr, err := strconv.Atoi(crRaw[0])
		if err != nil {

			return haveNoIdea()
		}

		//// evaluate the quantity
		for _, v := range qtyRaw {
			if val, ok := stored.RulesRoman.CheckValue(v); ok {
				rom += val
			} else {

				return haveNoIdea()
			}
		}

		quantity, err = roman.Rtoi(rom)
		if err != nil {

			return haveNoIdea()
		}

		//// evaluate the value
		value := float64(cr) / float64(quantity)
		stored.MetalsValue["SILVER"] = value
		return fmt.Sprintf("Silver value updated:: %v", value)
	case strings.Contains(input, "GOLD IS"):
		var quantity int
		var rom string

		x := strings.Split(input, "GOLD IS")
		x[0] = strings.TrimSpace(x[0])
		x[1] = strings.TrimSpace(x[1])
		crRaw := strings.Split(x[1], " ")
		qtyRaw := strings.Split(x[0], " ")

		cr, err := strconv.Atoi(crRaw[0])
		if err != nil {

			return haveNoIdea()
		}

		//// evaluate the quantity
		for _, v := range qtyRaw {
			if val, ok := stored.RulesRoman.CheckValue(v); ok {
				rom += val
			} else {

				return haveNoIdea()
			}
		}

		quantity, err = roman.Rtoi(rom)
		if err != nil {

			return haveNoIdea()
		}

		//// evaluate the value
		value := float64(cr) / float64(quantity)
		stored.MetalsValue["GOLD"] = value

		return fmt.Sprintf("GOLD value updated:: %v", value)
	case strings.Contains(input, "IRON IS"):
		var quantity int
		var rom string

		x := strings.Split(input, "IRON IS")
		x[0] = strings.TrimSpace(x[0])
		x[1] = strings.TrimSpace(x[1])
		crRaw := strings.Split(x[1], " ")
		qtyRaw := strings.Split(x[0], " ")

		cr, err := strconv.Atoi(crRaw[0])
		if err != nil {

			return haveNoIdea()
		}

		//// evaluate the quantity
		for _, v := range qtyRaw {
			if val, ok := stored.RulesRoman.CheckValue(v); ok {
				rom += val
			} else {

				return haveNoIdea()
			}
		}

		quantity, err = roman.Rtoi(rom)
		if err != nil {

			return haveNoIdea()
		}

		//// evaluate the value
		value := float64(cr) / float64(quantity)
		stored.MetalsValue["IRON"] = value

		return fmt.Sprintf("IRON value updated:: %v", value)
	}

	return haveNoIdea()
}
