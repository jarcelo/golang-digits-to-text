package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	var input = get_user_input()
	fmt.Println(convert_digits_to_words(input))
}

func get_user_input() int {
	fmt.Println("Enter a number to convert: ")
	scanner := bufio.NewScanner(os.Stdin)
	var text string
	scanner.Scan()
	text = scanner.Text()
	number, _  := strconv.Atoi(text)

	return number
}

func convert_digits_to_words(digit int) string {
	var one_million = 1000000
	var one_thousand = 1000
	var one_hundred = 100

	var millions = digit / one_million
	var millions_remainder = digit  % one_million

	var thousands = millions_remainder / one_thousand
	var thousands_remainder = millions_remainder % one_thousand

	var hundreds = thousands_remainder / one_hundred
	var hundreds_remainder = thousands_remainder % one_hundred

	var millions_part = translate_digit(millions)
	var thousands_part = translate_digit(thousands)
	var hundreds_part = translate_digit(hundreds)
	var remainder_part = translate_digit(hundreds_remainder)

	if (millions > 0) {
		millions_part = millions_part + " million "
	}
	if (thousands > 0) {
		thousands_part = thousands_part + " thousand "
	}
	if (hundreds > 0) {
		hundreds_part = hundreds_part + " hundred "
	} 

	var text_output =  millions_part + thousands_part + hundreds_part + remainder_part

	return text_output
}

func translate_digit(digit int) string {
    if (digit < 20) {
		return translate_twenty_below(digit)
	} else {
        var tens = digit / 10
        var tens_remainder = digit % 10
		return translate_divisible_by_ten(tens) + " " + translate_twenty_below(tens_remainder)
	}
}

func translate_twenty_below(digit int) string {
	var number = [20]string {
		" ", "one", "two", "three", "four", "five",
		"six", "seven", "eight", "nine", "ten", "eleven",
		"twelve", "thirteen", "fourteen", "fifteen", "sixteen",
		"seventeen", "eighteen", "nineteen",
	}
	return number[digit]
}

func translate_divisible_by_ten(digit int) string {
	var number = [10]string {" " , " ","twenty","thirty","forty",
		"fifty","sixty","seventy","eighty","ninety",
	}
	return number[digit]
}