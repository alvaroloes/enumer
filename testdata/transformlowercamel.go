package main

import "fmt"

type CamelCaseValue int

const (
	CamelCaseValueOne CamelCaseValue = iota
	CamelCaseValueTwo
	CamelCaseValueThree
)

func main() {
	ck(CamelCaseValueOne, "camelCaseValueOne")
	ck(CamelCaseValueTwo, "camelCaseValueTwo")
	ck(CamelCaseValueThree, "camelCaseValueThree")
	ck(-127, "CamelCaseValue(-127)")
	ck(127, "CamelCaseValue(127)")
}

func ck(value CamelCaseValue, str string) {
	if fmt.Sprint(value) != str {
		panic("transformlowercamel.go: " + str)
	}
}
