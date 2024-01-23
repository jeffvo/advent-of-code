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

	lines := strings.Split(string(file), "\n")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		cardNumber, _ := strconv.Atoi(strings.Fields(strings.Split(line, ":")[:1][0])[1])
		str := strings.Split(line, ":")[1:][0]
		rowSplit := strings.Split(str, "|")
		answers := createArray(rowSplit[0])
		values := createArray(rowSplit[1])

		winningNumbers := winningNumbers(answers, values)

		for x := 1; x <= winningNumbers; x++ {
			lines = append(lines, lines[cardNumber+x-1])
		}

	}
	fmt.Println(len(lines))

}

func createArray(line string) []int {
	array := []int{}
	for _, val := range strings.Fields(line) {
		value, _ := strconv.Atoi(val)
		array = append(array, value)
	}
	return array
}

func winningNumbers(answers []int, values []int) int {
	counter := 0
	for _, value := range values {
		for _, answer := range answers {
			if answer == value {
				counter++
			}
		}
	}

	return counter
}
