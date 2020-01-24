package numerals

import "strings"

// RomanNumeral is mapping between arabic and roman
type RomanNumeral struct {
	Value  uint16
	Symbol string
}

// RomanNumerals is a list of mappings
type RomanNumerals []RomanNumeral

// ValueOf looks up for arabic value of roman symbol in the mapping
func (r RomanNumerals) ValueOf(symbols ...byte) uint16 {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}

var romanNumerals = RomanNumerals{
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

// ConvertToRoman covnerts a roman string to arabic integer
func ConvertToRoman(arabic uint16) string {

	var result strings.Builder

	for _, numeral := range romanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return result.String()
}

// ConvertToArabic covnerts arabic integer to a roman string
func ConvertToArabic(roman string) uint16 {
	var total uint16 = 0

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		if i+1 < len(roman) && (symbol == 'I' || symbol == 'X' || symbol == 'C') {
			if value := romanNumerals.ValueOf(symbol, roman[i+1]); value != 0 {
				total += value
				i++
			} else {
				total += romanNumerals.ValueOf(symbol)
			}
		} else {
			total += romanNumerals.ValueOf(symbol)
		}
	}
	return total
}
