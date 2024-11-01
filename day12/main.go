package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile, err := os.Open("day12/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := bufio.NewScanner(inputFile)

	unfold := false
	if len(os.Args) > 1 {
		part := os.Args[1]
		if part == "part2" {
			unfold = true
		} else {
			unfold = false
		}
	}

	totalCount := CountArrangements(input, unfold)
	fmt.Println("total possible arrangements :", totalCount)
}

func Unfold(springs, groups string) (string, string) {
	totSprings := springs
	totGroups := groups
	for range 4 {
		totSprings += "?" + springs
		totGroups += "," + groups
	}
	return totSprings, totGroups
}

func CountArrangements(input *bufio.Scanner, unfold bool) int {
	totalCount := 0
	for input.Scan() {
		line := input.Text()
		parts := strings.Fields(line)
		var springs string
		var groups string
		if unfold {
			springs, groups = Unfold(parts[0], parts[1])
		} else {
			springs = parts[0]
			groups = parts[1]
		}
		fmt.Println(springs, groups)

		expectedGroups := strings.FieldsFunc(groups, func(r rune) bool {
			return r == ','
		})

		possibleRawArrangements := []string{""}
		for _, c := range springs {
			switch {
			case c == '.' || c == '#':
				for i := range possibleRawArrangements {
					possibleRawArrangements[i] += string(c)
				}
			case c == '?':
				for i, rawArrangement := range possibleRawArrangements {
					possibleRawArrangements[i] = rawArrangement + "."
					possibleRawArrangements = append(possibleRawArrangements, rawArrangement+"#")
				}
			}
		}

		for _, rawArrangement := range possibleRawArrangements {
			groups := make([]int, 0)
			inGroup := false
			var start int
			for i, c := range rawArrangement {
				if c == '.' {
					if inGroup {
						groups = append(groups, i-start)
						inGroup = false
					}
				} else {
					if !inGroup {
						start = i
						inGroup = true
					}
				}
			}
			if inGroup {
				groups = append(groups, len(rawArrangement)-start)
			}
			if len(groups) != len(expectedGroups) {
				continue
			}
			mismatch := false
			for i, group := range groups {
				expectedGroup, err := strconv.Atoi(expectedGroups[i])
				if err != nil {
					log.Fatal(err)
				}
				if group != expectedGroup {
					mismatch = true
					continue
				}
			}
			if !mismatch {
				totalCount++
			}
		}
	}
	return totalCount
}
