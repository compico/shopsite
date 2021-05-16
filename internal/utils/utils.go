package utils

import (
	"math"
	"strings"

	"github.com/mehanizm/iuliia-go"
)

func Round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

func Transcript(s string) string {
	if s[0] == ' ' {
		s = s[1:]
		Transcript(s)
	}
	x := iuliia.Wikipedia.Translate(s)
	return strings.ReplaceAll(x, " ", "_")
}
