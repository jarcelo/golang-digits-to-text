package main

import "testing"

func TestConvertDigitsToWords(t *testing.T)  {
	var tests = []struct {
		in int
		out string
	}{
		{2, "two"},
		{10, "ten"},
		{100, "one hundred"},
		{999, "nine hundred ninety nine"},
		{1010, "one thousand ten"},
		{9999, "nine thousand nine hundred ninety nine"},
		{99999, "ninety nine thousand nine hundred ninety nine"},
		{999999, "nine hundred ninety nine thousand nine hundred ninety nine"},  //TODO: This test fails
		{1000000, "one million"},
	}

	for _, test := range tests {
		actual := ConvertDigitsToWords(test.in)
		if actual != test.out {
			t.Errorf("Test failed. Expected: %v; Actual: %s", test.out, actual)
		}
	}
}