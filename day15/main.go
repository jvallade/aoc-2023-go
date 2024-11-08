package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	// inputFile, err = os.Open("day15/input.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// input = bufio.NewScanner(inputFile)
	// res = Part2(input)
	// fmt.Println("part 2 :", res)
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

func Part2(input *bufio.Scanner) int {
	res := 0
	return res
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
