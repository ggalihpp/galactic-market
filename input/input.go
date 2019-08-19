package input

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/ggalihpp/galactic-market/stored"
)

// Msg - will check pattern
func Msg(originMsg string) string {
	msg := strings.ToUpper(originMsg)
	msg = strings.TrimSuffix(msg, "\n") // remove new line
	msg = strings.TrimSpace(msg)
	m := strings.Split(msg, " ")

	if strings.Contains(msg, "LIST") {
		return list(msg)
	}

	//// STORED CUSTOM ROMAN RULES
	if len(m) == 3 {
		return addCustomRomanRules(m)
	}

	//// SAVING METALS VALUE
	matched, err := regexp.MatchString("^([A-Za-z\\s]*) ([A-Z]+[A-Za-z]+) IS ([0-9]+) CREDITS$", msg)
	if err != nil {
		return haveNoIdea(err)
	}

	if matched {
		return insertMaterial(msg)
	}

	//// TRANSACTION
	if strings.Contains(msg, "HOW MUCH IS") {
		return getRomansVal(msg)
	}

	if strings.Contains(msg, "HOW MANY CREDITS IS") {
		return getMaterialValue(msg)
	}

	return haveNoIdea(fmt.Errorf("Unrecognize input"))
}

func haveNoIdea(errInput error) string {
	var debugMode bool
	dm, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		debugMode = false
	} else {
		debugMode = dm
	}

	if debugMode {
		if errInput != nil {
			return fmt.Sprintf("DEBUGMODE:: %v", errInput.Error())
		}
	}

	return "I have no idea what you are talking about"

}

func list(msg string) string {
	switch msg {
	case "LIST MATERIALS":
		var mtrls string

		for k, v := range stored.MaterialsValue {
			mtrls += fmt.Sprintf("%v -> %v Credits\n", k, v)
		}

		return mtrls
	case "LIST ROMANS":
		var rmns string

		for k, v := range stored.RulesRoman {
			rmns += fmt.Sprintf("%v -> %v\n", k, v)
		}

		return rmns
	}
	return "Insert command \"List materials\" or \"List romans\""
}
