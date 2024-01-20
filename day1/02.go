package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	b, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal("input file not found")
	}

	t := 0
	for _, line := range strings.Split(string(b), "\n") {
		l, r, ld, rd := 0, len(line)-1, 0, 0
		for (ld == 0 || rd == 0) && len(line) > 0 {
			if IsDigit(line[l]) && ld == 0 {
				ld = ToInt(line[l])
			}

			lw, found := checkWords(line, l)
			if found && ld == 0 {
				ld = lw
			}

			if IsDigit(line[r]) && rd == 0 {
				rd = ToInt(line[r])
			}

			rw, found := checkWords(line, r)
			if found && rd == 0 {
				rd = rw
			}

			if ld != 0 && rd != 0 {
				t += (ld * 10) + rd
				break
			}
			l++
			r--
		}
	}

	fmt.Printf("%v", t)
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
