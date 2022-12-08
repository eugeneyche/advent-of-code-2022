package main

import (
	"fmt"
	"os"
	"strings"
)

type coord struct {
	r int
	c int
}

func main() {
	content, _ := os.ReadFile(os.Args[1])
	lines := strings.Split(string(content), "\n")
	grid := [][]int{}

	for _, line := range lines {
		row := []int{}
		for _, char := range line {
			num := int(char - '0')
			row = append(row, num)
		}
		grid = append(grid, row)
	}

	numRows := len(grid)
	numCols := len(grid[0])

	p1 := func() {
		isTreeSeen := map[coord]bool{}
		for c := 0; c < numCols; c++ {
			// Up
			tallestSoFar := -1
			for r := numRows - 1; r >= 0; r -= 1 {
				tree := grid[r][c]
				if tree > tallestSoFar {
					tallestSoFar = tree
					isTreeSeen[coord{r, c}] = true
				}
			}

			// Down
			tallestSoFar = -1
			for r := 0; r < numRows; r++ {
				tree := grid[r][c]
				if tree > tallestSoFar {
					tallestSoFar = tree
					isTreeSeen[coord{r, c}] = true
				}
			}

		}
		for r := 0; r < numRows; r++ {
			// Right
			tallestSoFar := -1
			for c := 0; c < numCols; c++ {
				tree := grid[r][c]
				if tree > tallestSoFar {
					tallestSoFar = tree
					isTreeSeen[coord{r, c}] = true
				}
			}
			// Left
			tallestSoFar = -1
			for c := numCols - 1; c >= 0; c -= 1 {
				tree := grid[r][c]
				if tree > tallestSoFar {
					tallestSoFar = tree
					isTreeSeen[coord{r, c}] = true
				}
			}
		}

		fmt.Printf("%d\n", len(isTreeSeen))
	}
	p2 := func() {
		maxScenicScore := 0
		for r := 0; r < numRows; r++ {
			for c := 0; c < numCols; c++ {
				tree := grid[r][c]

				// Up
				upViewDist := 0
				for rr := r - 1; rr >= 0; rr -= 1 {
					upViewDist++
					if grid[rr][c] >= tree {
						break
					}
				}

				// Down
				downViewDist := 0
				for rr := r + 1; rr < numRows; rr++ {
					downViewDist++
					if grid[rr][c] >= tree {
						break
					}
				}

				// Left
				leftViewDist := 0
				for cc := c - 1; cc >= 0; cc -= 1 {
					leftViewDist++
					if grid[r][cc] >= tree {
						break
					}
				}

				// Right
				rightViewDist := 0
				for cc := c + 1; cc < numCols; cc++ {
					rightViewDist++
					if grid[r][cc] >= tree {
						break
					}
				}

				scenicScore := upViewDist * downViewDist * leftViewDist * rightViewDist

				if scenicScore > maxScenicScore {
					maxScenicScore = scenicScore
				}
			}
		}
		fmt.Printf("%d\n", maxScenicScore)
	}

	_ = p1
	p2()
}
