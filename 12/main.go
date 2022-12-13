package main

import (
	"bufio"
	"fmt"
	"os"
)

func enqueue[T any](queue []T, element T) []T {
	queue = append(queue, element)
	return queue
}

func dequeue[T any](queue []T) (T, []T) {
	element := queue[0]
	if len(queue) == 1 {
		return element, []T{}
	}
	return element, queue[1:]
}

type vec2 struct {
	row int
	col int
}

type node struct {
	pos      vec2
	numSteps int
}

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	grid := [][]int{}
	var start, end vec2

	for rowIdx := 0; scanner.Scan(); rowIdx++ {
		row := []int{}
		for colIdx, char := range []rune(scanner.Text()) {
			height := int(char - 'a')
			if char == 'S' {
				start = vec2{rowIdx, colIdx}
				height = 0
			} else if char == 'E' {
				end = vec2{rowIdx, colIdx}
				height = 25
			}
			row = append(row, height)
		}
		grid = append(grid, row)
	}
	numGridRows := len(grid)
	numGridCols := len(grid[0])

	p1 := func() {
		queue := []node{}
		explored := map[vec2]bool{}
		queue = enqueue(queue, node{start, 0})
		explored[start] = true

		for len(queue) > 0 {
			var currNode node
			currNode, queue = dequeue(queue)
			currPos := currNode.pos
			currHeight := grid[currPos.row][currPos.col]

			if currPos == end {
				fmt.Printf("%v\n", currNode.numSteps)
				return
			}

			for _, offset := range []vec2{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				nextPos := vec2{currPos.row + offset.row, currPos.col + offset.col}
				if nextPos.row < 0 || nextPos.row >= numGridRows || nextPos.col < 0 || nextPos.col >= numGridCols {
					continue
				}
				if _, ok := explored[nextPos]; ok {
					continue
				}
				nextHeight := grid[nextPos.row][nextPos.col]
				heightDiff := nextHeight - currHeight
				if heightDiff <= 1 {
					queue = enqueue(queue, node{nextPos, currNode.numSteps + 1})
					explored[nextPos] = true
				}
			}
		}
	}

	p2 := func() {
		queue := []node{}
		explored := map[vec2]bool{}
		for r := 0; r < numGridRows; r++ {
			for c := 0; c < numGridCols; c++ {
				if grid[r][c] == 0 {
					queue = enqueue(queue, node{vec2{r, c}, 0})
					explored[start] = true
				}
			}
		}

		for len(queue) > 0 {
			var currNode node
			currNode, queue = dequeue(queue)
			currPos := currNode.pos
			currHeight := grid[currPos.row][currPos.col]

			if currPos == end {
				fmt.Printf("%v\n", currNode.numSteps)
				return
			}

			for _, offset := range []vec2{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				nextPos := vec2{currPos.row + offset.row, currPos.col + offset.col}
				if nextPos.row < 0 || nextPos.row >= numGridRows || nextPos.col < 0 || nextPos.col >= numGridCols {
					continue
				}
				if _, ok := explored[nextPos]; ok {
					continue
				}
				nextHeight := grid[nextPos.row][nextPos.col]
				heightDiff := nextHeight - currHeight
				if heightDiff <= 1 {
					queue = enqueue(queue, node{nextPos, currNode.numSteps + 1})
					explored[nextPos] = true
				}
			}
		}
	}

	_ = p1
	p2()
}
