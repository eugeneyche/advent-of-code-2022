package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ivec2 struct {
	r int
	c int
}

type path struct {
	points []ivec2
}

func iSorted(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	paths := []path{}

	var r, c int
	for scanner.Scan() {
		pointStrs := strings.Split(scanner.Text(), " -> ")
		points := []ivec2{}
		for _, pointStr := range pointStrs {
			fmt.Sscanf(pointStr, "%d,%d", &c, &r)
			points = append(points, ivec2{r, c})
		}
		paths = append(paths, path{points})
	}

	grid := map[ivec2]bool{}
	for _, p := range paths {
		for i := 0; i < len(p.points)-1; i++ {
			p0 := p.points[i]
			p1 := p.points[i+1]
			if p0.r != p1.r {
				c := p0.c
				minR, maxR := iSorted(p0.r, p1.r)
				for r := minR; r <= maxR; r++ {
					grid[ivec2{r, c}] = true
				}
			} else /* aPoint.c != bPoint.c */ {
				r := p0.r
				minC, maxC := iSorted(p0.c, p1.c)
				for c := minC; c <= maxC; c++ {
					grid[ivec2{r, c}] = true
				}
			}
		}
	}

	p1 := func() {
		maxR := 0
		for point := range grid {
			if maxR < point.r {
				maxR = point.r
			}
		}

		numSand := 0
		sandEye := ivec2{0, 500}
		for true {
			sand := sandEye
			for sand.r <= maxR {
				below := ivec2{sand.r + 1, sand.c}
				belowLeft := ivec2{sand.r + 1, sand.c - 1}
				belowRight := ivec2{sand.r + 1, sand.c + 1}
				if _, ok := grid[below]; !ok {
					sand = below
				} else if _, ok := grid[belowLeft]; !ok {
					sand = belowLeft
				} else if _, ok := grid[belowRight]; !ok {
					sand = belowRight
				} else {
					break
				}
			}

			grid[sand] = true
			if sand.r > maxR {
				break
			}

			numSand++
		}
		fmt.Printf("%d\n", numSand)
	}

	p2 := func() {
		numSand := 0
		sandEye := ivec2{0, 500}

		maxR := 0
		for pt := range grid {
			if maxR < pt.r {
				maxR = pt.r
			}
		}

		for true {
			if _, ok := grid[sandEye]; ok {
				break
			}

			sand := sandEye
			// Stop when sand.r is maxR + 1 since there is a wall at maxR + 2
			for sand.r < maxR+1 {
				below := ivec2{sand.r + 1, sand.c}
				belowLeft := ivec2{sand.r + 1, sand.c - 1}
				belowRight := ivec2{sand.r + 1, sand.c + 1}
				if _, ok := grid[below]; !ok {
					sand = below
				} else if _, ok := grid[belowLeft]; !ok {
					sand = belowLeft
				} else if _, ok := grid[belowRight]; !ok {
					sand = belowRight
				} else {
					break
				}
			}

			grid[sand] = true
			numSand++
		}
		fmt.Printf("%d\n", numSand)
	}

	p1()
	p2()
}
