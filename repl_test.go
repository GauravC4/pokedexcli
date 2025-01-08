package main

import "testing"

func TestCleanInput(t *testing.T) {
	type TestCase struct {
		input    string
		expected []string
	}

	cases := []TestCase{
		{
			input:    " charmander BulbaSaur ",
			expected: []string{"charmander", "bulbasaur"},
		},
		{
			input:    " squirtle    mewoth  kingdra SLOWBROW  slaKING",
			expected: []string{"squirtle", "mewoth", "kingdra", "slowbrow", "slaking"},
		},
		{
			input:    " PikAchu",
			expected: []string{"pikachu"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "charmander charmeleon charizard",
			expected: []string{"charmander", "charmeleon", "charizard"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("slice length mismatch for input %v\n", c.input)
			continue
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("wanted %v got %v\n", c.expected[i], actual[i])
			}
		}
	}
}
