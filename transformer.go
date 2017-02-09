package main

import (
	"strings"

	"github.com/fatih/camelcase"
)

func transform(src, delim string) string {
	entries := camelcase.Split(src)
	if len(entries) <= 1 {
		return strings.ToLower(src)
	}

	result := strings.ToLower(entries[0])
	for i := 1; i < len(entries); i++ {
		result += delim + strings.ToLower(entries[i])
	}
	return result
}

func toSnakeCase(src string) string {
	return transform(src, "_")
}

func toKebabCase(src string) string {
	return transform(src, "-")
}
