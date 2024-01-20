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
			if IsDigit(line[r]) && rd == 0 {
				rd = ToInt(line[r])
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
