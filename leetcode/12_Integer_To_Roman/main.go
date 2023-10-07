package main

var R = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

func intToRomanRec(num int, result string) string {
	if num < 1 {
		return result
	}

	if num >= 1000 {
		return intToRomanRec(num%1000, result+R[1000])
	}
	if num >= 900 {
		return intToRomanRec(num%900, result+R[900])
	}
	if num >= 500 {
		return intToRomanRec(num%500, result+R[500])
	}
	if num >= 400 {
		return intToRomanRec(num%400, result+R[400])
	}
	if num >= 100 {
		return intToRomanRec(num%100, result+R[100])
	}
	if num >= 90 {
		return intToRomanRec(num%90, result+R[90])
	}
	if num >= 50 {
		return intToRomanRec(num%50, result+R[50])
	}
	if num >= 40 {
		return intToRomanRec(num%40, result+R[40])
	}
	if num >= 10 {
		return intToRomanRec(num%10, result+R[10])
	}
	if num >= 9 {
		return intToRomanRec(num%9, result+R[9])
	}
	if num >= 5 {
		return intToRomanRec(num%5, result+R[5])
	}
	if num >= 4 {
		return intToRomanRec(num%4, result+R[4])
	}
	if num < 4 {
		return intToRomanRec(num-1, result+R[1])
	}

	return result
}

func intToRoman(num int) string {
	result := ""

	return intToRomanRec(num, result)
}
