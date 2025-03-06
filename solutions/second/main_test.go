package main

import "testing"

type testCase struct {
	name     string
	input    string
	expected string
}

func TestZaiKwa(t *testing.T) {
	tests := []testCase{
		{
			name:     "case 1",
			input:    "LLRR=",
			expected: "210122",
		},
		{
			name:     "case 2",
			input:    "==RLL",
			expected: "000210",
		},
		{
			name:     "case 3",
			input:    "=LLRR",
			expected: "221012",
		},
		{
			name:     "case 4",
			input:    "RRL=R",
			expected: "012001",
		},
	}

	for _, test := range tests {
		t.Run("Test ZaiKwa", func(t *testing.T) {
			result := zaiKwa(test.input)
			if result != test.expected {
				t.Errorf(`
					==========
					| FAILED |
					==========
					case:     %s
					expected: <%s>
					but got:  <%s>
			`, test.name, test.expected, result)
			}
		})
	}
}
