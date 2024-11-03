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
	fmt.Println("memo hits :", memoHits)
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

type MemoInput struct {
	springs    string
	groupSizes string
}

func CreateMemoInput(springs string, groupSizes []int) MemoInput {
	sizes := ""
	for _, size := range groupSizes {
		sizes += strconv.Itoa(size) + ","
	}
	return MemoInput{springs, sizes}
}

var memoHits int = 0

func FindNextPossibleGroup(springs string, groupSizes []int, memo map[MemoInput]int) int {

	// fmt.Println("finding new group in", springs)
	// fmt.Println("searching next group size", groupSizes)
	// fmt.Println("offset", offset)

	if len(groupSizes) == 0 {
		// we do not expect any group
		if !strings.Contains(springs, "#") {
			return 1
		}
		// fmt.Println("no group left")
		return 0
	}

	memoInput := CreateMemoInput(springs, groupSizes)
	if val, ok := memo[memoInput]; ok {
		// fmt.Println("found in memo")
		memoHits++
		return val
	}

	requiredSize := 0
	for _, size := range groupSizes {
		requiredSize += size
	}
	requiredSize += len(groupSizes) - 1
	if len(springs) < requiredSize {
		// fmt.Println("not enough springs left")
		return 0
	}

	// we can search from the next character
	if springs[0] == '.' {
		return FindNextPossibleGroup(springs[1:], groupSizes, memo)
	}

	currentGroupSize := groupSizes[0]
	allBrokenSprings := !strings.ContainsAny(springs[:currentGroupSize], ".")
	var lastCharValid bool
	if len(springs) == currentGroupSize {
		lastCharValid = true
	} else {
		// the character after the group must be a '.' or a '?'
		lastCharValid = springs[currentGroupSize] != '#'
	}

	arrangementCount := 0

	if allBrokenSprings && lastCharValid {
		nextIndex := min(currentGroupSize+1, len(springs))
		arrangementCount += FindNextPossibleGroup(springs[nextIndex:], groupSizes[1:], memo)
	}

	// we need to check starting one char further as this '?' could be a '.'
	if springs[0] == '?' {
		arrangementCount += FindNextPossibleGroup(springs[1:], groupSizes, memo)
	}

	memo[memoInput] = arrangementCount
	return arrangementCount
}

func CountArrangements(input *bufio.Scanner, unfold bool) int {
	globalCounter := 0
	memo := make(map[MemoInput]int)
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
		// Convert the slice of strings to a slice of integers
		groupSizes := make([]int, len(expectedGroups))
		for i, groupSize := range expectedGroups {
			size, err := strconv.Atoi(groupSize)
			if err != nil {
				log.Fatal(err)
			}
			groupSizes[i] = size
		}
		count := FindNextPossibleGroup(springs, groupSizes, memo)
		fmt.Println("count", count)
		globalCounter += count
	}
	return globalCounter
}
