package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

func main() {
	inputFile, err := os.Open("day16/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := bufio.NewScanner(inputFile)
	res := Part1(input)
	fmt.Println("part 1 :", res)

	inputFile, err = os.Open("day16/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input = bufio.NewScanner(inputFile)
	res = Part2(input)
	fmt.Println("part 2 :", res)
}

type Tile struct {
	TileType    rune
	Energized   bool
	Energizedby []*Tile
	Direction   []string
	X           int
	Y           int
}

type Matrix [][]Tile

func Part1(input *bufio.Scanner) int {
	matrix := make(Matrix, 0)
	j := 0
	for input.Scan() {
		row := make([]Tile, 0)
		for i, c := range input.Text() {
			row = append(row, Tile{TileType: c, Energized: false, X: j, Y: i})
		}
		matrix = append(matrix, row)
		j++
	}

	matrix.energize(0, 0, "right")
	return matrix.count()
}

func Part2(input *bufio.Scanner) int {
	matrix := make(Matrix, 0)
	j := 0
	for input.Scan() {
		row := make([]Tile, 0)
		for i, c := range input.Text() {
			row = append(row, Tile{TileType: c, Energized: false, X: j, Y: i})
		}
		matrix = append(matrix, row)
		j++
	}

	res := 0
	for i := 0; i < len(matrix); i++ {
		matrix.energize(i, 0, "right")
		tmpRes := matrix.count()
		if tmpRes > res {
			res = tmpRes
		}
		matrix.reset()

		matrix.energize(i, len(matrix[0])-1, "left")
		tmpRes = matrix.count()
		if tmpRes > res {
			res = tmpRes
		}
		matrix.reset()

		matrix.energize(0, i, "down")
		tmpRes = matrix.count()
		if tmpRes > res {
			res = tmpRes
		}
		matrix.reset()

		matrix.energize(len(matrix)-1, i, "up")
		tmpRes = matrix.count()
		if tmpRes > res {
			res = tmpRes
		}
		matrix.reset()
	}

	return res
}

func (matrix *Matrix) energize(x, y int, direction string) {

	currentBeams := make([]*Tile, 0)
	(*matrix)[x][y].Energized = true
	(*matrix)[x][y].Direction = append((*matrix)[x][y].Direction, direction)
	currentBeams = append(currentBeams, &(*matrix)[x][y])
	for {
		if len(currentBeams) == 0 {
			break
		}
		newBeams := make([]*Tile, 0)
		for _, beam := range currentBeams {
			switch beam.TileType {
			case '|':
				for _, direction := range beam.Direction {
					if direction == "right" || direction == "left" {
						if beam.X-1 >= 0 && !slices.Contains((*matrix)[beam.X-1][beam.Y].Energizedby, beam) {
							if slices.Contains(newBeams, &(*matrix)[beam.X-1][beam.Y]) {
								(*matrix)[beam.X-1][beam.Y].Direction = append((*matrix)[beam.X-1][beam.Y].Direction, "up")
							} else {
								(*matrix)[beam.X-1][beam.Y].Energized = true
								(*matrix)[beam.X-1][beam.Y].Direction = []string{"up"}
								newBeams = append(newBeams, &(*matrix)[beam.X-1][beam.Y])
							}
							(*matrix)[beam.X-1][beam.Y].Energizedby = append((*matrix)[beam.X-1][beam.Y].Energizedby, beam)
						}
						if beam.X+1 < len((*matrix)) && !slices.Contains((*matrix)[beam.X+1][beam.Y].Energizedby, beam) {
							if slices.Contains(newBeams, &(*matrix)[beam.X+1][beam.Y]) {
								(*matrix)[beam.X+1][beam.Y].Direction = append((*matrix)[beam.X+1][beam.Y].Direction, "down")
							} else {
								(*matrix)[beam.X+1][beam.Y].Energized = true
								(*matrix)[beam.X+1][beam.Y].Direction = []string{"down"}
								newBeams = append(newBeams, &(*matrix)[beam.X+1][beam.Y])
							}
							(*matrix)[beam.X+1][beam.Y].Energizedby = append((*matrix)[beam.X+1][beam.Y].Energizedby, beam)
						}
					} else if direction == "up" {
						if beam.X-1 >= 0 && !slices.Contains((*matrix)[beam.X-1][beam.Y].Energizedby, beam) {
							if slices.Contains(newBeams, &(*matrix)[beam.X-1][beam.Y]) {
								(*matrix)[beam.X-1][beam.Y].Direction = append((*matrix)[beam.X-1][beam.Y].Direction, "up")
							} else {
								(*matrix)[beam.X-1][beam.Y].Energized = true
								(*matrix)[beam.X-1][beam.Y].Direction = []string{"up"}
								newBeams = append(newBeams, &(*matrix)[beam.X-1][beam.Y])
							}
							(*matrix)[beam.X-1][beam.Y].Energizedby = append((*matrix)[beam.X-1][beam.Y].Energizedby, beam)
						}
					} else if direction == "down" {
						if beam.X+1 < len((*matrix)) && !slices.Contains((*matrix)[beam.X+1][beam.Y].Energizedby, beam) {
							if slices.Contains(newBeams, &(*matrix)[beam.X+1][beam.Y]) {
								(*matrix)[beam.X+1][beam.Y].Direction = append((*matrix)[beam.X+1][beam.Y].Direction, "down")
							} else {
								(*matrix)[beam.X+1][beam.Y].Energized = true
								(*matrix)[beam.X+1][beam.Y].Direction = []string{"down"}
								newBeams = append(newBeams, &(*matrix)[beam.X+1][beam.Y])
							}
							(*matrix)[beam.X+1][beam.Y].Energizedby = append((*matrix)[beam.X+1][beam.Y].Energizedby, beam)
						}
					}
				}
			case '-':
				for _, direction := range beam.Direction {
					if direction == "up" || direction == "down" {
						if beam.Y-1 >= 0 && !slices.Contains((*matrix)[beam.X][beam.Y-1].Energizedby, beam) {
							if slices.Contains(newBeams, &(*matrix)[beam.X][beam.Y-1]) {
								(*matrix)[beam.X][beam.Y-1].Direction = append((*matrix)[beam.X][beam.Y-1].Direction, "left")
							} else {
								(*matrix)[beam.X][beam.Y-1].Energized = true
								(*matrix)[beam.X][beam.Y-1].Direction = []string{"left"}
								newBeams = append(newBeams, &(*matrix)[beam.X][beam.Y-1])
							}
							(*matrix)[beam.X][beam.Y-1].Energizedby = append((*matrix)[beam.X][beam.Y-1].Energizedby, beam)
						}
						if beam.Y+1 < len((*matrix)[0]) && !slices.Contains((*matrix)[beam.X][beam.Y+1].Energizedby, beam) {
							if slices.Contains(newBeams, &(*matrix)[beam.X][beam.Y+1]) {
								(*matrix)[beam.X][beam.Y+1].Direction = append((*matrix)[beam.X][beam.Y+1].Direction, "right")
							} else {
								(*matrix)[beam.X][beam.Y+1].Energized = true
								(*matrix)[beam.X][beam.Y+1].Direction = []string{"right"}
								newBeams = append(newBeams, &(*matrix)[beam.X][beam.Y+1])
							}
							(*matrix)[beam.X][beam.Y+1].Energizedby = append((*matrix)[beam.X][beam.Y+1].Energizedby, beam)
						}
					} else if direction == "right" {
						if beam.Y+1 < len((*matrix)[0]) && !slices.Contains((*matrix)[beam.X][beam.Y+1].Energizedby, beam) {
							if slices.Contains(newBeams, &(*matrix)[beam.X][beam.Y+1]) {
								(*matrix)[beam.X][beam.Y+1].Direction = append((*matrix)[beam.X][beam.Y+1].Direction, "right")
							} else {
								(*matrix)[beam.X][beam.Y+1].Energized = true
								(*matrix)[beam.X][beam.Y+1].Direction = []string{"right"}
								newBeams = append(newBeams, &(*matrix)[beam.X][beam.Y+1])
							}
							(*matrix)[beam.X][beam.Y+1].Energizedby = append((*matrix)[beam.X][beam.Y+1].Energizedby, beam)
						}
					} else if direction == "left" {
						if beam.Y-1 >= 0 && !slices.Contains((*matrix)[beam.X][beam.Y-1].Energizedby, beam) {
							if slices.Contains(newBeams, &(*matrix)[beam.X][beam.Y-1]) {
								(*matrix)[beam.X][beam.Y-1].Direction = append((*matrix)[beam.X][beam.Y-1].Direction, "left")
							} else {
								(*matrix)[beam.X][beam.Y-1].Energized = true
								(*matrix)[beam.X][beam.Y-1].Direction = []string{"left"}
								newBeams = append(newBeams, &(*matrix)[beam.X][beam.Y-1])
							}
							(*matrix)[beam.X][beam.Y-1].Energizedby = append((*matrix)[beam.X][beam.Y-1].Energizedby, beam)
						}
					}
				}
			case '\\':
				for _, direction := range beam.Direction {
					if direction == "right" {
						if beam.X+1 < len((*matrix)) && !slices.Contains((*matrix)[beam.X+1][beam.Y].Energizedby, beam) {
							if slices.Contains(newBeams, &(*matrix)[beam.X+1][beam.Y]) {
								(*matrix)[beam.X+1][beam.Y].Direction = append((*matrix)[beam.X+1][beam.Y].Direction, "down")
							} else {
								(*matrix)[beam.X+1][beam.Y].Energized = true
								(*matrix)[beam.X+1][beam.Y].Direction = []string{"down"}
								newBeams = append(newBeams, &(*matrix)[beam.X+1][beam.Y])
							}
							(*matrix)[beam.X+1][beam.Y].Energizedby = append((*matrix)[beam.X+1][beam.Y].Energizedby, beam)
						}
					} else if direction == "left" {
						if beam.X-1 >= 0 && !slices.Contains((*matrix)[beam.X-1][beam.Y].Energizedby, beam) {
							if slices.Contains(newBeams, &(*matrix)[beam.X-1][beam.Y]) {
								(*matrix)[beam.X-1][beam.Y].Direction = append((*matrix)[beam.X-1][beam.Y].Direction, "up")
							} else {
								(*matrix)[beam.X-1][beam.Y].Energized = true
								(*matrix)[beam.X-1][beam.Y].Direction = []string{"up"}
								newBeams = append(newBeams, &(*matrix)[beam.X-1][beam.Y])
							}
							(*matrix)[beam.X-1][beam.Y].Energizedby = append((*matrix)[beam.X-1][beam.Y].Energizedby, beam)
						}
					} else if direction == "down" {
						if beam.Y+1 < len((*matrix)[0]) && !slices.Contains((*matrix)[beam.X][beam.Y+1].Energizedby, beam) {
							if slices.Contains(newBeams, &(*matrix)[beam.X][beam.Y+1]) {
								(*matrix)[beam.X][beam.Y+1].Direction = append((*matrix)[beam.X][beam.Y+1].Direction, "right")
							} else {
								(*matrix)[beam.X][beam.Y+1].Energized = true
								(*matrix)[beam.X][beam.Y+1].Direction = []string{"right"}
								newBeams = append(newBeams, &(*matrix)[beam.X][beam.Y+1])
							}
							(*matrix)[beam.X][beam.Y+1].Energizedby = append((*matrix)[beam.X][beam.Y+1].Energizedby, beam)
						}
					} else if direction == "up" {
						if beam.Y-1 >= 0 && !slices.Contains((*matrix)[beam.X][beam.Y-1].Energizedby, beam) {
							if slices.Contains(newBeams, &(*matrix)[beam.X][beam.Y-1]) {
								(*matrix)[beam.X][beam.Y-1].Direction = append((*matrix)[beam.X][beam.Y-1].Direction, "left")
							} else {
								(*matrix)[beam.X][beam.Y-1].Energized = true
								(*matrix)[beam.X][beam.Y-1].Direction = []string{"left"}
								newBeams = append(newBeams, &(*matrix)[beam.X][beam.Y-1])
							}
							(*matrix)[beam.X][beam.Y-1].Energizedby = append((*matrix)[beam.X][beam.Y-1].Energizedby, beam)
						}
					}
				}
			case '/':
				for _, direction := range beam.Direction {
					if direction == "right" {
						if beam.X-1 >= 0 && !slices.Contains((*matrix)[beam.X-1][beam.Y].Energizedby, beam) {
							if slices.Contains(newBeams, &(*matrix)[beam.X-1][beam.Y]) {
								(*matrix)[beam.X-1][beam.Y].Direction = append((*matrix)[beam.X-1][beam.Y].Direction, "up")
							} else {
								(*matrix)[beam.X-1][beam.Y].Energized = true
								(*matrix)[beam.X-1][beam.Y].Direction = []string{"up"}
								newBeams = append(newBeams, &(*matrix)[beam.X-1][beam.Y])
							}
							(*matrix)[beam.X-1][beam.Y].Energizedby = append((*matrix)[beam.X-1][beam.Y].Energizedby, beam)
						}
					} else if direction == "left" {
						if beam.X+1 < len((*matrix)) && !slices.Contains((*matrix)[beam.X+1][beam.Y].Energizedby, beam) {
							if slices.Contains(newBeams, &(*matrix)[beam.X+1][beam.Y]) {
								(*matrix)[beam.X+1][beam.Y].Direction = append((*matrix)[beam.X+1][beam.Y].Direction, "down")
							} else {
								(*matrix)[beam.X+1][beam.Y].Energized = true
								(*matrix)[beam.X+1][beam.Y].Direction = []string{"down"}
								newBeams = append(newBeams, &(*matrix)[beam.X+1][beam.Y])
							}
							(*matrix)[beam.X+1][beam.Y].Energizedby = append((*matrix)[beam.X+1][beam.Y].Energizedby, beam)
						}
					} else if direction == "up" {
						if beam.Y+1 < len((*matrix)[0]) && !slices.Contains((*matrix)[beam.X][beam.Y+1].Energizedby, beam) {
							if slices.Contains(newBeams, &(*matrix)[beam.X][beam.Y+1]) {
								(*matrix)[beam.X][beam.Y+1].Direction = append((*matrix)[beam.X][beam.Y+1].Direction, "right")
							} else {
								(*matrix)[beam.X][beam.Y+1].Energized = true
								(*matrix)[beam.X][beam.Y+1].Direction = []string{"right"}
								newBeams = append(newBeams, &(*matrix)[beam.X][beam.Y+1])
							}
							(*matrix)[beam.X][beam.Y+1].Energizedby = append((*matrix)[beam.X][beam.Y+1].Energizedby, beam)
						}
					} else if direction == "down" {
						if beam.Y-1 >= 0 && !slices.Contains((*matrix)[beam.X][beam.Y-1].Energizedby, beam) {
							if slices.Contains(newBeams, &(*matrix)[beam.X][beam.Y-1]) {
								(*matrix)[beam.X][beam.Y-1].Direction = append((*matrix)[beam.X][beam.Y-1].Direction, "left")
							} else {
								(*matrix)[beam.X][beam.Y-1].Energized = true
								(*matrix)[beam.X][beam.Y-1].Direction = []string{"left"}
								newBeams = append(newBeams, &(*matrix)[beam.X][beam.Y-1])
							}
							(*matrix)[beam.X][beam.Y-1].Energizedby = append((*matrix)[beam.X][beam.Y-1].Energizedby, beam)
						}
					}
				}
			case '.':
				for _, direction := range beam.Direction {
					if direction == "right" {
						if beam.Y+1 < len((*matrix)[0]) && !slices.Contains((*matrix)[beam.X][beam.Y+1].Energizedby, beam) {
							if slices.Contains(newBeams, &(*matrix)[beam.X][beam.Y+1]) {
								(*matrix)[beam.X][beam.Y+1].Direction = append((*matrix)[beam.X][beam.Y+1].Direction, "right")
							} else {
								(*matrix)[beam.X][beam.Y+1].Energized = true
								(*matrix)[beam.X][beam.Y+1].Direction = []string{"right"}
								newBeams = append(newBeams, &(*matrix)[beam.X][beam.Y+1])
							}
							(*matrix)[beam.X][beam.Y+1].Energizedby = append((*matrix)[beam.X][beam.Y+1].Energizedby, beam)
						}
					} else if direction == "left" {
						if beam.Y-1 >= 0 && !slices.Contains((*matrix)[beam.X][beam.Y-1].Energizedby, beam) {
							if slices.Contains(newBeams, &(*matrix)[beam.X][beam.Y-1]) {
								(*matrix)[beam.X][beam.Y-1].Direction = append((*matrix)[beam.X][beam.Y-1].Direction, "left")
							} else {
								(*matrix)[beam.X][beam.Y-1].Energized = true
								(*matrix)[beam.X][beam.Y-1].Direction = []string{"left"}
								newBeams = append(newBeams, &(*matrix)[beam.X][beam.Y-1])
							}
							(*matrix)[beam.X][beam.Y-1].Energizedby = append((*matrix)[beam.X][beam.Y-1].Energizedby, beam)
						}
					} else if direction == "up" {
						if beam.X-1 >= 0 && !slices.Contains((*matrix)[beam.X-1][beam.Y].Energizedby, beam) {
							if slices.Contains(newBeams, &(*matrix)[beam.X-1][beam.Y]) {
								(*matrix)[beam.X-1][beam.Y].Direction = append((*matrix)[beam.X-1][beam.Y].Direction, "up")
							} else {
								(*matrix)[beam.X-1][beam.Y].Energized = true
								(*matrix)[beam.X-1][beam.Y].Direction = []string{"up"}
								newBeams = append(newBeams, &(*matrix)[beam.X-1][beam.Y])
							}
							(*matrix)[beam.X-1][beam.Y].Energizedby = append((*matrix)[beam.X-1][beam.Y].Energizedby, beam)
						}
					} else if direction == "down" {
						if beam.X+1 < len((*matrix)) && !slices.Contains((*matrix)[beam.X+1][beam.Y].Energizedby, beam) {
							if slices.Contains(newBeams, &(*matrix)[beam.X+1][beam.Y]) {
								(*matrix)[beam.X+1][beam.Y].Direction = append((*matrix)[beam.X+1][beam.Y].Direction, "down")
							} else {
								(*matrix)[beam.X+1][beam.Y].Energized = true
								(*matrix)[beam.X+1][beam.Y].Direction = []string{"down"}
								newBeams = append(newBeams, &(*matrix)[beam.X+1][beam.Y])
							}
							(*matrix)[beam.X+1][beam.Y].Energizedby = append((*matrix)[beam.X+1][beam.Y].Energizedby, beam)
						}
					}
				}
			}
		}
		currentBeams = newBeams
	}
}

func (matrix *Matrix) reset() {
	for i := range len(*matrix) {
		for j := range len((*matrix)[i]) {
			(*matrix)[i][j].Energized = false
			(*matrix)[i][j].Energizedby = nil
			(*matrix)[i][j].Direction = nil
		}
	}
}

func (matrix *Matrix) count() int {
	res := 0
	for _, r := range *matrix {
		for _, c := range r {
			if c.Energized {
				res++
			}
		}
	}
	return res
}
