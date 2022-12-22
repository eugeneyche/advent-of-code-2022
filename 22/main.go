package main

import (
	"bufio"
	"fmt"
	"os"
)

type state struct {
	x   int
	y   int
	dir int
}

const (
	right int = iota
	down
	left
	up
)

func (s state) forwardDelta() (int, int) {
	switch s.dir {
	case right:
		return 1, 0
	case down:
		return 0, 1
	case left:
		return -1, 0
	case up:
		return 0, -1
	}
	return 0, 0
}

func (s state) turnRight() state {
	return state{s.x, s.y, (s.dir + 1) % 4}
}

func (s state) turnLeft() state {
	return state{s.x, s.y, (s.dir + 3) % 4}
}

type instructionOp int

const (
	forward instructionOp = iota
	turnRight
	turnLeft
)

type instruction struct {
	op     instructionOp
	amount int
}

func parseInstructions(s string) []instruction {
	instrs := []instruction{}
	numAcc := 0
	for _, c := range s {
		if c-'0' < 10 {
			numAcc = 10*numAcc + int(c-'0')
		} else if c == 'L' {
			instrs = append(
				instrs,
				instruction{forward, numAcc},
				instruction{turnLeft, 0},
			)
			numAcc = 0
		} else if c == 'R' {
			instrs = append(
				instrs,
				instruction{forward, numAcc},
				instruction{turnRight, 0},
			)
			numAcc = 0
		}
	}
	instrs = append(
		instrs,
		instruction{forward, numAcc},
	)
	return instrs
}

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := [][]byte{}
	gridWidth := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		if len(line) > gridWidth {
			gridWidth = len(line)
		}
		grid = append(grid, []byte(line))
	}

	minXAtY := make([]int, len(grid))
	maxXAtY := make([]int, len(grid))
	minYAtX := make([]int, gridWidth)
	maxYAtX := make([]int, gridWidth)

	for y := 0; y < len(grid); y++ {
		hasSeenMin := false
		for x := 0; x < len(grid[y]); x++ {
			cell := grid[y][x]
			if cell == ' ' {
				continue
			}
			if !hasSeenMin {
				hasSeenMin = true
				minXAtY[y] = x
			}
			maxXAtY[y] = x
		}
	}
	for x := 0; x < gridWidth; x++ {
		hasSeenMin := false
		for y := 0; y < len(grid); y++ {
			if x >= len(grid[y]) {
				continue
			}
			cell := grid[y][x]
			if cell == ' ' {
				continue
			}
			if !hasSeenMin {
				hasSeenMin = true
				minYAtX[x] = y
			}
			maxYAtX[x] = y
		}
	}

	scanner.Scan()
	instrs := parseInstructions(scanner.Text())

	upWrapAtX := map[int]state{}
	downWrapAtX := map[int]state{}
	leftWrapAtY := map[int]state{}
	rightWrapAtY := map[int]state{}

	traverse := func(initialState state, instructions []instruction) state {
		s := initialState
		for _, instr := range instrs {
			switch instr.op {
			case forward:
				for i := 0; i < instr.amount; i++ {
					dx, dy := s.forwardDelta()
					ns := state{s.x + dx, s.y + dy, s.dir}
					if dx != 0 && !(minXAtY[ns.y] <= ns.x && ns.x <= maxXAtY[ns.y]) {
						if dx > 0 {
							ns = rightWrapAtY[ns.y]
						} else {
							ns = leftWrapAtY[ns.y]
						}
					}
					if dy != 0 && !(minYAtX[ns.x] <= ns.y && ns.y <= maxYAtX[ns.x]) {
						if dy > 0 {
							ns = downWrapAtX[ns.x]
						} else {
							ns = upWrapAtX[ns.x]
						}
					}
					if grid[ns.y][ns.x] == '.' {
						s = ns
					}
				}
			case turnLeft:
				s = s.turnLeft()
			case turnRight:
				s = s.turnRight()
			}
		}
		return s
	}

	p1 := func() {
		for y := 0; y < len(grid); y++ {
			leftWrapAtY[y] = state{maxXAtY[y], y, left}
			rightWrapAtY[y] = state{minXAtY[y], y, right}
		}
		for x := 0; x < gridWidth; x++ {
			upWrapAtX[x] = state{x, maxYAtX[x], up}
			downWrapAtX[x] = state{x, minYAtX[x], down}
		}
		s := traverse(state{minXAtY[0], 0, right}, instrs)
		passwd := 1000*(s.y+1) + 4*(s.x+1) + s.dir
		fmt.Printf("%d\n", passwd)
	}

	p2 := func() {
		if os.Args[1] == "sample.txt" {
			for i := 0; i < 4; i++ {
				rightWrapAtY[4+i] = state{15 - i, 8, down}
				downWrapAtX[8+i] = state{3 - i, 7, up}
				upWrapAtX[4+i] = state{8, i, right}
			}
		} else {
			//  21
			//  3
			// 54
			// 6
			for i := 0; i < 50; i++ {
				// 1 <-> 3
				downWrapAtX[100+i] = state{99, 50 + i, left}
				rightWrapAtY[50+i] = state{100 + i, 49, up}

				// 1 <-> 4
				rightWrapAtY[i] = state{99, 149 - i, left}
				rightWrapAtY[100+i] = state{149, 49 - i, left}

				// 1 <-> 6
				upWrapAtX[100+i] = state{i, 199, up}
				downWrapAtX[i] = state{100 + i, 0, down}

				// 2 <-> 6
				upWrapAtX[50+i] = state{0, 150 + i, right}
				leftWrapAtY[150+i] = state{50 + i, 0, down}

				// 2 <-> 5
				leftWrapAtY[i] = state{0, 149 - i, right}
				leftWrapAtY[100+i] = state{50, 49 - i, right}

				// 3 <-> 5
				leftWrapAtY[50+i] = state{i, 100, down}
				upWrapAtX[i] = state{50, 50 + i, right}

				// 4 <-> 6
				rightWrapAtY[150+i] = state{50 + i, 149, up}
				downWrapAtX[50+i] = state{49, 150 + i, left}
			}
		}
		s := traverse(state{minXAtY[0], 0, 0}, instrs)
		passwd := 1000*(s.y+1) + 4*(s.x+1) + s.dir
		fmt.Printf("%d\n", passwd)
	}

	p1()
	p2()
}
