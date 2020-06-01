package main

import (
	"strings"

	iuliia "github.com/mehanizm/iuliia-go"
)

func transcript(s string) string {
	x := iuliia.Wikipedia.Translate(s)
	return strings.ReplaceAll(x, " ", "_")
}
