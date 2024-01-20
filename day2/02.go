package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input_2.txt")

	if err != nil {
		log.Fatal("input file could not found")
	}

	total := 0
	for _, line := range strings.Split(string(file), "\n") {
		split := strings.Split(line, ":")
		minRed, minGreen, minBlue := 0, 0, 0

		games := strings.Split(split[1], ";")

		for _, game := range games {
			dices := strings.Split(game, ",")
			for _, dice := range dices {
				diceSplit := strings.Fields(dice)
				diceColour := diceSplit[1]

				diceValue, _ := strconv.Atoi(diceSplit[0])

				switch diceColour {
				case "red":
					minRed = max(diceValue, minRed)
				case "green":
					minGreen = max(diceValue, minGreen)
				case "blue":
					minBlue = max(diceValue, minBlue)
				default:
					panic(fmt.Sprintf("unknown color %s", diceColour))
				}
			}

		}
		total += minRed * minBlue * minGreen
	}
	fmt.Printf("%v", total)
}
