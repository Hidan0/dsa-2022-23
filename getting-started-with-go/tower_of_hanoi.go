package main

import "fmt"

func main() {
	fmt.Printf("The number of moves are %d\n", hanoi(3, 0, 1, 2, true))
	// fmt.Printf("The end of the world in %d days?\n", hanoi(64, 0, 1, 2, false)) // Yes it is... 2^64-1
}

func hanoi(n int, from int, tmp int, to int, print bool) int {
	if n == 1 {
		if print {
			fmt.Printf("%d -> %d\n", from, to)
		}
		return 1
	}
	hanoi(n-1, from, to, tmp, print)
	if print {
		fmt.Printf("%d -> %d\n", from, to)
	}
	return 2*hanoi(n-1, tmp, from, to, print) + 1
}
