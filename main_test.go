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
			expected: []string{"--", "12test.", "aaa_$!"},
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

func TestGetCommands(t *testing.T) {
	commands := getCommands()
	if len(commands) == 0 {
		t.Errorf("No items in the returned map")
	}

	for _, command := range commands {
		if command.name == "" {
			t.Errorf("Empty command name")
		}

		if command.description == "" {
			t.Errorf("Empty command description")
		}

		if command.callback == nil {
			t.Errorf("Empty command callback function")
		}
	}
}
