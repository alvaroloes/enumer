package main

import "fmt"

type TransformKebabCase int

const (
	TransformKebabCaseOne TransformKebabCase = iota
	TransformKebabCaseTwo
	TransformKebabCaseThree
)

func main() {
	ck(TransformKebabCaseOne, "transform-kebab-case-one")
	ck(TransformKebabCaseTwo, "transform-kebab-case-two")
	ck(TransformKebabCaseThree, "transform-kebab-case-three")
	ck(-127, "TransformKebabCase(-127)")
	ck(127, "TransformKebabCase(127)")
}

func ck(value TransformKebabCase, str string) {
	if fmt.Sprint(value) != str {
		panic("transform_kebab_case.go: " + str)
	}
}
