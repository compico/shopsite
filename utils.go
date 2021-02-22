package main

import (
	"math"
	"strings"

	"github.com/mehanizm/iuliia-go"
)

func round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

func transcript(s string) string {
	if s[0] == ' ' {
		s = s[1:]
		transcript(s)
	}
	x := iuliia.Wikipedia.Translate(s)
	return strings.ReplaceAll(x, " ", "_")
}
