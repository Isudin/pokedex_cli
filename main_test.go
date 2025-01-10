package main

import (
	"strings"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "-- 12Test. aaa_$!",
			expected: []string{"--", "12Test.", "aaa_$!"},
		},
	}

	for _, c := range cases {
		output := cleanInput(c.input)
		allOutput := strings.Join(output, ", ")
		allExpected := strings.Join(c.expected, ", ")
		if len(output) != len(c.expected) {
			t.Errorf("Error matching output\nExpected: [%v]\nOutput: [%v]", allExpected, allOutput)
		}
		for i := range output {
			if output[i] != c.expected[i] {
				t.Errorf("Error matching output\nExpected: [%v]\nOutput: [%v]", allExpected, allOutput)
			}
		}
	}
}
