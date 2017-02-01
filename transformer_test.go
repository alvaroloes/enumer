package main

import (
	"testing"
)

func TestTransform(t *testing.T) {
	result := transform("CamelCaseStringValue", ".")

	const expected = "camel.case.string.value"
	if result != expected {
		t.Errorf("\ngot: %s\n====\nexpected: %s", result, expected)
	}
}
