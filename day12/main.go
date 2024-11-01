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

	totalCount := CountArrangements(input)
	fmt.Println("total possible arrangements :", totalCount)
}

func CountArrangements(input *bufio.Scanner) int {
	totalCount := 0
	for input.Scan() {
		line := input.Text()
		parts := strings.Fields(line)
		springs := parts[0]
		expectedGroups := strings.FieldsFunc(parts[1], func(r rune) bool {
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
