package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal("input file not found")
	}

	schematic := createArraySchematic(&file)

	currentValue := 0
	total := 0
	hasCharAround := false
	for y, line := range schematic {
		for x, char := range line {
			if !isDigit(char) {
				if hasCharAround {
					fmt.Println("currentNumber: ", currentValue, " line: ", y)
					total += currentValue
				}
				hasCharAround = false
				currentValue = 0
				continue
			}

			if !hasCharAround {
				hasCharAround = checkAround(schematic, y, x, len(schematic), len(schematic[y]))
			}

			value, _ := strconv.Atoi(char)
			currentValue = currentValue*10 + value
		}
	}

	fmt.Printf("%v", total)

}

func createArraySchematic(file *[]byte) [][]string {
	characters := [][]string{}
	lines := strings.Split(string(*file), "\n")
	for _, line := range lines {
		var splitLine []string
		for _, char := range line {
			splitLine = append(splitLine, string(char))
		}
		characters = append(characters, splitLine)
	}
	return characters
}

func checkAround(schematic [][]string, y int, x int, ylength int, xLength int) bool {
	positions := [][]int{
		{-1, 0},  // Above
		{1, 0},   // Below
		{0, -1},  // Left
		{0, 1},   // Right
		{-1, -1}, // Diagonal top-left
		{-1, 1},  // Diagonal top-right
		{1, -1},  // Diagonal bottom-left
		{1, 1},   // Diagonal bottom-right
	}

	for _, pos := range positions {
		rowIndex, colIndex := y+pos[0], x+pos[1]

		if rowIndex >= 0 && rowIndex < ylength && colIndex >= 0 && colIndex < xLength {
			if schematic[rowIndex][colIndex] != "." && !isDigit(schematic[rowIndex][colIndex]) {
				return true
			}
		}
	}

	return false

}

func isDigit(char string) bool {
	_, err := strconv.Atoi(char)
	if err != nil {
		// If conversion fails, the character is not a number
		return false
	}

	return true
}
