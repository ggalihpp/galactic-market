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
	"glob glob glob Dirt is 7500 Credits",
	"pish Wood is 1000 Credits",
}

var transactionInput = map[string]string{ // INPUT:EXPECTED RESPONSE
	"how much is pish tegj glob glob ?":                                       "pish tegj glob glob is 42",
	"how many Credits is glob prok Silver ?":                                  "glob prok Silver is 68 Credits",
	"how many Credits is glob prok Gold ?":                                    "glob prok Gold is 57800 Credits",
	"how many Credits is glob prok Iron ?":                                    "glob prok Iron is 782 Credits",
	"how many Credits is glob Dirt ?":                                         "glob Dirt is 2500 Credits",
	"how many Credits is glob glob glob Wood ?":                               "glob glob glob Wood is 300 Credits",
	"how much wood could a woodchuck chuck if a woodchuck could chuck wood ?": "I have no idea what you are talking about",
	"is glob glob glob credits larger than glob glob":                         "yes",
	"is pish smaller than prok":                                               "no",
	"is glob silver smaller than glob gold":                                   "yes",
	"is glob silver larger than glob gold":                                    "no",
}

func Test(t *testing.T) {
	//// INPUT INVENTORY
	for _, v := range inventoryInput {
		input.Msg(v)
	}

	for k, v := range transactionInput {
		res := input.Msg(k)
		if res != v {
			t.Errorf("Wrong response for \"%v\" \nExpected: %v \n Got: %v", k, v, res)
		}
	}
}
