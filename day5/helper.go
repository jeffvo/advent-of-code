package main

import (
	"strconv"
	"strings"
)

func ParseInput(s string, on ...string) []string {
	sep := "\n"
	if len(on) > 0 {
		sep = on[0]
	}
	return strings.Split(strings.Trim(strings.TrimSpace(s), sep), sep)
}

func StringToIntArray(s []string) []int {
	ints := make([]int, len(s))
	for i, str := range s {
		ints[i], _ = strconv.Atoi(str)
	}
	return ints
}
