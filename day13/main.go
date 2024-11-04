package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	inputFile, err := os.Open("day13/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := bufio.NewScanner(inputFile)
	res := Part1(input)
	fmt.Println("part 1 :", res)
}

func Part1(input *bufio.Scanner) int {
	res := 0
	currentPattern := make([]string, 0)
	for input.Scan() {
		if input.Text() == "" {
			res += processLines(currentPattern)
			res += processColumns(currentPattern)
			currentPattern = make([]string, 0)
			continue
		}

		currentPattern = append(currentPattern, input.Text())
	}

	res += processLines(currentPattern)
	res += processColumns(currentPattern)
	return res
}

func processLines(lines []string) int {
	possiblePositions := make([]int, len(lines[0])-1)
	for i := range possiblePositions {
		possiblePositions[i] = i
	}

	for _, line := range lines {
		possiblePositions = searchSym(line, possiblePositions)
		if len(possiblePositions) == 0 {
			// fmt.Println("no symetric pattern found on lines")
			return 0
		}
	}
	if len(possiblePositions) > 1 {
		fmt.Println("multiple symetric patterns found on lines", possiblePositions)
	}

	// add + 1 to get the number of columns
	return possiblePositions[0] + 1
}

func processColumns(lines []string) int {
	possiblePositions := make([]int, len(lines)-1)
	for i := range possiblePositions {
		possiblePositions[i] = i
	}

	for i := 0; i < len(lines[0]); i++ {
		pattern := ""
		for _, line := range lines {
			pattern += string(line[i])
		}
		possiblePositions = searchSym(pattern, possiblePositions)
		if len(possiblePositions) == 0 {
			// fmt.Println("no symetric pattern found on columns")
			return 0
		}
	}
	if len(possiblePositions) > 1 {
		fmt.Println("multiple symetric patterns found on columns", possiblePositions)
	}
	return 100 * (possiblePositions[0] + 1)
}

func searchSym(pattern string, possiblePositions []int) []int {
	res := make([]int, 0)
	for _, i := range possiblePositions {
		maxJ := min(i+1, len(pattern)-i-1)
		foundSym := true
		for j := 0; j < maxJ; j++ {
			foundSym = foundSym && pattern[i-j] == pattern[i+1+j]
		}
		if foundSym {
			res = append(res, i)
		}
	}
	return res
}
