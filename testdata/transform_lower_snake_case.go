package main

import "fmt"

type TransformLowerSnakeCase int

const (
	TransformLowerSnakeCaseOne TransformLowerSnakeCase = iota
	TransformLowerSnakeCaseTwo
	TransformLowerSnakeCaseThree
)

func main() {
	ck(TransformLowerSnakeCaseOne, "transform_lower_snake_case_one")
	ck(TransformLowerSnakeCaseTwo, "transform_lower_snake_case_two")
	ck(TransformLowerSnakeCaseThree, "transform_lower_snake_case_three")
	ck(-127, "TransformLowerSnakeCase(-127)")
	ck(127, "TransformLowerSnakeCase(127)")
}

func ck(value TransformLowerSnakeCase, str string) {
	if fmt.Sprint(value) != str {
		panic("transform_lower_snake_case.go: " + str)
	}
}
