package main

import (
	"errors"
	"fmt"
	"log"
)

type Tower struct {
	Height int
	disks  []rune
}

func NewTower(n int) (*Tower, error) {
	if n < 0 || n > 25 {
		return nil, errors.New("Do you want to wait the end of the world?")
	}
	return &Tower{0, make([]rune, n)}, nil
}

func (tower *Tower) Init() {
	tower.Height = len(tower.disks)
	for i := 0; i < tower.Height; i++ {
		tower.disks[i] = rune('A' + i)
	}
}

func (tower *Tower) Pop() (rune, error) {
	if tower.Height <= 0 {
		return ' ', errors.New("Empty tower")
	}
	ch := tower.disks[tower.Height-1]
	tower.Height--
	return ch, nil
}

func (tower *Tower) TakeFrom(other *Tower) {
	ch, err := other.Pop()
	if err != nil {
		log.Fatal(err)
	}

	tower.disks[tower.Height] = ch
	tower.Height++
}

func (tower *Tower) AsString() string {
	var out string
	for i := 0; i < tower.Height; i++ {
		out += string(tower.disks[i])
	}
	return out
}

type Hanoi struct {
	N      int
	towers [3]Tower
}

func (hanoi *Hanoi) Print() {
	fmt.Printf("%s, %s, %s\n", hanoi.towers[0].AsString(), hanoi.towers[1].AsString(), hanoi.towers[2].AsString())
}

func (hanoi *Hanoi) _solve(n int, from *Tower, aux *Tower, to *Tower) {
	if n == 1 {
		to.TakeFrom(from)
		hanoi.Print()
		return
	}

	hanoi._solve(n-1, from, to, aux)
	to.TakeFrom(from)
	hanoi.Print()
	hanoi._solve(n-1, aux, from, to)
}

func (hanoi *Hanoi) SolveAndPrint() {
	hanoi.Print()
	hanoi._solve(hanoi.N, &hanoi.towers[0], &hanoi.towers[1], &hanoi.towers[2])
}

func NewTowerOfHanoi(numberOfDisks int) (*Hanoi, error) {
	var hanoi Hanoi

	hanoi.N = numberOfDisks

	tower, err := NewTower(numberOfDisks)
	if err != nil {
		return nil, err
	}
	hanoi.towers[0] = *tower
	hanoi.towers[0].Init()

	tower, _ = NewTower(numberOfDisks)
	hanoi.towers[1] = *tower

	tower, _ = NewTower(numberOfDisks)
	hanoi.towers[2] = *tower

	return &hanoi, nil
}

func main() {
	hanoi, err := NewTowerOfHanoi(3)
	if err != nil {
		log.Fatal(err)
	}
	hanoi.SolveAndPrint()
}
