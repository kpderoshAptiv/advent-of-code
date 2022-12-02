package main

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/kpderoshAptiv/advent-of-code/dec_02_2022/hand"
)

func main() {

	content, err := ioutil.ReadFile("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	fileContentsString := string(content)
	rounds := strings.Split(fileContentsString, "\n")
	var finalScore int = 0
	for roundIndex, roundString := range rounds {
		hands := strings.Split(roundString, " ")
		enemyHand, err := hand.NewHand(hands[0])
		if err != nil {
			panic(err)
		}
		playerHand, err := hand.NewHand(hands[1])
		if err != nil {
			panic(err)
		}
		log.Printf("Round %d, FIGHT\n", roundIndex)
		log.Printf("\tEnemy uses: %v\n", enemyHand)
		log.Printf("\tYou use: %v\n", playerHand)
		roundScore := hand.ScoreHands(enemyHand, playerHand)
		log.Printf("Your score: %d", roundScore)
		finalScore += roundScore
	}
	log.Printf("Final score: %d", finalScore)
}
