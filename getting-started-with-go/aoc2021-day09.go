package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
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

func basinContains(basin *[]Position, o Position) bool {
	for _, el := range *basin {
		if el == o {
			return true
		}
	}
	return false
}

func findBasinSize(hMap *[][]int, basin *[]Position, this Position) int {
	if (*hMap)[this.y][this.x] >= 9 {
		return 0
	}

	if !basinContains(basin, this) {
		*basin = append(*basin, this)
	}

	// TOP
	if this.y > 0 && !basinContains(basin, NewPos(this.x, this.y-1)) {
		findBasinSize(hMap, basin, NewPos(this.x, this.y-1))
	}

	// RIGHT
	if this.x < len((*hMap)[0])-1 && !basinContains(basin, NewPos(this.x+1, this.y)) {
		findBasinSize(hMap, basin, NewPos(this.x+1, this.y))
	}
	// BOTTOM
	if this.y < len(*hMap)-1 && !basinContains(basin, NewPos(this.x, this.y+1)) {
		findBasinSize(hMap, basin, NewPos(this.x, this.y+1))
	}

	// LEFT
	if this.x > 0 && !basinContains(basin, NewPos(this.x-1, this.y)) {
		findBasinSize(hMap, basin, NewPos(this.x-1, this.y))
	}

	return len(*basin)
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
	pos := make([]Position, 0)
	fmt.Println(findBasinSize(&h_map, &pos, NewPos(0, 0)))
	fmt.Println(pos)

	pos = make([]Position, 0)
	fmt.Println(findBasinSize(&h_map, &pos, NewPos(2, 2)))
	fmt.Println(pos)
}
