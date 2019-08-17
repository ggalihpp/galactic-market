package main

import (
	"testing"

	"github.com/ggalihpp/galactic-market/input"
)

var inventoryInput = []string{
	"glob is I",
	"prok is V",
	"pish is X",
	"tegj is L",
	"glob glob Silver is 34 Credits",
	"glob prok Gold is 57800 Credits",
	"pish pish Iron is 3910 Credits ",
}

var transactionInput = map[string]string{ // INPUT:RESPONSE
	"how much is pish tegj glob glob ?":                                       "pish tegj glob glob is 42",
	"how many Credits is glob prok Silver ?":                                  "glob prok Silver is 68 Credits",
	"how many Credits is glob prok Gold ?":                                    "glob prok Gold is 57800 Credits",
	"how many Credits is glob prok Iron ?":                                    "glob prok Iron is 782 Credits",
	"how much wood could a woodchuck chuck if a woodchuck could chuck wood ?": "I have no idea what you are talking about",
}

func Test(t *testing.T) {
	x := "aspkpgpkp"

	tt := input.Msg(x)

	if tt != "I have no idea what you are talking about" {
		t.Error()
	}
}
