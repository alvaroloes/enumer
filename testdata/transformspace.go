package main

import "fmt"

type CamelCaseValue int

const (
	CamelCaseValueOne CamelCaseValue = iota
	CamelCaseValueTwo
	CamelCaseValueThree
)

func main() {
	ck(CamelCaseValueOne, "camel case value one")
	ck(CamelCaseValueTwo, "camel case value two")
	ck(CamelCaseValueThree, "camel case value three")
	ck(-127, "CamelCaseValue(-127)")
	ck(127, "CamelCaseValue(127)")
}

func ck(value CamelCaseValue, str string) {
	if fmt.Sprint(value) != str {
		panic("transformspace.go: " + str)
	}
}
