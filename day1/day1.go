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
	file, err := os.Open("../inputs/day1.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		entry := scanner.Text()
		for i := 0; i < len(entry); i++ {
			digitOne, e := strconv.Atoi(string(entry[i]))
			if e == nil {
				sum += digitOne * 10
				break
			}
		}
		for i := len(entry) - 1; i >= 0; i-- {
			digitTwo, e := strconv.Atoi(string(entry[i]))
			if e == nil {
				sum += digitTwo
				break
			}
		}
	}
	fmt.Println(sum)
}

// #2
func p2() {
	file, err := os.Open("../inputs/day1.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	numStrings := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	sum := 0
	for scanner.Scan() {
		entry := scanner.Text()
		foundDigitOne := false
		foundDigitTwo := false
		numStringOne := ""
		for i := 0; i < len(entry); i++ {
			numStringOne += string(entry[i])
			for i, numString := range numStrings {
				if strings.HasSuffix(numStringOne, numString) {
					sum += (i + 1) * 10
					foundDigitOne = true
					break
				}
			}
			if foundDigitOne {
				break
			}
			digitOne, e := strconv.Atoi(string(entry[i]))
			if e == nil {
				sum += digitOne * 10
				break
			}
		}
		numStringTwo := ""
		for i := len(entry) - 1; i >= 0; i-- {
			numStringTwo = string(entry[i]) + numStringTwo
			for i, numString := range numStrings {
				if strings.HasPrefix(numStringTwo, numString) {
					sum += i + 1
					foundDigitTwo = true
					break
				}
			}
			if foundDigitTwo {
				break
			}
			digitTwo, e := strconv.Atoi(string(entry[i]))
			if e == nil {
				sum += digitTwo
				break
			}
		}
	}
	fmt.Println(sum)
}
