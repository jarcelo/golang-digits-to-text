package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func main() {
	var input string
	for input != "q" {
		input = GetUserInput()
		var number = ValidateUserInput(input)
		fmt.Println(ConvertDigitsToWords(number))
	}
}

func ValidateUserInput(input string) int {
	number, err  := strconv.Atoi(input)
	if err != nil && input != "q" {
		fmt.Println ("Invalid input. Please try again.")
	}
	return number
}

func GetUserInput() string {
	fmt.Println("Enter a number to convert. Hit 'q' to exit.")
	scanner := bufio.NewScanner(os.Stdin)
	var text string
	scanner.Scan()
	text = scanner.Text()
	text = strings.TrimSpace(text)
	return text
}

func ConvertDigitsToWords(digit int) string {
	var one_million = 1000000
	var one_thousand = 1000
	var one_hundred = 100

	var millions = digit / one_million
	var millions_remainder = digit  % one_million

	var thousands = millions_remainder / one_thousand
	var thousands_remainder = millions_remainder % one_thousand

	var hundreds = thousands_remainder / one_hundred
	var hundreds_remainder = thousands_remainder % one_hundred

	var millions_part = TranslateDigit(millions)
	var thousands_part = TranslateDigit(thousands)
	var hundreds_part = TranslateDigit(hundreds)
	var remainder_part = TranslateDigit(hundreds_remainder)

	if (millions > 0) {
		millions_part = millions_part + " million "
	}
	if (thousands > 0) {
		thousands_part = thousands_part + " thousand "
	}
	if (hundreds > 0) {
		hundreds_part = hundreds_part + " hundred "
	} 
	return strings.TrimSpace(millions_part + thousands_part + hundreds_part + remainder_part)
}

func TranslateDigit(digit int) string {
    if (digit < 20) {
		return TranslateTwentyBelow(digit)
	} else {
        var tens = digit / 10
        var tens_remainder = digit % 10
		return TranslateDivisibleByTen(tens) + " " + TranslateTwentyBelow(tens_remainder)
	}
}

func TranslateTwentyBelow(digit int) string {
	var number = [20]string {
		"", "one", "two", "three", "four", "five",
		"six", "seven", "eight", "nine", "ten", "eleven",
		"twelve", "thirteen", "fourteen", "fifteen", "sixteen",
		"seventeen", "eighteen", "nineteen",
	}
	return number[digit]
}

func TranslateDivisibleByTen(digit int) string {
	var number = [10]string {"" , "","twenty","thirty","forty",
		"fifty","sixty","seventy","eighty","ninety",
	}
	return number[digit]
}