package hand

import "fmt"

const SCORE_FOR_ROCK int = 1
const SCORE_FOR_PAPER int = 2
const SCORE_FOR_SCISSORS int = 3
const ROCK string = "ROCK"
const PAPER string = "PAPER"
const SCISSORS string = "SCISSORS"

func NewHand(hand string) (Hand, error) {
	if hand == "X" || hand == "A" {
		return Hand{ROCK}, nil
	} else if hand == "Y" || hand == "B" {
		return Hand{PAPER}, nil
	} else if hand == "Z" || hand == "C" {
		return Hand{SCISSORS}, nil
	}
	return Hand{}, fmt.Errorf("unable to normalize hand")
}

type Hand struct {
	shapeThrown string
}

func (hand Hand) shapeScore() int {
	if hand.shapeThrown == ROCK {
		return SCORE_FOR_ROCK
	} else if hand.shapeThrown == PAPER {
		return SCORE_FOR_PAPER
	} else if hand.shapeThrown == SCISSORS {
		return SCORE_FOR_SCISSORS
	}
	return 0
}

func ScoreHands(enemyHand Hand, playerHand Hand) int {
	winScore := 6
	drawScore := 3
	playerShapeScore := playerHand.shapeScore()
	win := winScore + playerShapeScore
	// Draw
	if enemyHand.shapeThrown == playerHand.shapeThrown {
		return drawScore + playerShapeScore
	} else if playerHand.shapeThrown == ROCK && enemyHand.shapeThrown == SCISSORS {
		// ROCK beats SCISSORS
		return win
	} else if playerHand.shapeThrown == PAPER && enemyHand.shapeThrown == ROCK {
		// PAPER beats ROCK
		return win
	} else if playerHand.shapeThrown == SCISSORS && enemyHand.shapeThrown == PAPER {
		return win
	}
	return 0 + playerShapeScore
}
