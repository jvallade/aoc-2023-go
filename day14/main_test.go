package main

import (
	"bufio"
	"strings"
	"testing"
)

const example string = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

const expectedResultPart1 int = 136
const expectedResultPart2 int = 64

// .....#....
// ....#...O#  9
// .....##...  0
// ..O#......  7
// ......OO#.  12
// .O#...O#.#  10
// ....O#....  4
// ......OOOO  12
// #...O###.O  4
// #.OOO#..OO  5

func TestPart1(t *testing.T) {
	input := strings.NewReader(example)
	res := Part1(bufio.NewScanner(input))
	if res != expectedResultPart1 {
		t.Errorf("got %d, expected %d", res, expectedResultPart1)
	}
}

func TestPart2(t *testing.T) {
	input := strings.NewReader(example)
	res := Part2(bufio.NewScanner(input))
	if res != expectedResultPart2 {
		t.Errorf("got %d, expected %d", res, expectedResultPart2)
	}
}
