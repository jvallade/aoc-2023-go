package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Group struct {
	start      int
	end        int
	nextGroups []Group
	valid      bool
}

// Define custom error types
type GroupSizeMismatch struct{}
type GroupNotFound struct{}
type FoundUnexpectedGroup struct{}

// Implement the Error method for each custom error type
func (e GroupSizeMismatch) Error() string {
	return "group size mismatch"
}
func (e GroupNotFound) Error() string {
	return "group not found"
}
func (e FoundUnexpectedGroup) Error() string {
	return "found unexpected group"
}

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

func FindNextPossibleGroup(springs string, groupSizes []int, offset int) ([]Group, error) {
	inGroup := false
	var start int

	// fmt.Println("finding new group in", springs)
	// fmt.Println("searching next group size", groupSizes)
	// fmt.Println("offset", offset)

	groups := make([]Group, 0)

	for i, c := range springs {
		switch c {
		case '.':
			// we found a complete group...
			if inGroup {
				// ... but we did not expect a new group (i.e groupSizes is empty)
				if len(groupSizes) == 0 {
					// we should not reach this point
					log.Fatal("we should not reach this point - this should be handled at group start")
				}

				// now we need to check if the group we found has the expected size
				if i-start == groupSizes[0] {
					// we found a group with the expected size
					group := Group{start + offset, i + offset, make([]Group, 0), true}

					// now we need to find the next possible groups from this one
					nextGroups, _ := FindNextPossibleGroup(springs[i+1:], groupSizes[1:], offset+i+1)
					for _, g := range nextGroups {
						if g.valid {
							group.nextGroups = append(group.nextGroups, g)
						}
					}
					groups = append(groups, group)
					return groups, nil
				} else {
					// we found a group but it does not have the expected size
					group := Group{valid: false}
					groups = append(groups, group)
					return groups, nil
				}
			}
		case '#':
			if len(groupSizes) == 0 {
				// we found a group but we did not expect any group
				group := Group{valid: false}
				groups = append(groups, group)
				return groups, nil
			}

			if !inGroup {
				inGroup = true
				start = i
			}
		case '?':
			if inGroup {
				if i-start > groupSizes[0] {
					// we found a group but it is too big
					group := Group{valid: false}
					groups = append(groups, group)
					return groups, nil

				} else if i-start == groupSizes[0] {
					group := Group{start + offset, i + offset, make([]Group, 0), true}

					// now we need to find the next possible groups from this one
					nextGroups, _ := FindNextPossibleGroup(springs[i+1:], groupSizes[1:], offset+i+1)
					for _, g := range nextGroups {
						if g.valid {
							group.nextGroups = append(group.nextGroups, g)
						}
					}
					groups = append(groups, group)
					return groups, nil
				}
			} else {
				if len(groupSizes) == 0 {
					// we do not expect any group - we can skip this character
					continue
				}

				// we start by skipping this character (i.e. we search a group starting at the next character)
				nextGroups, _ := FindNextPossibleGroup(springs[i+1:], groupSizes, offset+i+1)
				for _, g := range nextGroups {
					if g.valid {
						groups = append(groups, g)
					}
				}

				// then we start a new group
				start = i
				inGroup = true
			}
		}
	}

	if inGroup {
		if len(springs)-start == groupSizes[0] {
			group := Group{start + offset, len(springs) + offset, make([]Group, 0), true}
			if len(groupSizes) == 1 {
				// we found all the groups we were looking for
				groups = append(groups, group)
				return groups, nil
			} else {
				// we found a group but we did not find all the groups we were looking for
				group.valid = false
				groups = append(groups, group)
				return groups, nil
			}
		} else {
			group := Group{valid: false}
			groups = append(groups, group)
			return groups, nil
		}
	}

	if len(groupSizes) == 0 {
		// we did not find any group but we did not expect any group
		return groups, nil
	} else {
		// we did not find any group but we did expect a group
		group := Group{valid: false}
		groups = append(groups, group)
		return groups, nil
	}
}

func FlattenGroup(group Group) [][]Group {
	if len(group.nextGroups) == 0 {
		return [][]Group{{group}}
	}

	groups := make([][]Group, 0)
	for _, g := range group.nextGroups {
		flattenedGroups := FlattenGroup(g)
		for _, fg := range flattenedGroups {
			groups = append(groups, append([]Group{group}, fg...))
		}
	}
	return groups
}

func CountArrangements(input *bufio.Scanner, unfold bool) int {
	totalCount := 0
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

		foundGroups, err := FindNextPossibleGroup(springs, groupSizes, 0)
		if err != nil {
			fmt.Println(foundGroups, err)
			continue
		}

		count := 0
		for _, g := range foundGroups {
			fgs := FlattenGroup(g)
			for _, fg := range fgs {
				if len(fg) == len(expectedGroups) {
					count++
				}
			}
		}
		totalCount += count
	}
	return totalCount
}
