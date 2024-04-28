package main

import "fmt"

type validationCode int8

const (
	valid validationCode = iota
	arabicAndRomanNumerals
	incorrectSymbol
	inputIsEmpty
)

func (code validationCode) String() string {
	switch code {
	case valid:
		return ""
	case arabicAndRomanNumerals:
		return "Input contains both arabic and roman numerals."
	case incorrectSymbol:
		return "Input contains incorrect symbol."
	case inputIsEmpty:
		return "Input is empty."
	default:
		return ""
	}
}

func ensureValidInput(input string) validationCode {
	if len(input) == 0 {
		return inputIsEmpty
	}
	hasArabicNumerals := false
	hasRomanNumerals := false
	hasIncorrectSymbol := false
	for _, char := range input {
		if char >= '0' && char <= '9' { // there is arabic numerals
			hasArabicNumerals = true
		} else if char == 'I' || char == 'V' || char == 'X' { // there is roman numerals
			hasRomanNumerals = true
		} else if !(char == ' ' || char == '*' || char == '/' || char == '+' || char == '-') {
			hasIncorrectSymbol = true
		}
	}
	if hasIncorrectSymbol {
		return incorrectSymbol
	}
	if hasArabicNumerals && hasRomanNumerals {
		return arabicAndRomanNumerals
	}
	return valid
}

func arabicToRoman(arabic int) string {
	return ""
}

func romanToArabic(roman string) int {
	return 0
}

func main() {
	exitMsg := "Use CTRL + C to exit the program or type another expression."
	var input string
	for {
		fmt.Scanln(&input)
		validationCode := ensureValidInput(input)
		if validationCode != valid {
			fmt.Println("Experssion is invalid : " + validationCode.String() + exitMsg)
			continue
		}

	}
}
