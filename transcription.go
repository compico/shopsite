package main

import (
	"strings"

	iuliia "github.com/mehanizm/iuliia-go"
)

func transcript(s string) string {
	if s[0] == ' ' {
		s = s[1:]
		transcript(s)
	}
	x := iuliia.Wikipedia.Translate(s)
	return strings.ReplaceAll(x, " ", "_")
}
