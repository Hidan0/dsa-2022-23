package main

import (
	"fmt"
)

func main() {
	values := []int{9, 1, 3, 5, 2, 0, 8, 9, 6, 7, 7, 0}

	var min_i int = 0
	var max_i int = 0

	for i := 0; i < len(values); i++ {
		if values[i] < values[min_i] {
			min_i = i
			if min_i > max_i {
				max_i = min_i
			}
		} else if values[i] > values[max_i] {
			max_i = i
		} else {
			if max_i-min_i >= 1 {
				fmt.Println(values[min_i : max_i+1])
			}
			min_i = i
			max_i = i
		}
	}
	if max_i-min_i >= 1 {
		fmt.Println(values[min_i : max_i+1])
	}
}
