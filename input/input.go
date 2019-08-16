package input

import (
	"fmt"
	"strings"
)

// Msg - will check pattern
func Msg(msg string) {
	msg = strings.ToUpper(msg)
	m := strings.Split(msg, " ")

	if len(m) == 3 {
		if string(m[1]) == "IS" {
			//// Store Based Roman
			fmt.Println(m)
		}
	}

}

var romanDict map[string]struct{}
