package main

import "testing"

func TestConvertDigitsToWords(t *testing.T)  {
	expected := "two"
	actual := ConvertDigitsToWords(2)
	if actual != expected {
		t.Errorf("Test failed.")
	}
}