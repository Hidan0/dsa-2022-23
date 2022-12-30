package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseInput(r io.Reader) ([][]int, error) {
	scanner := bufio.NewScanner(r)

	scanner.Split(bufio.ScanLines)

	var out [][]int
	i := 0
	for scanner.Scan() {
		line := scanner.Text()

		out = append(out, make([]int, 0))
		lineScanner := bufio.NewScanner(strings.NewReader(line))
		lineScanner.Split(bufio.ScanRunes)
		for lineScanner.Scan() {
			h, err := strconv.Atoi(lineScanner.Text())
			if err != nil {
				return nil, err
			}
			if h < 0 || h > 9 {
				return nil, errors.New("Invalid height location")
			}
			out[i] = append(out[i], h)
		}
		i++
	}

	return out, nil
}

func findRiskLevel(hMap *[][]int) int {
	var out int

	for i := 0; i < len(*hMap); i++ {
		for j := 0; j < len((*hMap)[0]); j++ {
			lowest := true
			target := (*hMap)[i][j]

			if i > 0 && target >= (*hMap)[i-1][j] {
				lowest = false
			}

			if j < len((*hMap)[0])-1 && target >= (*hMap)[i][j+1] {
				lowest = false
			}

			if i < len(*hMap)-1 && target >= (*hMap)[i+1][j] {
				lowest = false
			}

			if j > 0 && target >= (*hMap)[i][j-1] {
				lowest = false
			}

			if lowest {
				out += 1 + (*hMap)[i][j]
			}

		}
	}
	return out
}

type Position struct {
	x int
	y int
}

func NewPos(x int, y int) Position {
	return Position{x, y}
}

type Basin []Position

func NewBasin() Basin {
	return make([]Position, 0)
}

func (basin *Basin) contains(o Position) bool {
	for _, el := range *basin {
		if el == o {
			return true
		}
	}
	return false
}

func findBasinSize(hMap *[][]int, basin *Basin, this Position) int {
	if (*hMap)[this.y][this.x] >= 9 {
		return 0
	}

	*basin = append(*basin, this)

	// TOP
	if this.y > 0 && !basin.contains(NewPos(this.x, this.y-1)) {
		findBasinSize(hMap, basin, NewPos(this.x, this.y-1))
	}

	// RIGHT
	if this.x < len((*hMap)[0])-1 && !basin.contains(NewPos(this.x+1, this.y)) {
		findBasinSize(hMap, basin, NewPos(this.x+1, this.y))
	}
	// BOTTOM
	if this.y < len(*hMap)-1 && !basin.contains(NewPos(this.x, this.y+1)) {
		findBasinSize(hMap, basin, NewPos(this.x, this.y+1))
	}

	// LEFT
	if this.x > 0 && !basin.contains(NewPos(this.x-1, this.y)) {
		findBasinSize(hMap, basin, NewPos(this.x-1, this.y))
	}

	return len(*basin)
}

func _isInOneBasin(basins *[]Basin, position Position) bool {
	for _, basin := range *basins {
		if basin.contains(position) {
			return true
		}
	}

	return false
}

func threeLargestBasins(hMap *[][]int) int {
	sizes := make([]int, 0)
	basins := make([]Basin, 0)
	for i := 0; i < len(*hMap); i++ {
		for j := 0; j < len((*hMap)[0]); j++ {
			pos := NewPos(j, i)
			if !_isInOneBasin(&basins, pos) {
				basin := NewBasin()
				sizes = append(sizes, findBasinSize(hMap, &basin, pos))
				basins = append(basins, basin)
			}
		}
	}

	sort.Ints(sizes[:])
	return sizes[len(sizes)-1] * sizes[len(sizes)-2] * sizes[len(sizes)-3]
}

func main() {
	test_part1()
	f, err := os.Open("input09.txt")
	if err != nil {
		log.Fatal(err)
	}

	hMap, err := parseInput(bufio.NewReader(f))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Risk Level:", findRiskLevel(&hMap))

	test_part2()
	fmt.Println("Three largest basins:", threeLargestBasins(&hMap))
}

var INPUT = `2199943210
3987894921
9856789892
8767896789
9899965678`

func test_part1() {
	h_map, err := parseInput(strings.NewReader(INPUT))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("TEST:", findRiskLevel(&h_map))
}

func test_part2() {
	h_map, err := parseInput(strings.NewReader(INPUT))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("TEST:", threeLargestBasins(&h_map))
}
