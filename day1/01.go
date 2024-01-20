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
			if IsDigit(line[rightIndex]) && rightNumber == 0 {
				rightNumber = ToInt(line[rightIndex])
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
