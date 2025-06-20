package main

import ( 
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
	
	{	input:    " This is a Test  ",
		expected: []string{"this", "is","a","test"},
	},

	{	input:    " This is ONLY a Test  ",
		expected: []string{"this", "is","only", "a","test"},
	},
}

for _, c := range cases {
	actual := cleanInput(c.input)
	if len(actual) != len(c.expected) {
		t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
		continue
	}
	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
		if word != expectedWord {
			t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
		}
	}
}
}

