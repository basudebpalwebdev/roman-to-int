package main

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	listOfTestCasesWithResults := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "Vi": 6, "iX": 9, "Xi": 11, "XX": 20, "XL": 40, "LXI": 61,
		"MMXIX": 2019,
	}
	for k, v := range listOfTestCasesWithResults {
		got := RomanToInt(k)
		want := v

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}
}
