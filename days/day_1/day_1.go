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
	for i, line := range splitInput {
		calibrationValue := getCalibrationValue(line)
		fmt.Println(fmt.Sprintf("Line %v calibration value: %v", i, calibrationValue))

		sum += getCalibrationValue(line)
	}
	fmt.Printf("Sum: %v", sum)
}

func getCalibrationValue(line string) int {
	digits := make([]rune, 0)

	for _, character := range line {
		if unicode.IsDigit(character) {
			digits = append(digits, character)
		}
	}

	if len(digits) == 0 {
		return 0
	}

	if len(digits) == 1 {
		digits = append(digits, digits[0])
	}

	firstDigit := digits[0]
	lastDigit := digits[len(digits)-1]
	calibrationValueString := string([]rune{firstDigit, lastDigit})

	calibrationValue, err := strconv.Atoi(calibrationValueString)

	if err != nil {
		panic(fmt.Sprintf("Non-numeric digit entered. Error: %v", err))
	}

	return calibrationValue
}
