package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal("input file not found")
	}

	total := 0
	for _, line := range strings.Split(string(file), "\n") {
		leftIndex, rightIndex, leftNumber, rightNumber := 0, len(line)-1, 0, 0
		for (leftNumber == 0 || rightNumber == 0) && len(line) > 0 {
			if IsDigit(line[leftIndex]) && leftNumber == 0 {
				leftNumber = ToInt(line[leftIndex])
			}

			lw, found := checkWords(line, leftIndex)
			if found && leftNumber == 0 {
				leftNumber = lw
			}

			if IsDigit(line[rightIndex]) && rightNumber == 0 {
				rightNumber = ToInt(line[rightIndex])
			}

			rw, found := checkWords(line, rightIndex)
			if found && rightNumber == 0 {
				rightNumber = rw
			}

			if leftNumber != 0 && rightNumber != 0 {
				total += (leftNumber * 10) + rightNumber
				break
			}
			leftIndex++
			rightIndex--
		}
	}

	fmt.Printf("%v", total)
}

func checkWords(line string, i int) (int, bool) {

	words := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for j, word := range words {
		if checkWord(line, i, word) {
			return j, true
		}
	}
	return 0, false
}

func checkWord(line string, i int, word string) bool {
	if i+len(word) > len(line) {
		return false
	}
	return line[i:i+len(word)] == word
}
