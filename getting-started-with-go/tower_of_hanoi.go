package main

import "fmt"

func main() {
	hanoi(3, 0, 1, 2)
}

func hanoi(n int, from int, tmp int, to int) {
	if n == 1 {
		fmt.Printf("%d -> %d\n", from, to)
		return
	}
	hanoi(n-1, from, to, tmp)
	fmt.Printf("%d -> %d\n", from, to)
	hanoi(n-1, tmp, from, to)
}
