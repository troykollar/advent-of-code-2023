package days

import (
	"advent-of-code-2023/read_input"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Day1() {
	input := read_input.ReadInput("./days/day_1/input.txt")

	splitInput := strings.Split(input, "\n")
	sum := 0
	for _, line := range splitInput {
		calibrationValue := getCalibrationValue(line)
		// fmt.Println(fmt.Sprintf("Line %v, Calibration Value: %v", i, calibrationValue))

		sum += calibrationValue
	}
	fmt.Printf("Sum: %v", sum)
}

func getCalibrationValue(line string) int {
	digits := getDigits(line)

	if len(digits) == 0 {
		return 0
	}

	firstDigit := digits[0]
	lastDigit := digits[0]

	for _, digit := range digits {
		if digit.Index < firstDigit.Index {
			firstDigit = digit
		}

		if digit.Index > lastDigit.Index {
			lastDigit = digit
		}
	}

	calibrationValue, _ := strconv.Atoi(firstDigit.Digit + lastDigit.Digit)
	return calibrationValue
}

func getDigits(line string) []Digit {
	digits := make([]Digit, 0)
	digits = append(digits, getNumericDigits(line)...)
	digits = append(digits, getTextDigits(line)...)
	return digits
}

func getNumericDigits(line string) []Digit {
	digits := make([]Digit, 0)
	for i, char := range line {
		if unicode.IsDigit(char) {
			digits = append(digits, Digit{Digit: string(char), Index: i})
		}
	}
	return digits
}

func getTextDigits(line string) []Digit {
	digitsAsText := []DigitAsText{
		{Text: "one", Digit: "1"},
		{Text: "two", Digit: "2"},
		{Text: "three", Digit: "3"},
		{Text: "four", Digit: "4"},
		{Text: "five", Digit: "5"},
		{Text: "six", Digit: "6"},
		{Text: "seven", Digit: "7"},
		{Text: "eight", Digit: "8"},
		{Text: "nine", Digit: "9"},
	}

	textDigits := make([]Digit, 0)
	for _, textDigit := range digitsAsText {
		firstIndex := strings.Index(line, textDigit.Text)
		if firstIndex != -1 {
			textDigits = append(textDigits, Digit{Digit: textDigit.Digit, Index: firstIndex})
		}

		lastIndex := strings.LastIndex(line, textDigit.Text)
		if lastIndex != -1 {
			textDigits = append(textDigits, Digit{Digit: textDigit.Digit, Index: lastIndex})
		}
	}

	return textDigits
}

type DigitAsText struct {
	Text  string
	Digit string
}

type Digit struct {
	Digit string
	Index int
}
