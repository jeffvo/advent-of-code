package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input_1.txt")

	if err != nil {
		log.Fatal("input file could not found")
	}

	total := 0
	for _, line := range strings.Split(string(file), "\n") {
		split := strings.Split(line, ":")

		gameNumber, _ := strconv.Atoi(strings.TrimPrefix(split[0], "Game "))
		isValid := true

		games := strings.Split(split[1], ";")
	out:
		for _, game := range games {
			dices := strings.Split(game, ",")
			for _, dice := range dices {
				if !isValid {
					break out
				}
				diceSplit := strings.Fields(dice)
				diceColour := diceSplit[1]

				diceValue, _ := strconv.Atoi(diceSplit[0])

				switch diceColour {
				case "red":
					isValid = diceValue <= 12
				case "green":
					isValid = diceValue <= 13
				case "blue":
					isValid = diceValue <= 14
				default:
					panic(fmt.Sprintf("unknown color %s", diceColour))
				}
			}

		}
		if isValid {
			total += gameNumber
		}
	}
	fmt.Printf("%v", total)
}
