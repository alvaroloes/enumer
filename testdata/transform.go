package main

import "fmt"

type CamelCaseValue int

const (
	CamelCaseValueOne CamelCaseValue = iota
	CamelCaseValueTwo
	CamelCaseValueThree
)

func main() {
	ck(CamelCaseValueOne, "camel_case_value_one")
	ck(CamelCaseValueTwo, "camel_case_value_two")
	ck(CamelCaseValueThree, "camel_case_value_three")
	ck(-127, "CamelCaseValue(-127)")
	ck(127, "CamelCaseValue(127)")
}

func ck(value CamelCaseValue, str string) {
	if fmt.Sprint(value) != str {
		panic("transform.go: " + str)
	}
}
