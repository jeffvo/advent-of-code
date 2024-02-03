package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type submapping struct {
	rangeStart, sourceStart, offset int
}

//go:embed input.txt
var input string

func main() {
	file := parseInput(input)

	type mapping []submapping
	var seeds []int
	var mappings []mapping

	currentMapping := mapping{}
	for _, line := range file {
		if line == "" {
			continue
		}

		if strings.Contains(line, "seeds: ") {
			line = strings.ReplaceAll(line, "seeds: ", "")
			seeds = stringToIntArray(strings.Split(line, " "))
			continue
		}

		if strings.Contains(line, "-to-") {
			if len(currentMapping) > 0 {
				mappings = append(mappings, currentMapping)
			}
			currentMapping = mapping{}
			continue
		}

		values := stringToIntArray(strings.Split(line, " "))
		currentMapping = append(currentMapping, submapping{
			rangeStart:  values[1],
			sourceStart: values[2],
			offset:      values[0] - values[1],
		})
	}
	if len(currentMapping) > 0 {
		mappings = append(mappings, currentMapping)
	}

	lowest := math.MaxInt
	for _, seed := range seeds {
		val := seed

		for _, mapping := range mappings {
			for _, submapping := range mapping {
				if val >= submapping.rangeStart && val <= submapping.rangeStart+submapping.sourceStart {
					val += submapping.offset
					break
				}
			}
		}
		lowest = min(lowest, val)
	}

	fmt.Println(lowest)

}

func parseInput(s string, on ...string) []string {
	sep := "\n"
	if len(on) > 0 {
		sep = on[0]
	}
	return strings.Split(strings.Trim(strings.TrimSpace(s), sep), sep)
}

func stringToIntArray(s []string) []int {
	ints := make([]int, len(s))
	for i, str := range s {
		ints[i], _ = strconv.Atoi(str)
	}
	return ints
}
