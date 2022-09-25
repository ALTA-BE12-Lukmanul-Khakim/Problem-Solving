package main

import "fmt"

func RomanNumerals(number int) string {
	konv := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""
	for _, konv := range konv {
		for number >= konv.value {
			roman += konv.digit
			number -= konv.value
		}
	}
	return roman
}

func main() {
	fmt.Println(RomanNumerals(6))
	fmt.Println(RomanNumerals(9))
	fmt.Println(RomanNumerals(23))
	fmt.Println(RomanNumerals(2021))
	fmt.Println(RomanNumerals(1646))
}
