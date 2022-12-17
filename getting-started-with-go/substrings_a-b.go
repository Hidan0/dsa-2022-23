package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "ccbaacbabbcbabb"
	count := 0
	for i, ch := range str {
		if ch == 'a' {
			for _, ch := range str[i:] {
				if ch == 'b' {
					count++
				}
			}
		}
	}
	fmt.Println("The number of a-b subsrings is: " + strconv.Itoa(count))
}
