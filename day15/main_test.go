package main

import (
	"bufio"
	"strings"
	"testing"
)

const example string = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

const expectedResultPart1 int = 1320
const expectedResultPart2 int = 145

func TestHash(t *testing.T) {
	input := "HASH"
	res := Hash(input)
	if res != 52 {
		t.Errorf("got %d, expected %d", res, 52)
	}
}

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
