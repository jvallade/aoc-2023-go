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
	inputFile, err := os.Open("day15/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := bufio.NewScanner(inputFile)
	res := Part1(input)
	fmt.Println("part 1 :", res)

	inputFile, err = os.Open("day15/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input = bufio.NewScanner(inputFile)
	res = Part2(input)
	fmt.Println("part 2 :", res)
}

func Part1(input *bufio.Scanner) int {
	res := 0

	for input.Scan() {
		inputs := strings.Split(input.Text(), ",")
		for _, i := range inputs {
			res += Hash(i)
		}
	}

	return res
}

type Box struct {
	Lens *Lens
}

func (b Box) String() string {
	res := ""
	nextLens := b.Lens
	for {
		if nextLens == nil {
			break
		}
		res += fmt.Sprintf("(%s, %d) -> ", nextLens.Label, nextLens.FocalLength)
		nextLens = nextLens.Next
	}
	return res
}

type Lens struct {
	Label       string
	FocalLength int
	Next        *Lens
}

func Part2(input *bufio.Scanner) int {

	boxes := make([]Box, 256)
	for i := range boxes {
		boxes[i] = Box{nil}
	}

	for input.Scan() {
		inputs := strings.Split(input.Text(), ",")
		for _, i := range inputs {
			label, operation := ExtractLabel(i)
			boxNumber := Hash(label)
			switch operation {
			case '=':
				focalLength := ExtractFocalLength(i)
				currentLens := boxes[boxNumber].Lens
				if currentLens == nil {
					boxes[boxNumber].Lens = &Lens{label, focalLength, nil}
				} else {
					for {
						nextLens := currentLens.Next
						if currentLens.Label == label {
							currentLens.FocalLength = focalLength
							break
						}

						if nextLens == nil {
							currentLens.Next = &Lens{label, focalLength, nil}
							break
						} else {
							currentLens = nextLens
						}
					}
				}

			case '-':
				currentLens := boxes[boxNumber].Lens
				if currentLens == nil {
					continue
				} else if currentLens.Label == label {
					boxes[boxNumber].Lens = currentLens.Next
					continue
				}

				for {
					nextLens := currentLens.Next
					if nextLens == nil {
						// we have reach the end without finding the label
						break
					}
					if nextLens.Label == label {
						currentLens.Next = nextLens.Next
						break
					} else {
						currentLens = nextLens
					}
				}
			}

		}
	}

	res := 0
	for i, b := range boxes {
		currentLens := b.Lens
		slot := 1
		if currentLens != nil {
			res += (i + 1) * slot * currentLens.FocalLength
		} else {
			continue
		}
		for {
			nextLens := currentLens.Next
			if nextLens == nil {
				break
			}
			slot++
			res += (i + 1) * slot * nextLens.FocalLength
			currentLens = nextLens
		}
	}

	return res
}

func ExtractLabel(input string) (string, rune) {
	if strings.Contains(input, "=") {
		return strings.Split(input, "=")[0], '='
	} else if strings.Contains(input, "-") {
		return strings.Split(input, "-")[0], '-'
	} else {
		log.Fatal("Invalid input")
		return "", '?'
	}
}

func ExtractFocalLength(input string) int {
	l := strings.Split(input, "=")[1]
	result, err := strconv.Atoi(l)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func Hash(input string) int {
	res := 0
	for _, c := range input {
		res += int(c)
		res *= 17
		res %= 256
	}
	return res
}
