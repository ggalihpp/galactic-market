package input

import (
	"fmt"
	"strings"

	"github.com/ggalihpp/galactic-market/roman"
	"github.com/ggalihpp/galactic-market/stored"
)

// Msg - will check pattern
func Msg(msg string) {
	msg = strings.ToUpper(msg)
	msg = strings.TrimSuffix(msg, "\n") // remove new line

	m := strings.Split(msg, " ")

	if len(m) == 3 {
		if string(m[1]) == "IS" {
			//// Store Based Roman
			v := m[2]
			if _, ok := roman.Dict[v]; ok {
				stored.RulesRoman[v] = m[0]
				fmt.Printf("Rules roman %v has been updated to %v\n", v, m[0])
			}
		}
	}

}
