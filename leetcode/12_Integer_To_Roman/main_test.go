package main

import (
	"fmt"
	"testing"
)

var cases = []struct {
	Exp   string
	Input int
}{
	{
		Exp:   "",
		Input: 0,
	},
	{
		Exp:   "I",
		Input: 1,
	},
	{
		Exp:   "II",
		Input: 2,
	},
	{
		Exp:   "III",
		Input: 3,
	},
	{
		Exp:   "IV",
		Input: 4,
	},
	{
		Exp:   "VII",
		Input: 7,
	},
	{
		Exp:   "IX",
		Input: 9,
	},
	{
		Exp:   "XI",
		Input: 11,
	},
	{
		Exp:   "XIII",
		Input: 13,
	},
	{
		Exp:   "XX",
		Input: 20,
	},
	{
		Exp:   "LVIII",
		Input: 58,
	},
	{
		Exp:   "C",
		Input: 100,
	},
	{
		Exp:   "CD",
		Input: 400,
	},
	{
		Exp:   "D",
		Input: 500,
	},
	{
		Exp:   "M",
		Input: 1000,
	},
	{
		Exp:   "MI",
		Input: 1001,
	},
	{
		Exp:   "MCMXCIV",
		Input: 1994,
	},
}

func TestIntToRoman(t *testing.T) {
	for _, e := range cases {
		t.Run(fmt.Sprintf("input=%d exp=%s", e.Input, e.Exp), func(t *testing.T) {
			exp := e.Exp
			got := intToRoman(e.Input)
			if exp != got {
				t.Fatalf("exp %s; got %s", exp, got)
			}
		})
	}
}

func TestIntToRomanRec(t *testing.T) {
	for _, e := range cases {
		t.Run(fmt.Sprintf("input=%d exp=%s", e.Input, e.Exp), func(t *testing.T) {
			exp := e.Exp
			got := intToRomanRec(e.Input, "")
			if exp != got {
				t.Fatalf("exp %s; got %s", exp, got)
			}
		})
	}
}
