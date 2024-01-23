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
	total := 0
	for i, line := range lines {
		str := strings.Split(line, ":")[1:][0]
		rowSplit := strings.Split(str, "|")
		answers := createArray(rowSplit[0])
		values := createArray(rowSplit[1])

		total += getPoints(answers, values)

		fmt.Println("game ", i, " total", total)
	}

	fmt.Println(total)

}

func createArray(line string) []int {
	array := []int{}
	for _, val := range strings.Fields(line) {
		value, _ := strconv.Atoi(val)
		array = append(array, value)
	}
	return array
}

func getPoints(answers []int, values []int) int {
	counter := 0
	for _, value := range values {
		for _, answer := range answers {
			if answer == value {
				if counter >= 2 {
					counter *= 2
				} else {
					counter++
				}
			}
		}
	}

	return counter
}
