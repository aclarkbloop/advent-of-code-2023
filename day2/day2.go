package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	p1()
	p2()
}

// #1
func p1() {
	file, err := os.Open("../inputs/day2.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	colors := map[string]int{"red": 12, "green": 13, "blue": 14}
	for scanner.Scan() {
		gameLog := scanner.Text()
		logComponents := strings.Split(gameLog, ": ")
		gameInfo := strings.Split(logComponents[0], " ")
		gameNum, e := strconv.Atoi(gameInfo[1])
		if e != nil {
			log.Fatalf("unable to convert game number to int: %v", e)
		}
		rounds := strings.Split(logComponents[1], "; ")
		possible := true
		for _, round := range rounds {
			cubes := strings.Split(round, ", ")
			for i := 0; i < len(cubes); i++ {
				colorInfo := strings.Split(cubes[i], " ")
				colorCount, e := strconv.Atoi(colorInfo[0])
				if e != nil {
					log.Fatalf("unable to convert color count to int: %v", e)
				}
				color := colorInfo[1]
				if colorCount > colors[color] {
					possible = false
					break
				}
			}
		}
		if possible {
			sum += gameNum
		}
	}
	fmt.Println("The answer to problem #1 is:", sum)
}

// #2
func p2() {
	file, err := os.Open("../inputs/day2.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		colors := map[string]int{"red": 0, "green": 0, "blue": 0}
		gameLog := scanner.Text()
		logComponents := strings.Split(gameLog, ": ")
		rounds := strings.Split(logComponents[1], "; ")
		for _, round := range rounds {
			cubes := strings.Split(round, ", ")
			for i := 0; i < len(cubes); i++ {
				colorInfo := strings.Split(cubes[i], " ")
				colorCount, e := strconv.Atoi(colorInfo[0])
				if e != nil {
					log.Fatalf("unable to convert color count to int: %v", e)
				}
				color := colorInfo[1]
				if colorCount > colors[color] {
					colors[color] = colorCount
				}
			}
		}
		sum += colors["red"] * colors["green"] * colors["blue"]
	}
	fmt.Println("The answer to problem #2 is:", sum)
}
