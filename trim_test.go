package main

import "testing"

var lcpTests = []struct {
	expected string
	in       []string
}{
	{"Proto", []string{"ProtoOne", "ProtoTwo", "ProtoThree"}},
	// An empty string is OK when one value is the prefix
	{"Proto", []string{"Proto", "ProtoLonger"}},
	{"", []string{}},
	{"", []string{"aardvark"}},
	{"", []string{"abc", "def", "deg"}},
	{"ab", []string{"ab", "abc", "abcd"}},
}

// TestLcp checks that the longest common prefix is generated correctly
func TestLcp(t *testing.T) {
	for _, tt := range lcpTests {
		values := make([]Value, len(tt.in))
		for i := range tt.in {
			values[i] = Value{
				name: tt.in[i],
			}
		}
		prefix := autoPrefix(values)
		if prefix != tt.expected {
			t.Errorf("%q => %s, expected %s", tt.in, tt.expected, prefix)
		}
	}
}
