package main

import "fmt"

type TransformUpperSnakeCase int

const (
	TransformUpperSnakeCaseOne TransformUpperSnakeCase = iota
	TransformUpperSnakeCaseTwo
	TransformUpperSnakeCaseThree
)

func main() {
	ck(TransformUpperSnakeCaseOne, "TRANSFORM_UPPER_SNAKE_CASE_ONE")
	ck(TransformUpperSnakeCaseTwo, "TRANSFORM_UPPER_SNAKE_CASE_TWO")
	ck(TransformUpperSnakeCaseThree, "TRANSFORM_UPPER_SNAKE_CASE_THREE")
	ck(-127, "TransformUpperSnakeCase(-127)")
	ck(127, "TransformUpperSnakeCase(127)")
}

func ck(value TransformUpperSnakeCase, str string) {
	if fmt.Sprint(value) != str {
		panic("transform_upper_snake_case.go: " + str)
	}
}
