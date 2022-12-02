package hand

import (
	"testing"
)

// Testing to make sure the normalization of the hand symbols work as expected
func TestNormalizedHand(t *testing.T) {
	inputHands := []string{"X", "Y", "Z", "A", "B", "C"}
	expectedOutputs := []string{"ROCK", "PAPER", "SCISSORS", "ROCK", "PAPER", "SCISSORS"}
	for index, inputHand := range inputHands {
		normalizedHand, err := NewHand(inputHand)
		if err != nil {
			t.Fatal(err)
		}
		if expectedOutputs[index] != normalizedHand.shapeThrown {
			t.Fatalf("Normalized hand failed with input %s and expected output %s", inputHand, expectedOutputs[index])
		}
	}
}

func TestThrowErrorWhenHandInputIsNotUnderstood(t *testing.T) {
	input := []string{"H"}
	_, err := NewHand(input[0])
	if err == nil {
		t.Fatalf("Excepted to get an error when creating new hand with unknown input")
	}
}

func TestShapeScoring(t *testing.T) {
	hands := []Hand{
		Hand{shapeThrown: "ROCK"},
		Hand{shapeThrown: "PAPER"},
		Hand{shapeThrown: "SCISSORS"},
	}
	expectedOutputs := []int{1, 2, 3}
	for index, hand := range hands {
		actual := hand.shapeScore()
		expected := expectedOutputs[index]
		if expected != actual {
			t.Fatalf("Hand value %d is incorrect for normalized shape %v, should be %d", actual, hand, expected)
		}
	}
}

func TestScoreHand(t *testing.T) {
	// Hand matrix [0] = playerHand, [1] = enemyHand
	rounds := [][]Hand{
		// Draws
		{Hand{shapeThrown: "ROCK"}, Hand{shapeThrown: "ROCK"}},
		{Hand{shapeThrown: "PAPER"}, Hand{shapeThrown: "PAPER"}},
		{Hand{shapeThrown: "SCISSORS"}, Hand{shapeThrown: "SCISSORS"}},
		// Player wins with rock
		{Hand{shapeThrown: "ROCK"}, Hand{shapeThrown: "SCISSORS"}},
		// Player wins with PAPER
		{Hand{shapeThrown: "PAPER"}, Hand{shapeThrown: "ROCK"}},
		// Player wins with SCISSORS
		{Hand{shapeThrown: "SCISSORS"}, Hand{shapeThrown: "PAPER"}},
		// Player loses with rock
		{Hand{shapeThrown: "ROCK"}, Hand{shapeThrown: "PAPER"}},
		// Player loses with PAPER
		{Hand{shapeThrown: "PAPER"}, Hand{shapeThrown: "SCISSORS"}},
		// Player loses with SCISSORS
		{Hand{shapeThrown: "SCISSORS"}, Hand{shapeThrown: "ROCK"}},
	}

	expectedValues := []int{4, 5, 6, 7, 8, 9, 1, 2, 3}
	for index, round := range rounds {
		actual := ScoreHands(round[1], round[0])
		if actual != expectedValues[index] {
			t.Fatalf("Round value is incorrect, Player threw: %v, enemy threw %v. actual score %d, expected score %d", round[0], round[1], actual, expectedValues[index])
		}
	}

}
