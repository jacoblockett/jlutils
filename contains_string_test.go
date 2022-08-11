package jlutils

import "testing"

func TestTableContainsString(t *testing.T) {
	var tests = []struct {
		arga     []string
		argb     string
		expected bool
	}{
		{[]string{"a", "b", "c"}, "a", true},
		{[]string{"a", "b", "c"}, "b", true},
		{[]string{"a", "b", "c"}, "c", true},
		{[]string{"a", "b", "c"}, "d", false},
	}

	for _, test := range tests {
		if output := ContainsString(test.arga, test.argb); output != test.expected {
			t.Errorf("Test failed:\n1st arg:  %s\n2nd arg:  %s\nexpected: %t\nreceived: %t", test.arga, test.argb, test.expected, output)
		}
	}
}
