package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

const N_CYCLES = 1000000000

type Matrix [][]rune

func main() {
	inputFile, err := os.Open("day14/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := bufio.NewScanner(inputFile)
	res := Part1(input)
	fmt.Println("part 1 :", res)

	inputFile, err = os.Open("day14/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input = bufio.NewScanner(inputFile)
	res = Part2(input)
	fmt.Println("part 2 :", res)
}

func Part1(input *bufio.Scanner) int {
	res := 0
	columns := make([]string, 0)
	firstLine := true
	for input.Scan() {
		line := input.Text()
		for i := 0; i < len(line); i++ {
			if firstLine {
				columns = append(columns, "")
			}
			columns[i] += string(line[i])
		}
		firstLine = false
	}

	for _, column := range columns {
		// fmt.Println("raw   ", column)
		roundRockBuffer := 0
		offset := 0
		tiltedColumn := []rune(strings.Repeat(".", len(column)))

		for k, c := range column {
			switch c {
			case 'O':
				roundRockBuffer++
			case '#':
				for i := 0; i < roundRockBuffer; i++ {
					tiltedColumn[i+offset] = 'O'
					res += len(column) - i - offset
				}
				roundRockBuffer = 0
				tiltedColumn[k] = '#'
				offset = k + 1
			case '.':
			}
		}
		for i := 0; i < roundRockBuffer; i++ {
			tiltedColumn[i+offset] = 'O'
			res += len(column) - i - offset
		}

		// fmt.Println("tilted", string(tiltedColumn))
		// fmt.Println("res   ", res)
	}

	return res
}

func Part2(input *bufio.Scanner) int {
	res := 0
	matrix := createMatrix(input)

	memo := make([]string, 0)
	memo = append(memo, matrix.hash())
	memoCount := make(map[string]int)
	indexFirstRepeat := 0
	nCycle := 0

	for range N_CYCLES {
		matrix.tiltNorth()
		matrix.tiltWest()
		matrix.tiltSouth()
		matrix.tiltEast()
		nCycle++

		hash := matrix.hash()
		if slices.Contains(memo, hash) {
			indexFirstRepeat = slices.Index(memo, hash)
			break
		} else {
			memo = append(memo, matrix.hash())
			memoCount[hash] = matrix.count()
		}
	}

	cycleLength := nCycle - indexFirstRepeat
	remainingCycles := (N_CYCLES - nCycle) % cycleLength

	hash := memo[indexFirstRepeat+remainingCycles]
	res = memoCount[hash]
	return res
}

func createMatrix(input *bufio.Scanner) Matrix {
	matrix := make([][]rune, 0)
	j := 0
	for input.Scan() {
		line := input.Text()
		matrix = append(matrix, make([]rune, len(line)))
		for i := 0; i < len(line); i++ {
			matrix[j][i] = rune(line[i])
		}
		j++
	}
	return matrix
}

func (m *Matrix) tiltNorth() {
	// iterate columns first
	for i := 0; i < len((*m)[0]); i++ {
		roundRockBuffer := 0
		offset := 0

		for j := 0; j < len((*m)); j++ {
			switch (*m)[j][i] {
			case 'O':
				(*m)[j][i] = '.'
				roundRockBuffer++
			case '#':
				for k := 0; k < roundRockBuffer; k++ {
					(*m)[k+offset][i] = 'O'
				}
				roundRockBuffer = 0
				offset = j + 1
			case '.':
			}

		}
		for k := 0; k < roundRockBuffer; k++ {
			(*m)[k+offset][i] = 'O'
		}
	}
}

func (m *Matrix) tiltSouth() {
	// iterate columns first
	for i := 0; i < len((*m)[0]); i++ {
		roundRockBuffer := 0

		for j := 0; j < len(*m); j++ {
			switch (*m)[j][i] {
			case 'O':
				(*m)[j][i] = '.'
				roundRockBuffer++
			case '#':
				for k := 0; k < roundRockBuffer; k++ {
					(*m)[j-1-k][i] = 'O'
				}
				roundRockBuffer = 0
			case '.':
			}
		}
		for k := 0; k < roundRockBuffer; k++ {
			(*m)[len(*m)-1-k][i] = 'O'
		}
	}
}

func (m *Matrix) tiltEast() {
	for i := 0; i < len(*m); i++ {
		roundRockBuffer := 0

		for j := 0; j < len((*m)[i]); j++ {
			switch (*m)[i][j] {
			case 'O':
				(*m)[i][j] = '.'
				roundRockBuffer++
			case '#':
				for k := 0; k < roundRockBuffer; k++ {
					(*m)[i][j-1-k] = 'O'
				}
				roundRockBuffer = 0
			case '.':
			}
		}
		for k := 0; k < roundRockBuffer; k++ {
			(*m)[i][len(*m)-1-k] = 'O'
		}
	}
}

func (m *Matrix) tiltWest() {
	for i := 0; i < len(*m); i++ {
		roundRockBuffer := 0
		offset := 0

		for j := 0; j < len((*m)[i]); j++ {
			switch (*m)[i][j] {
			case 'O':
				(*m)[i][j] = '.'
				roundRockBuffer++
			case '#':
				for k := 0; k < roundRockBuffer; k++ {
					(*m)[i][k+offset] = 'O'
				}
				roundRockBuffer = 0
				offset = j + 1
			case '.':
			}
		}
		for k := 0; k < roundRockBuffer; k++ {
			(*m)[i][k+offset] = 'O'
		}
	}
}

func (m *Matrix) print() {
	for _, row := range *m {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func (m Matrix) count() int {
	res := 0
	for i, row := range m {
		for _, c := range row {
			if c == 'O' {
				res += len(m) - i
			}
		}
	}
	return res
}

func (m Matrix) hash() string {
	res := ""
	for _, row := range m {
		for _, c := range row {
			res += string(c)
		}
	}
	return res
}
