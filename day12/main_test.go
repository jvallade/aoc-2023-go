package main

import (
	"bufio"
	"strings"
	"testing"
)

const example = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

const example2 = `????.#...#... 4,1,1`

const expectResultPart1 = 21
const expectResultPart2 = 525152

func TestCounterPart1(t *testing.T) {
	input := bufio.NewScanner(strings.NewReader(example))
	totalCount := CountArrangements(input, false)
	if totalCount != expectResultPart1 {
		t.Errorf("got %d, expected %d", totalCount, expectResultPart1)
	}
}

func TestCounterPart2(t *testing.T) {
	input := bufio.NewScanner(strings.NewReader(example))
	totalCount := CountArrangements(input, true)
	if totalCount != expectResultPart2 {
		t.Errorf("got %d, expected %d", totalCount, expectResultPart2)
	}
}
func TestCounterPart3(t *testing.T) {
	input := bufio.NewScanner(strings.NewReader(example2))
	totalCount := CountArrangements(input, true)
	if totalCount != expectResultPart2 {
		t.Errorf("got %d, expected %d", totalCount, expectResultPart2)
	}
}
