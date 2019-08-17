package input

import (
	"fmt"

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

	return haveNoIdea()
}
