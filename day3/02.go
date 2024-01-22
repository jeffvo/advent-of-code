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

	total := 0

	for x, line := range schematic {
		for y, char := range line {

			if char == "*" {
				total += checkAround(schematic, x, y, len(schematic), len(schematic[x]))
			}
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

func checkAround(schematic [][]string, x int, y int, xLength int, yLength int) int {
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

	firstValue, secondValue := 0, 0

	for _, pos := range positions {
		rowIndex, colIndex := x+pos[0], y+pos[1]

		value := 0
		str := ""
		if rowIndex >= 0 && rowIndex < xLength && colIndex >= 0 && colIndex < yLength && isDigit(schematic[rowIndex][colIndex]) {
			str = schematic[rowIndex][colIndex]
			if colIndex-1 >= 0 && isDigit(schematic[rowIndex][colIndex-1]) {
				str = schematic[rowIndex][colIndex-1] + str

				if colIndex-2 >= 0 && isDigit(schematic[rowIndex][colIndex-2]) {
					str = schematic[rowIndex][colIndex-2] + str
				}
			}

			if colIndex+1 < yLength && isDigit(schematic[rowIndex][colIndex+1]) {
				str = str + schematic[rowIndex][colIndex+1]

				if colIndex+2 < yLength && isDigit(schematic[rowIndex][colIndex+2]) {
					str = str + schematic[rowIndex][colIndex+2]
				}

			}

			char, _ := strconv.Atoi(str)
			value = char

		}

		if firstValue == 0 {
			firstValue = value
		} else if secondValue == 0 && value != firstValue {
			secondValue = value
		}

		if firstValue > 0 && secondValue > 0 {
			return firstValue * secondValue
		}
	}

	return 0

}

func isDigit(char string) bool {
	_, err := strconv.Atoi(char)
	if err != nil {
		// If conversion fails, the character is not a number
		return false
	}

	return true
}
