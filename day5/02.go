package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := ParseInput(input)
	type submapping struct {
		source, size, offset int
	}
	type mapping []submapping

	convert := func(m mapping, in int) int {
		for _, c := range m {
			if in >= c.source && in <= c.source+c.size {
				return in + c.offset
			}
		}
		return in
	}
	splitRangesAt := func(s [][2]int, n int) [][2]int {
		for i, ss := range s {
			if n > ss[0] && n <= ss[1] {
				s[i][1] = n - 1
				s = append(s, [2]int{n, ss[1]})
				return s
			}
		}
		return s
	}

	var seedPairs [][2]int
	var mappings []mapping

	currentMapping := mapping{}
	for _, line := range s {
		if line == "" {
			continue
		}

		if strings.Contains(line, "seeds: ") {
			line = strings.ReplaceAll(line, "seeds: ", "")
			seeds := StringToIntArray(strings.Split(line, " "))
			for i := 0; i < len(seeds); i += 2 {
				seedPairs = append(seedPairs, [2]int{seeds[i], seeds[i] + seeds[i+1]})
			}
			continue
		}

		if strings.Contains(line, "-to-") {
			if len(currentMapping) > 0 {
				mappings = append(mappings, currentMapping)
			}
			currentMapping = mapping{}
			continue
		}

		values := StringToIntArray(strings.Split(line, " "))
		currentMapping = append(currentMapping, submapping{
			source: values[1],
			size:   values[2],
			offset: values[0] - values[1],
		})
	}
	if len(currentMapping) > 0 {
		mappings = append(mappings, currentMapping)
	}

	lowest := math.MaxInt
	for _, pair := range seedPairs {
		ranges := [][2]int{pair}

		for _, mapping := range mappings {
			for _, submapping := range mapping {
				ranges = splitRangesAt(ranges, submapping.source)
			}

			for i := range ranges {
				ranges[i][0] = convert(mapping, ranges[i][0])
				ranges[i][1] = convert(mapping, ranges[i][1])
			}
		}

		for i := range ranges {
			if lowest == -1 {
				lowest = ranges[i][0]
			}

			if (lowest > ranges[i][0] && ranges[i][0] != -1) || lowest == -1 {
				fmt.Println(ranges[i][0], i)
				lowest = min(lowest, ranges[i][0])
			}
		}
	}

	fmt.Print(lowest)
	return lowest
}
