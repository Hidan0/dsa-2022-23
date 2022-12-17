package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

type Lanternfish int

func main() {
	testInput()
}

func parseInput(r io.Reader) (lanternfish []Lanternfish, err error) {
	scanner := bufio.NewScanner(r)

	lanternfish = make([]Lanternfish, 0)

	if scanner.Scan() {
		src := strings.Split(scanner.Text(), ",")
		for _, s_fish := range src {
			fish, err := strconv.Atoi(s_fish)
			if err != nil {
				return lanternfish, err
			}
			lanternfish = append(lanternfish, Lanternfish(fish))
		}
	} else {
		return lanternfish, fmt.Errorf("Can not scan the input")
	}

	return lanternfish, nil
}

func calculateFish(lanterfish []Lanternfish, days int, verbose ...bool) int {
	var print bool = false
	if len(verbose) == 1 {
		print = verbose[0]
	}

	if print {
		fmt.Printf("Initial state: ")
		fmt.Println(lanterfish)
	}

	for i := 1; i <= days; i++ {
		max := len(lanterfish)
		for j := 0; j < max; j++ {
			lanterfish[j]--
			if lanterfish[j] < 0 {
				lanterfish[j] = 6
				lanterfish = append(lanterfish, Lanternfish(8))
			}
		}
		if print {
			fmt.Printf("After %d days: ", i)
			fmt.Println(lanterfish)
		}
	}

	return len(lanterfish)
}

func testInput() {
	const TEST_SRC string = "3,4,3,1,2"

	fish, err := parseInput(strings.NewReader(TEST_SRC))
	if err != nil {
		log.Fatal(err)
	}
	total := calculateFish(fish, 18, true)
	fmt.Printf("Total of fish after %d days is %d\n", 18, total)
}
