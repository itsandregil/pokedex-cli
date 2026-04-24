package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		result := cleanInput(c.input)
		if len(result) != len(c.expected) {
			t.Errorf("length mismatch: got %d, want %d", len(result), len(c.expected))
		}
		for i := range result {
			word := result[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("words do not match: got %s - want %s", word, expectedWord)
			}
		}
	}
}
