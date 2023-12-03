package day_2

import (
	"testing"
)

func TestParseLine(t *testing.T) {
	line := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	game, err := parseLine(line)

	if err != nil {
		t.Fatalf("Returned error for a valid line.")
	}

	if game.Id != 1 {
		t.Fatalf("Returned incorrect Id. Expected %v, Recieved %v", 1, game.Id)
	}

	if game.CubeSets[0].Blue != 3 {
		t.Fatalf("Returned incorrect Blue value for CubeSet 0")
	}

	if game.CubeSets[0].Red != 4 {
		t.Fatalf("Returned incorrect Red value for CubeSet 0")
	}

	if game.CubeSets[1].Red != 1 {
		t.Fatalf("Returned incorrect Red value for CubeSet 1")
	}

	if game.CubeSets[1].Green != 2 {
		t.Fatalf("Returned incorrect Green value for CubeSet 1")
	}

	if game.CubeSets[1].Blue != 6 {
		t.Fatalf("Returned incorrect Blue value for CubeSet 1")
	}

	if game.CubeSets[2].Green != 2 {
		t.Fatalf("Returned incorrect Green value for CubeSet 2")
	}
}
