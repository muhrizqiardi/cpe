package main

import "fmt"

var RomanNumeralsMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	result := 0

loopString:
	for i := 0; i < len(s); i++ {
		fmt.Printf("i=%d; s[i:i+2]=%s;\n", i, s[i:i+2])
		switch s[i : i+2] {
		case "IV":
			result += 4
			i++
			continue loopString
		case "IX":
			result += 9
			i++
			continue loopString
		case "XL":
			result += 40
			i++
			continue loopString
		case "XC":
			result += 90
			i++
			continue loopString
		case "CD":
			result += 400
			i++
			continue loopString
		case "CM":
			result += 900
			i++
			continue loopString
		default:
			result += RomanNumeralsMap[string(s[i])]
		}
	}

	return result
}
