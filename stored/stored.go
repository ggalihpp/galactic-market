package stored

type rulesRoman map[string]string

// CheckValue - will check value
func (m *rulesRoman) CheckValue(val string) (key string, ok bool) {
	for k, v := range *m {
		if v == val {
			key = k
			ok = true
			return
		}
	}
	return
}

// RulesRoman - Contains custom rules for roman number (The default will set)
var RulesRoman = rulesRoman{
	"I": "I",
	"V": "V",
	"X": "X",
	"L": "L",
	"C": "C",
	"D": "D",
	"M": "M",
}

// MaterialsValue -
var MaterialsValue = map[string]float64{
	"GOLD":   0,
	"SILVER": 0,
	"IRON":   0,
}
