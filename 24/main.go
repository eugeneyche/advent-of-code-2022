package main

import (
	"bufio"
	"fmt"
	"os"
)

type direction int

const (
	north direction = iota
	south
	east
	west
)

type ivec2 struct {
	x, y int
}

type queue[t any] []t

func enqueue[t any](q queue[t], v t) queue[t] {
	return append(q, v)
}

func dequeue[t any](q queue[t]) (t, queue[t]) {
	return q[0], q[1:]
}

type state struct {
	pos  ivec2
	step int
}

type stateKey struct {
	pos  ivec2
	wMod int
	hMod int
}

type blizzard struct {
	pos ivec2
	dir direction
}

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	grid := [][]byte{}
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}

	gridHeight := len(grid) - 2
	gridWidth := len(grid[0]) - 2

	blizzards := []blizzard{}

	for y := 0; y < gridHeight; y++ {
		for x := 0; x < gridWidth; x++ {
			switch grid[y+1][x+1] {
			case '^':
				blizzards = append(blizzards, blizzard{ivec2{x, y}, north})
			case 'v':
				blizzards = append(blizzards, blizzard{ivec2{x, y}, south})
			case '>':
				blizzards = append(blizzards, blizzard{ivec2{x, y}, east})
			case '<':
				blizzards = append(blizzards, blizzard{ivec2{x, y}, west})
			}
		}
	}

	entrance := ivec2{0, -1}
	exit := ivec2{gridWidth - 1, gridHeight}

	getStateKey := func(s state) stateKey {
		return stateKey{s.pos, s.step % gridWidth, s.step % gridHeight}
	}

	isInbounds := func(pos ivec2) bool {
		if pos == entrance || pos == exit {
			return true
		}
		return 0 <= pos.x && pos.x < gridWidth &&
			0 <= pos.y && pos.y < gridHeight
	}

	deltas := []ivec2{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
		{0, 0},
	}

	computeBlizzardCoverage := func(step int) map[ivec2]bool {
		hasBlizzard := map[ivec2]bool{}
		for _, b := range blizzards {
			switch b.dir {
			case north:
				hasBlizzard[ivec2{b.pos.x, (b.pos.y + gridHeight - (step % gridHeight)) % gridHeight}] = true
			case south:
				hasBlizzard[ivec2{b.pos.x, (b.pos.y + step) % gridHeight}] = true
			case east:
				hasBlizzard[ivec2{(b.pos.x + step) % gridWidth, b.pos.y}] = true
			case west:
				hasBlizzard[ivec2{(b.pos.x + gridWidth - (step % gridWidth)) % gridWidth, b.pos.y}] = true
			}
		}
		return hasBlizzard
	}

	computeMinSteps := func(start, end ivec2, startStep int) int {
		initS := state{start, startStep}
		q := queue[state]{initS}
		seen := map[stateKey]bool{getStateKey(initS): true}

		hasBlizzard := computeBlizzardCoverage(startStep)
		blizzardStep := startStep

		var s state
		var stepsToExit int
		for len(q) > 0 {
			s, q = dequeue(q)

			if s.pos == end {
				stepsToExit = s.step
				break
			}

			nextStep := s.step + 1
			if nextStep != blizzardStep {
				hasBlizzard = computeBlizzardCoverage(nextStep)
				blizzardStep = nextStep
			}

			for _, d := range deltas {
				nextS := state{ivec2{s.pos.x + d.x, s.pos.y + d.y}, nextStep}
				nextSKey := getStateKey(nextS)

				if seen[nextSKey] || !isInbounds(nextS.pos) || hasBlizzard[nextS.pos] {
					continue
				}

				q = enqueue(q, nextS)
				seen[nextSKey] = true
			}
		}
		return stepsToExit
	}

	p1 := func() {
		fmt.Printf("%d\n", computeMinSteps(entrance, exit, 0))
	}

	p2 := func() {
		toExit := computeMinSteps(entrance, exit, 0)
		andBack := computeMinSteps(exit, entrance, toExit)
		andBackToExit := computeMinSteps(entrance, exit, andBack)
		fmt.Printf("%d\n", andBackToExit)
	}

	p1()
	p2()
}
