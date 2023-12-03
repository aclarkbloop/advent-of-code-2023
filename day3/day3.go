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
	file, err := os.Open("../inputs/day3.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	grid := [][]string{}
	for scanner.Scan() {
		row := scanner.Text()
		rowComponents := strings.Split(row, "")
		grid = append(grid, rowComponents)
	}
	for i := 0; i < len(grid); i++ {
		foundNum := ""
		for j := 0; j < len(grid[i]); j++ {
			if isInt(grid[i][j]) {
				foundNum += grid[i][j]
				// check if value to the right exists
				if j+1 < len(grid[i]) {
					// check if value to the right is an int
					if isInt(grid[i][j+1]) {
						continue
					} else {
						isValid := isPartValid(grid, i, j, foundNum)
						if isValid {
							partNum, e := strconv.Atoi(foundNum)
							if e != nil {
								log.Fatalf("unable to convert part number to int: %v", e)
							} else {
								sum += partNum
							}
						}
					}
				} else {
					// we're at the end of the row
					partNum, e := strconv.Atoi(foundNum)
					if e != nil {
						// no foundNum to add
						log.Fatalf("unable to convert part number to int: %v", e)
						break
					} else {
						isValid := isPartValid(grid, i, j, foundNum)
						if isValid {
							sum += partNum
						}
					}
				}
			} else {
				// reset foundNum until we find another int
				foundNum = ""
			}
		}
	}

	fmt.Println("The answer to problem #1 is:", sum)
}

// #2
func p2() {
	file, err := os.Open("../inputs/day3.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	grid := [][]string{}
	gearAdjacentParts := []map[string]int{}
	for scanner.Scan() {
		row := scanner.Text()
		rowComponents := strings.Split(row, "")
		grid = append(grid, rowComponents)
	}
	for i := 0; i < len(grid); i++ {
		foundNum := ""
		for j := 0; j < len(grid[i]); j++ {
			if isInt(grid[i][j]) {
				foundNum += grid[i][j]
				// check if value to the right exists
				if j+1 < len(grid[i]) {
					// check if value to the right is an int
					if isInt(grid[i][j+1]) {
						continue
					} else {
						gearMap := updateGearMap(grid, i, j, foundNum, gearAdjacentParts)
						gearAdjacentParts = gearMap
					}
				} else {
					// we're at the end of the row
					if isInt(foundNum) {
						gearMap := updateGearMap(grid, i, j, foundNum, gearAdjacentParts)
						gearAdjacentParts = gearMap
					}
				}
			} else {
				// reset foundNum until we find another int
				foundNum = ""
			}
		}
	}

	foundGears := []map[string]int{}
	for i, gearLocationMapI := range gearAdjacentParts {
		y := gearLocationMapI["y"]
		x := gearLocationMapI["x"]
		dupeGear := false
		for _, gear := range foundGears {
			if gear["y"] == y && gear["x"] == x {
				dupeGear = true
			}
		}
		if dupeGear {
			continue
		}
		gearRatio := 0
		for j, gearLocationMapJ := range gearAdjacentParts {
			if i != j {
				if gearLocationMapJ["y"] == y && gearLocationMapJ["x"] == x {
					if gearRatio == 0 {
						foundGears = append(foundGears, gearLocationMapI)
						gearRatio = gearLocationMapI["partNum"] * gearLocationMapJ["partNum"]
					} else {
						// too many adjacent parts, not a gear
						gearRatio = 0
						break
					}
				}
			}
		}
		sum += gearRatio
	}

	fmt.Println("The answer to problem #2 is:", sum)
}

func isPartValid(grid [][]string, i int, j int, foundNum string) bool {
	isValid := false
	// check to the right
	if j+1 < len(grid[i]) {
		if grid[i][j+1] != "." && !isInt(grid[i][j+1]) {
			isValid = true
		}
	}
	// check to the left
	if j-len(foundNum) >= 0 {
		if grid[i][j-len(foundNum)] != "." && !isInt(grid[i][j-len(foundNum)]) {
			isValid = true
		}
	}
	// check above
	if i-1 >= 0 {
		for x := j - len(foundNum); x < j+2; x++ {
			if x >= 0 && x < len(grid[i]) && grid[i-1][x] != "." && !isInt(grid[i-1][x]) {
				isValid = true
			}
		}
	}
	// check below
	if i+1 < len(grid) {
		for x := j - len(foundNum); x < j+2; x++ {
			if x >= 0 && x < len(grid[i]) && grid[i+1][x] != "." && !isInt(grid[i+1][x]) {
				isValid = true
			}
		}
	}
	return isValid
}

func updateGearMap(grid [][]string, i int, j int, foundNum string, gearAdjacentParts []map[string]int) []map[string]int {
	// check to the right
	intPartNum, _ := strconv.Atoi(foundNum)
	if j+1 < len(grid[i]) {
		if grid[i][j+1] != "." && !isInt(grid[i][j+1]) {
			if grid[i][j+1] == "*" {
				gearAdjacentParts = append(gearAdjacentParts, map[string]int{"y": i, "x": j + 1, "partNum": intPartNum})
			}
		}
	}
	// check to the left
	if j-len(foundNum) >= 0 {
		if grid[i][j-len(foundNum)] != "." && !isInt(grid[i][j-len(foundNum)]) {
			if grid[i][j-len(foundNum)] == "*" {
				gearAdjacentParts = append(gearAdjacentParts, map[string]int{"y": i, "x": j - len(foundNum), "partNum": intPartNum})
			}
		}
	}
	// check above
	if i-1 >= 0 {
		for x := j - len(foundNum); x < j+2; x++ {
			if x >= 0 && x < len(grid[i]) && grid[i-1][x] != "." && !isInt(grid[i-1][x]) {
				if grid[i-1][x] == "*" {
					gearAdjacentParts = append(gearAdjacentParts, map[string]int{"y": i - 1, "x": x, "partNum": intPartNum})
				}
			}
		}
	}
	// check below
	if i+1 < len(grid) {
		for x := j - len(foundNum); x < j+2; x++ {
			if x >= 0 && x < len(grid[i]) && grid[i+1][x] != "." && !isInt(grid[i+1][x]) {
				if grid[i+1][x] == "*" {
					gearAdjacentParts = append(gearAdjacentParts, map[string]int{"y": i + 1, "x": x, "partNum": intPartNum})
				}
			}
		}
	}
	return gearAdjacentParts
}

func isInt(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}
