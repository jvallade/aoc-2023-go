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

	inputFile, err = os.Open("day13/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input = bufio.NewScanner(inputFile)
	res = Part2(input)
	fmt.Println("part 2 :", res)
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

func Part2(input *bufio.Scanner) int {
	res := 0
	currentPattern := make([]string, 0)
	for input.Scan() {
		if input.Text() == "" {
			res += processLinesWithSmudge(currentPattern)
			res += processColumnsWithSmudge(currentPattern)
			currentPattern = make([]string, 0)
			continue
		}

		currentPattern = append(currentPattern, input.Text())
	}

	res += processLinesWithSmudge(currentPattern)
	res += processColumnsWithSmudge(currentPattern)
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

func processLinesWithSmudge(lines []string) int {
	// possiblePositionsWithSmudge := make([][]int, len(lines[0])-1)
	// for i, line := range lines {
	// 	possiblePositionsWithSmudge[i] = searchSymWithSmudge(line)
	// }
	// possiblePositions := make([]int, len(lines[0])-1)
	// for i := range possiblePositions {
	// 	possiblePositions[i] = i
	// }

	symWithSmudge := make(map[int]struct{}, 0)
	for i, line := range lines {
		possiblePositions := searchSymWithSmudge(line)
		for j, line := range lines {
			if i != j {
				possiblePositions = searchSym(line, possiblePositions)
			}
		}
		for _, pos := range possiblePositions {
			symWithSmudge[pos] = struct{}{}
		}
	}

	res := 0
	count := 0
	for pos := range symWithSmudge {
		count++
		res += pos
	}
	if count == 0 {
		// fmt.Println("no symetric pattern with smudge found on lines")
		return 0
	}
	if count > 1 {
		fmt.Println("multiple symetric patterns with smudge found on lines", symWithSmudge)
		return 0
	}

	// add + 1 to get the number of columns
	return res + 1
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

func processColumnsWithSmudge(lines []string) int {
	symWithSmudge := make(map[int]struct{}, 0)

	patterns := make([]string, len(lines[0]))
	for i := 0; i < len(lines[0]); i++ {
		pattern := ""
		for _, line := range lines {
			pattern += string(line[i])
		}
		patterns[i] = pattern
	}

	for i, pattern := range patterns {
		possiblePositions := searchSymWithSmudge(pattern)
		for j, pattern := range patterns {
			if i != j {
				possiblePositions = searchSym(pattern, possiblePositions)
			}
		}

		for _, pos := range possiblePositions {
			symWithSmudge[pos] = struct{}{}
		}
	}

	res := 0
	count := 0
	for pos := range symWithSmudge {
		count++
		res += pos
	}
	if count == 0 {
		// fmt.Println("no symetric pattern with smudge found on columns")
		return 0
	}
	if count > 1 {
		fmt.Println("multiple symetric patterns with smudge found on columns", symWithSmudge)
		return 0
	}
	return 100 * (res + 1)
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

func searchSymWithSmudge(pattern string) []int {
	res := make([]int, 0)
	for i := 0; i < len(pattern)-1; i++ {
		maxJ := min(i+1, len(pattern)-i-1)
		foundSym := true
		count := 0
		for j := 0; j < maxJ; j++ {
			if pattern[i-j] != pattern[i+1+j] {
				count++
			}
			if count > 1 {
				foundSym = foundSym && pattern[i-j] == pattern[i+1+j]
			}
		}
		if foundSym && count == 1 {
			res = append(res, i)
		}
	}
	return res
}
