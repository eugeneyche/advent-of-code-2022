package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	north int = iota
	south
	west
	east
)

type ivec2 struct {
	x, y int
}

type sparseGrid[value any] map[ivec2]value

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	initialElves := sparseGrid[bool]{}
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		for x := 0; x < len(line); x++ {
			if line[x] == '#' {
				initialElves[ivec2{x, y}] = true
			}
		}
	}

	deltas := []ivec2{
		{-1, -1},
		{0, -1},
		{1, -1},
		{-1, 0},
		{1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
	}

	deltasToCheckByDir := map[int][]ivec2{
		north: {{-1, -1}, {0, -1}, {1, -1}},
		south: {{-1, 1}, {0, 1}, {1, 1}},
		west:  {{-1, -1}, {-1, 0}, {-1, 1}},
		east:  {{1, -1}, {1, 0}, {1, 1}},
	}

	getBounds := func(elves sparseGrid[bool]) (minX, maxX, minY, maxY int) {
		hasInitialValue := false
		for elf := range elves {
			if !hasInitialValue {
				minX = elf.x
				maxX = elf.x
				minY = elf.y
				maxY = elf.y
				hasInitialValue = true
				continue
			}
			if elf.x < minX {
				minX = elf.x
			}
			if elf.x > maxX {
				maxX = elf.x
			}
			if elf.y < minY {
				minY = elf.y
			}
			if elf.y > maxY {
				maxY = elf.y
			}
		}
		return
	}

	simulate := func(elves sparseGrid[bool], round int) int {
		// Returns the number of moves
		deltaByElf := sparseGrid[ivec2]{}
		for elf := range elves {
			hasElfNearby := false
			for _, delta := range deltas {
				tx, ty := elf.x+delta.x, elf.y+delta.y
				if elves[ivec2{tx, ty}] {
					hasElfNearby = true
					break
				}
			}

			if !hasElfNearby {
				continue
			}

			for d := 0; d < 4; d++ {
				dir := (round + d) % 4
				deltas := deltasToCheckByDir[dir]
				hasElf := false
				for _, delta := range deltas {
					tx, ty := elf.x+delta.x, elf.y+delta.y
					if elves[ivec2{tx, ty}] {
						hasElf = true
						break
					}
				}

				if !hasElf {
					deltaByElf[elf] = deltas[1]
					break
				}
			}
		}

		proposeCount := sparseGrid[int]{}

		for elf, delta := range deltaByElf {
			proposed := ivec2{elf.x + delta.x, elf.y + delta.y}
			if count, ok := proposeCount[proposed]; ok {
				proposeCount[proposed] = count + 1
			} else {
				proposeCount[proposed] = 1
			}
		}

		numMoves := 0
		for elf, delta := range deltaByElf {
			proposed := ivec2{elf.x + delta.x, elf.y + delta.y}
			if proposeCount[proposed] == 1 {
				numMoves++
				delete(elves, elf)
				elves[proposed] = true
			}
		}
		return numMoves
	}

	p1 := func() {
		elves := sparseGrid[bool]{}
		for elf := range initialElves {
			elves[elf] = true
		}
		for r := 0; r < 10; r++ {
			simulate(elves, r)
		}

		numElves := len(elves)
		minX, maxX, minY, maxY := getBounds(elves)
		fmt.Printf("%d\n", (maxX-minX+1)*(maxY-minY+1)-numElves)
	}

	p2 := func() {
		elves := sparseGrid[bool]{}
		for elf := range initialElves {
			elves[elf] = true
		}

		var r int
		for r = 0; true; r++ {
			if simulate(elves, r) == 0 {
				break
			}
		}

		fmt.Printf("%d\n", r+1)
	}

	p1()
	p2()
}
