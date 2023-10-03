package main

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
		var sRange string
		if i == len(s)-1 {
			sRange = string(s[i])
		} else {
			sRange = s[i : i+2]
		}

		switch sRange {
		case "IV":
			result += 4
			i += 1
			continue loopString
		case "IX":
			result += 9
			i += 1
			continue loopString
		case "XL":
			result += 40
			i += 1
			continue loopString
		case "XC":
			result += 90
			i += 1
			continue loopString
		case "CD":
			result += 400
			i += 1
			continue loopString
		case "CM":
			result += 900
			i += 1
			continue loopString
		default:
			result += RomanNumeralsMap[string(s[i])]
		}
	}

	return result
}
