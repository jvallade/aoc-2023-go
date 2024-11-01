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

const expectResult = 21

func TestCounter(t *testing.T) {
	input := bufio.NewScanner(strings.NewReader(example))
	totalCount := CountArrangements(input)
	if totalCount != expectResult {
		t.Errorf("got %d, expected %d", totalCount, expectResult)
	}
}
