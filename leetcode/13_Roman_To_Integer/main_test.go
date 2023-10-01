package main

import (
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	cases := []struct {
		Input string
		Exp   int
	}{
		{
			Input: "III",
			Exp:   3,
		},
		{
			Input: "LVIII",
			Exp:   58,
		},
		{
			Input: "MCMXCIV",
			Exp:   1994,
		},
	}

	for _, e := range cases {
		t.Run(fmt.Sprintf("input %s; exp; %d", e.Input, e.Exp), func(t *testing.T) {
			exp := e.Exp
			got := romanToInt(e.Input)

			if got != exp {
				t.Errorf("exp %d; got %d", exp, got)
			}
		})
	}
}
