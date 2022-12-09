package main

import (
	"bufio"
	"fmt"
	"os"
)

type direction int

const (
	up direction = iota
	down
	left
	right
)

func runeToDirection(b rune) direction {
	switch b {
	case 'U':
		return up
	case 'D':
		return down
	case 'L':
		return left
	case 'R':
		return right
	}
	return up
}

type instruction struct {
	dir   direction
	steps int
}

type position struct {
	r int
	c int
}

func absInt(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func clampInt(i int, l int, h int) int {
	if i < l {
		return l
	}
	if i > h {
		return h
	}
	return i
}

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	instructions := []instruction{}

	var dirRune rune
	var steps int
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%c %d", &dirRune, &steps)
		instructions = append(instructions, instruction{runeToDirection(dirRune), steps})
	}

	p1 := func() {
		head := position{0, 0}
		tail := position{0, 0}
		tailVisited := map[position]bool{}
		tailVisited[tail] = true

		for _, instr := range instructions {
			for i := 0; i < instr.steps; i++ {
				switch instr.dir {
				case up:
					head.r -= 1
				case down:
					head.r += 1
				case left:
					head.c -= 1
				case right:
					head.c += 1
				}

				// Fix tail pos
				dr := head.r - tail.r
				dc := head.c - tail.c

				if absInt(dr) >= 2 || absInt(dc) >= 2 {
					tail.r += clampInt(dr, -1, 1)
					tail.c += clampInt(dc, -1, 1)
				}
				tailVisited[tail] = true
			}
		}
		fmt.Printf("%d\n", len(tailVisited))
	}

	p2 := func() {
		segments := [10]position{}
		tailVisited := map[position]bool{}
		tailVisited[segments[9]] = true

		for _, instr := range instructions {
			for i := 0; i < instr.steps; i++ {
				switch instr.dir {
				case up:
					segments[0].r -= 1
				case down:
					segments[0].r += 1
				case left:
					segments[0].c -= 1
				case right:
					segments[0].c += 1
				}

				for j := 1; j < 10; j++ {
					head := &segments[j-1]
					tail := &segments[j]

					// Fix tail pos
					dr := head.r - tail.r
					dc := head.c - tail.c

					if absInt(dr) >= 2 || absInt(dc) >= 2 {
						tail.r += clampInt(dr, -1, 1)
						tail.c += clampInt(dc, -1, 1)
					}
				}

				tailVisited[segments[9]] = true
			}
		}
		fmt.Printf("%d\n", len(tailVisited))
	}

	_ = p1
	p2()
}
