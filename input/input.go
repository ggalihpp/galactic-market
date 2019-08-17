package input

import (
	"fmt"
	"regexp"
	"strings"
)

// Msg - will check pattern
func Msg(originMsg string) string {
	msg := strings.ToUpper(originMsg)
	msg = strings.TrimSuffix(msg, "\n") // remove new line
	msg = strings.TrimSpace(msg)
	m := strings.Split(msg, " ")

	//// STORED CUSTOM ROMAN RULES
	if len(m) == 3 {
		return addCustomRomanRules(m)
	}

	//// SAVING METALS VALUE
	matched, err := regexp.MatchString("^([A-Za-z\\s]*) ([A-Z]+[A-Za-z]+) IS ([0-9]+) CREDITS$", msg)
	if err != nil {
		fmt.Println("RERR:: ", err.Error())
		return haveNoIdea()
	}

	if matched {
		return insertMetals(msg)
	}

	//// TRANSACTION
	if strings.Contains(msg, "HOW MUCH IS") {
		return getRomansVal(msg)
	}

	if strings.Contains(msg, "HOW MANY CREDITS IS") {
		return getMetalsValue(msg)
	}

	return haveNoIdea()
}

func haveNoIdea() string {
	return "I have no idea what you are talking about"
}
