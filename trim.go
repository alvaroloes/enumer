package main

import "strings"

func longestCommonPrefix(values []Value) string {
	// LCP of min and max (lexigraphically)
	// is the LCP of the whole set.
	min := values[0].name
	max := min
	for _, s := range values[1:] {
		switch {
		case s.name < min:
			min = s.name
		case s.name > max:
			max = s.name
		}
	}
	for i := 0; i < len(min) && i < len(max); i++ {
		if min[i] != max[i] {
			return min[:i]
		}
	}
	// If all the bytes are equal but the lengths aren't then
	// min is a prefix of max, and hence the lcp
	return min
}

func autoPrefix(values []Value) string {
	// Can't trim a single value
	if len(values) < 2 {
		return ""
	}
	prefix := longestCommonPrefix(values)

	return prefix
}

func (g *Generator) trimValueNames(values []Value, prefix string) {
	for i := range values {
		values[i].name = strings.TrimPrefix(values[i].name, prefix)
	}
}

func (g *Generator) autoTrimValueNames(values []Value) {
	prefix := autoPrefix(values)
	g.trimValueNames(values, prefix)
}
