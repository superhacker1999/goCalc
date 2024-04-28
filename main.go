package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type validationCode int8
type operationCode int8

const (
	arabic validationCode = iota
	roman
	invalid
	arabicAndRomanNumerals
	incorrectSymbol
	inputIsEmpty
)

const (
	multiply operationCode = iota
	divide
	add
	subtract
)

func (code validationCode) String() string {
	switch code {
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
		if char >= '0' && char <= '9' {
			hasArabicNumerals = true
		} else if char == 'I' || char == 'V' || char == 'X' {
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
	// the only valid variant
	if hasArabicNumerals {
		return arabic
	} else {
		return roman
	}
}

func arabicToRoman(num int) string {
	if num < 0 {
		return "Error: Result can not be negative using roman numbers!"
	}

	// Define the mapping of Arabic to Roman numerals
	romanNumerals := []struct {
		arabic int
		roman  string
	}{
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

	// Convert Arabic to Roman numerals
	var result string
	for _, rn := range romanNumerals {
		for num >= rn.arabic {
			result += rn.roman
			num -= rn.arabic
		}
	}

	return result
}

func romanToArabic(roman string) int {
	// Define the mapping of Roman to Arabic numerals
	romanNumerals := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
	}

	var result int
	prevValue := 0
	for i := len(roman) - 1; i >= 0; i-- {
		value := romanNumerals[string(roman[i])]
		if value < prevValue {
			result -= value
		} else {
			result += value
		}
		prevValue = value
	}

	return result
}

func getNumbers(code validationCode, input string) (int, int) {
	var num1 int
	var num2 int

	if code == arabic {
		num1, _ = strconv.Atoi(input[0:strings.IndexByte(input, ' ')])
		num2, _ = strconv.Atoi(input[strings.LastIndexByte(input, ' ')+1:])
	} else {
		num1 = romanToArabic(input[0:strings.IndexByte(input, ' ')])
		num2 = romanToArabic(input[strings.LastIndexByte(input, ' ')+1:])
	}
	return num1, num2
}

func getOperation(input string) operationCode {
	for _, char := range input {
		if char == '*' {
			return multiply
		} else if char == '/' {
			return divide
		} else if char == '+' {
			return add
		} else if char == '-' {
			return subtract
		}
	}
	return -1
}

func applyOperation(operation operationCode, num1, num2 int) int {
	switch operation {
	case multiply:
		return num1 * num2
	case divide:
		return num1 / num2
	case add:
		return num1 + num2
	case subtract:
		return num1 - num2
	default:
		return 0
	}
}

func main() {
	exitMsg := "\nUse CTRL + C to exit the program or type another expression."
	fmt.Println("This is a simple calculator that can perform addition, subtraction, multiplication, and division.")
	fmt.Println("Enter an expression in the following format: <number> <operation> <number>")
	fmt.Println("For example: 1 + 2 or I + II")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		input := scanner.Text()
		validationCode := ensureValidInput(input)
		if validationCode > invalid {
			fmt.Println("Experssion is invalid : " + validationCode.String() + exitMsg)
			continue
		}
		num1, num2 := getNumbers(validationCode, input)
		operation := getOperation(input)
		result := applyOperation(operation, num1, num2)
		fmt.Print("Result is : ")
		if validationCode == arabic {
			fmt.Println(result, exitMsg)
		} else {
			fmt.Println(arabicToRoman(result) + exitMsg)
		}
	}
}
