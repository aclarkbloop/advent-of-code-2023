package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	p1()
	p2()
}

// #1
func p1() {
	file, err := os.Open("../inputs/day4.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		card := scanner.Text()
		cardComponents := strings.Split(card, ": ")
		cardNumbers := strings.Split(cardComponents[1], "|")
		winningNumbers := strings.Split(strings.TrimSpace(cardNumbers[0]), " ")
		yourNumbers := strings.Split(strings.TrimSpace(cardNumbers[1]), " ")
		cardScore := 0
		foundNums := []string{}
		for i := 0; i < len(yourNumbers); i++ {
			if yourNumbers[i] == "" {
				continue
			}
			for j := 0; j < len(winningNumbers); j++ {
				if yourNumbers[i] == winningNumbers[j] && !slices.Contains(foundNums, yourNumbers[i]) {
					if cardScore == 0 {
						cardScore = 1
					} else {
						cardScore += cardScore
					}
					foundNums = append(foundNums, yourNumbers[i])
					break
				}
			}
		}
		sum += cardScore
	}
	fmt.Println("The answer to problem #1 is:", sum)
}

// #2
func p2() {
	file, err := os.Open("../inputs/day4.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// sum := 0
	cards := []string{}
	for scanner.Scan() {
		cards = append(cards, scanner.Text())
	}
	var totalCards func(newCards []string, total int, index int) int
	totalCards = func(newCards []string, total int, index int) int {
		if len(newCards) == 0 {
			if index == len(cards)-1 {
				// we have found all subsequent cards for all cards
				return total
			}
			// found all subsequent cards for this card, move to next card
			return totalCards([]string{cards[index+1]}, total, index+1)
		}
		numCardsWon := 0
		newCardInputs := []string{}
		for j := 0; j < len(newCards); j++ {
			numCardsWon = processCard(newCards[j])
			total += numCardsWon
			newCardInputs = append(newCardInputs, getWonCards(cards, newCards[j], numCardsWon)...)
		}
		return totalCards(newCardInputs, total, index)
	}
	fmt.Println("The answer to problem #2 is:", totalCards([]string{cards[0]}, len(cards), 0))
}

func processCard(card string) int {
	cardComponents := strings.Split(card, ": ")
	cardNumbers := strings.Split(cardComponents[1], "|")
	winningNumbers := strings.Split(strings.TrimSpace(cardNumbers[0]), " ")
	yourNumbers := strings.Split(strings.TrimSpace(cardNumbers[1]), " ")
	cardScore := 0
	foundNums := []string{}
	for i := 0; i < len(yourNumbers); i++ {
		if yourNumbers[i] == "" {
			continue
		}
		for j := 0; j < len(winningNumbers); j++ {
			if yourNumbers[i] == winningNumbers[j] && !slices.Contains(foundNums, yourNumbers[i]) {
				cardScore += 1
				foundNums = append(foundNums, yourNumbers[i])
				break
			}
		}
	}
	return cardScore
}

func getWonCards(cards []string, currentCard string, cardsWon int) []string {
	cardComponents := strings.Split(currentCard, ": ")
	currGame := strings.Split(cardComponents[0], " ")
	currGameNum, e := strconv.Atoi(currGame[len(currGame)-1])
	if e != nil {
		log.Fatalf("unable to convert current game to int: %v", e)
	}
	currIndex := currGameNum
	maxIndex := currIndex + cardsWon
	if difference := (len(cards) - 1) - (maxIndex); difference < 0 {
		maxIndex = maxIndex + difference
	}
	newCards := []string{}
	for i := currIndex; i < maxIndex; i++ {
		newCards = append(newCards, cards[i])
	}
	return newCards
}
