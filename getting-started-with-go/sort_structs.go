package main

import "fmt"

type Person struct {
	id   int
	name string
}

func main() {
	ppl := []Person{{6, "Francesco"}, {1, "Andrea"}, {5, "Elisa"}, {2, "Beatrice"}, {3, "Carlo"}, {4, "Dino"}, {7, "Giorgio"}, {9, "Irene"}, {8, "Henry"}}

	for k := 1; k < len(ppl); k++ {
		x := ppl[k]
		j := k - 1
		for j >= 0 && ppl[j].id < x.id {
			ppl[j+1] = ppl[j]
			j--
		}
		ppl[j+1] = x
	}

	fmt.Println(ppl)
}
