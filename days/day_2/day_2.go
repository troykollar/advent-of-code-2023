package day_2

import (
	"advent-of-code-2023/read_input"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

func Day2() {
	input := read_input.ReadInput("./days/day_2/input.txt")
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		fmt.Println(line)
		game, err := parseLine(line)
		if err != nil {
			continue
		}

		log := fmt.Sprintf("Game %v impossible.", game.Id)
		if isGamePossible(game) {
			log = fmt.Sprintf("Game %v possible.", game.Id)
			sum += game.Id
		}
		fmt.Println(log)
	}
	fmt.Println(fmt.Sprintf("Sum = %v", sum))
}

func parseLine(line string) (*Game, error) {
	game := new(Game)
	if !strings.Contains(line, "Game ") {
		return game, errors.New(fmt.Sprintf("Invalid line: '%v'", line))
	}

	game.Id = getId(line)
	game.CubeSets = getCubeSets(line)
	return game, nil
}

func getId(line string) int {
	splitLine := strings.Split(line, ":")
	gameIdString := splitLine[0]
	id, _ := strconv.Atoi(strings.Split(gameIdString, " ")[1])
	return id
}

func getCubeSets(line string) []CubeSet {
	cubeSetsString := strings.Split(line, ": ")[1]
	cubeSetStrings := strings.Split(cubeSetsString, "; ")
	cubeSets := make([]CubeSet, 0)

	for _, cubeSetString := range cubeSetStrings {
		cubeSets = append(cubeSets, *parseCubeSetString(cubeSetString))
	}

	return cubeSets
}

func parseCubeSetString(cubeSetString string) *CubeSet {
	cubeSet := new(CubeSet)

	valueStrings := strings.Split(cubeSetString, ", ")
	for _, valueString := range valueStrings {
		values := strings.Split(valueString, " ")
		value, _ := strconv.Atoi(values[0])
		color := values[1]

		if color == "red" {
			cubeSet.Red = value
		}

		if color == "green" {
			cubeSet.Green = value
		}

		if color == "blue" {
			cubeSet.Blue = value
		}
	}

	return cubeSet
}

func isGamePossible(game *Game) bool {
	for _, cubeSet := range game.CubeSets {
		if cubeSet.Red > MAX_RED {
			return false
		}
		if cubeSet.Green > MAX_GREEN {
			return false
		}
		if cubeSet.Blue > MAX_BLUE {
			return false
		}
	}

	return true
}

type Game struct {
	Id       int
	CubeSets []CubeSet
}

type CubeSet struct {
	Red   int
	Green int
	Blue  int
}
