package day_3

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetNumbersInLine(t *testing.T) {
	line := "...*...262....300..........+........507"
	actual := getNumbersInLine(line)

	expected := []Number{{Number: 262, IndexStart: 7, IndexEnd: 9}, {Number: 300, IndexStart: 14, IndexEnd: 16}, {Number: 507, IndexStart: 36, IndexEnd: 38}}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf(fmt.Sprintf("Expected %v. Received %v.", expected, actual))
	}
}

func TestGetSymbolsInLine(t *testing.T) {
	line := "...*...262....300..........+........507...."
	actual := getSymbolsInLine(line)

	expected := []Symbol{{Symbol: rune('*'), Index: 3}, {Symbol: rune('+'), Index: 27}}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf(fmt.Sprintf("Expected %v. Received %v.", expected, actual))
	}
}

func TestGetPartNumbersInLine(t *testing.T) {
	prevLine := stringToLine(".......262....300...................507.....961..............668.....................189.906...........................624..................")
	currLine := stringToLine("..148.................805..130..880*...........*684.............*......*..............*..-......%.................$........17...65....91*...")
	nextLine := stringToLine("......272.....464.....=......*.........................208*.....260.967.38.......692*.........676............@247..652.585.#......@......74.")

	expected := []int{805, 130, 880, 684, 17, 65, 91}
	actual := getPartNumbersInLine(currLine, prevLine, nextLine)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf(fmt.Sprintf("Expected %v. Received %v.", expected, actual))
	}
}
