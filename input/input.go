package input

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"

	"github.com/ggalihpp/galactic-market/roman"
	"github.com/ggalihpp/galactic-market/stored"
)

// Msg - will check pattern
func Msg(msg string) {
	msg = strings.ToUpper(msg)
	msg = strings.TrimSuffix(msg, "\n") // remove new line
	m := strings.Split(msg, " ")

	//// STORED CUSTOM ROMAN RULES
	if len(m) == 3 {
		if string(m[1]) == "IS" {
			//// Store Based Roman
			v := m[2]
			if _, ok := roman.Dict[v]; ok {
				stored.RulesRoman[v] = m[0]
				fmt.Printf("Rules roman %v has been updated to %v\n", v, m[0])
				return
			}
		}
	}

	//// SAVING METALS VALUE
	switch {
	case strings.Contains(msg, "SILVER IS"):
		var quantity int
		var rom string

		x := strings.Split(msg, "SILVER IS")
		x[0] = strings.TrimSpace(x[0])
		x[1] = strings.TrimSpace(x[1])
		crRaw := strings.Split(x[1], " ")
		qtyRaw := strings.Split(x[0], " ")

		spew.Dump(crRaw)
		spew.Dump(qtyRaw)
		cr, err := strconv.Atoi(crRaw[0])
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		//// evaluate the quantity
		for _, v := range qtyRaw {
			if val, ok := stored.RulesRoman.CheckValue(v); ok {
				rom += val
			} else {
				fmt.Println("RULES NOT FOUND!!!::: ", v)
				return
			}
		}

		quantity, err = roman.Rtoi(rom)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		//// evaluate the value
		value := float64(cr) / float64(quantity)
		stored.MetalsValue["SILVER"] = value
		return
	case strings.Contains(msg, "GOLD IS"):
		var quantity int
		var rom string

		x := strings.Split(msg, "GOLD IS")
		x[0] = strings.TrimSpace(x[0])
		x[1] = strings.TrimSpace(x[1])
		crRaw := strings.Split(x[1], " ")
		qtyRaw := strings.Split(x[0], " ")

		spew.Dump(crRaw)
		spew.Dump(qtyRaw)
		cr, err := strconv.Atoi(crRaw[0])
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		//// evaluate the quantity
		for _, v := range qtyRaw {
			if val, ok := stored.RulesRoman.CheckValue(v); ok {
				rom += val
			} else {
				fmt.Println("RULES NOT FOUND!!!::: ", v)
				return
			}
		}

		quantity, err = roman.Rtoi(rom)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		//// evaluate the value
		value := float64(cr) / float64(quantity)
		stored.MetalsValue["GOLD"] = value

		return
	case strings.Contains(msg, "IRON IS"):
		var quantity int
		var rom string

		x := strings.Split(msg, "IRON IS")
		x[0] = strings.TrimSpace(x[0])
		x[1] = strings.TrimSpace(x[1])
		crRaw := strings.Split(x[1], " ")
		qtyRaw := strings.Split(x[0], " ")

		spew.Dump(crRaw)
		spew.Dump(qtyRaw)
		cr, err := strconv.Atoi(crRaw[0])
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		//// evaluate the quantity
		for _, v := range qtyRaw {
			if val, ok := stored.RulesRoman.CheckValue(v); ok {
				rom += val
			} else {
				fmt.Println("RULES NOT FOUND")
				return
			}
		}

		quantity, err = roman.Rtoi(rom)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		//// evaluate the value
		value := float64(cr) / float64(quantity)
		stored.MetalsValue["IRON"] = value

		return
	}

	//// TRANSACTION

	if strings.Contains(msg, "HOW MUCH IS") {
		var rom string

		x := strings.Split(msg, "HOW MUCH IS")
		xx := strings.TrimSpace(x[1])

		xxx := strings.Split(xx, " ")

		for i := range xx {
			if val, ok := stored.RulesRoman.CheckValue(xxx[i]); ok {
				rom += val
			} else {
				if xxx[i] != "?" {
					fmt.Println("RULES NOT FOUND")
					return
				}
			}
		}

	}

}
