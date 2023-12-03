package day_3

import (
	"advent-of-code-2023/read_input"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Day3() {
	input := read_input.ReadInput("./days/day_3/input.txt")
	lineStrings := strings.Split(input, "\n")

	lines := make([]Line, 0)
	for _, lineString := range lineStrings {
		lines = append(lines, stringToLine(lineString))
	}

	sum := 0
	gearRatioSum := 0

	for i, line := range lines {
		prevLine := Line{Numbers: make([]Number, 0), Symbols: make([]Symbol, 0)}
		nextLine := Line{Numbers: make([]Number, 0), Symbols: make([]Symbol, 0)}
		if i != 0 {
			prevLine = lines[i-1]
		}
		if i != len(lines)-1 {
			nextLine = lines[i+1]
		}

		partNumbers := getPartNumbersInLine(line, prevLine, nextLine)
		lineSum := 0
		for _, partNumber := range partNumbers {
			lineSum += partNumber
		}
		sum += lineSum

		for _, symbol := range line.Symbols {
			gearRatioSum += getGearRatio(symbol, line, prevLine, nextLine)
		}
	}

	fmt.Println(fmt.Sprintf("Sum: %v, Gear Ratio Sum: %v", sum, gearRatioSum))
}

func getGearRatio(symbol Symbol, line Line, prevLine Line, nextLine Line) int {
	if symbol.Symbol != '*' {
		return 0
	}

	adjacentNumbers := make([]int, 0)

	for _, number := range line.Numbers {
		if symbol.Index >= number.IndexStart-1 && symbol.Index <= number.IndexEnd+1 {
			adjacentNumbers = append(adjacentNumbers, number.Number)
		}
	}

	for _, number := range prevLine.Numbers {
		if symbol.Index >= number.IndexStart-1 && symbol.Index <= number.IndexEnd+1 {
			adjacentNumbers = append(adjacentNumbers, number.Number)
		}
	}

	for _, number := range nextLine.Numbers {
		if symbol.Index >= number.IndexStart-1 && symbol.Index <= number.IndexEnd+1 {
			adjacentNumbers = append(adjacentNumbers, number.Number)
		}
	}

	if len(adjacentNumbers) != 2 {
		return 0
	}

	return adjacentNumbers[0] * adjacentNumbers[1]
}

func getPartNumbersInLine(line Line, prevLine Line, nextLine Line) []int {
	partNumbers := make([]int, 0)
	for _, number := range line.Numbers {
		if lineContainsAdjacentSymbol(number, line) || lineContainsAdjacentSymbol(number, prevLine) || lineContainsAdjacentSymbol(number, nextLine) {
			partNumbers = append(partNumbers, number.Number)
		}
	}
	return partNumbers
}

func lineContainsAdjacentSymbol(number Number, line Line) bool {
	for _, symbol := range line.Symbols {
		if symbol.Index >= number.IndexStart-1 && symbol.Index <= number.IndexEnd+1 {
			return true
		}
	}
	return false
}

func stringToLine(lineString string) Line {
	return Line{Numbers: getNumbersInLine(lineString), Symbols: getSymbolsInLine(lineString)}
}

func getNumbersInLine(line string) []Number {
	numbers := make([]Number, 0)

	number := ""
	indexStart := 0
	indexEnd := 0
	for index, char := range line {
		if unicode.IsDigit(char) {
			if len(number) == 0 {
				indexStart = index
			}

			number += string(char)
		} else if len(number) != 0 {
			indexEnd = index - 1

			num, err := strconv.Atoi(number)
			if err == nil {
				numbers = append(numbers, Number{Number: num, IndexStart: indexStart, IndexEnd: indexEnd})
				number = ""
				indexStart = 0
				indexEnd = 0
			}
		}
	}

	if len(number) != 0 {
		num, err := strconv.Atoi(number)
		if err == nil {
			numbers = append(numbers, Number{Number: num, IndexStart: indexStart, IndexEnd: len(line) - 1})
			number = ""
			indexStart = 0
			indexEnd = 0
		}
	}

	return numbers
}

func getSymbolsInLine(line string) []Symbol {
	symbols := make([]Symbol, 0)

	for index, char := range line {
		if unicode.IsDigit(char) || char == '.' {
			continue
		}

		symbols = append(symbols, Symbol{Symbol: char, Index: index})
	}

	return symbols
}

type Symbol struct {
	Symbol rune
	Index  int
}

type Number struct {
	Number     int
	IndexStart int
	IndexEnd   int
}

type Line struct {
	Numbers []Number
	Symbols []Symbol
}
