package main

import (
	"bufio"
	"fmt"
	"os"
)

type ivec2 struct {
	x, y int
}

type block struct {
	points []ivec2
}

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	jetStream := scanner.Text()

	blocks := []block{
		{
			[]ivec2{
				{0, 0},
				{1, 0},
				{2, 0},
				{3, 0},
			},
		},
		{
			[]ivec2{
				{1, 0},
				{0, 1},
				{1, 1},
				{2, 1},
				{1, 2},
			},
		},
		{
			[]ivec2{
				{0, 0},
				{1, 0},
				{2, 0},
				{2, 1},
				{2, 2},
			},
		},
		{
			[]ivec2{
				{0, 0},
				{0, 1},
				{0, 2},
				{0, 3},
			},
		},
		{
			[]ivec2{
				{0, 0},
				{1, 0},
				{0, 1},
				{1, 1},
			},
		},
	}

	getHeightAfterDroppingNBlocks := func(numBlocks int) int {
		type memKey struct {
			blockIdx     int
			jetStreamIdx int
			gridSample   [5][7]bool
		}

		type state struct {
			nthBlock int
			height   int
		}

		grid := [10_000][7]bool{}
		jetStreamIdx := 0
		maxBlockHeight := 0
		mem := map[memKey]state{}

		heightAddedFromRepeats := 0

		for i := 0; i < numBlocks; i++ {
			if maxBlockHeight >= 5 {
				preDropState := memKey{i % 5, jetStreamIdx, [5][7]bool{}}
				for y := 0; y < 5; y++ {
					for x := 0; x < 7; x++ {
						preDropState.gridSample[y][x] = grid[maxBlockHeight-5+y][x]
					}
				}
				if lastSeen, ok := mem[preDropState]; ok {
					numBlocksToRepeat := i - lastSeen.nthBlock
					blocksLeft := numBlocks - i

					heightPerRepeat := maxBlockHeight - lastSeen.height
					numRepeats := blocksLeft / numBlocksToRepeat
					heightAddedFromRepeats += numRepeats * heightPerRepeat

					i += numRepeats * numBlocksToRepeat
				} else {
					mem[preDropState] = state{i, maxBlockHeight}
				}
			}

			block := &blocks[i%len(blocks)]
			isOpen := func(dx, dy int) bool {
				for _, p := range block.points {
					pX := p.x + dx
					pY := p.y + dy
					if pX < 0 || 7 <= pX || pY < 0 || grid[pY][pX] {
						return false
					}
				}
				return true
			}

			dx := 2
			dy := maxBlockHeight + 3

			for true {
				jet := jetStream[jetStreamIdx]
				jetStreamIdx = (jetStreamIdx + 1) % len(jetStream)
				if jet == '<' && isOpen(dx-1, dy) {
					dx--
				} else if jet == '>' && isOpen(dx+1, dy) {
					dx++
				}

				if isOpen(dx, dy-1) {
					dy--
				} else {
					break
				}
			}

			for _, p := range block.points {
				pX := p.x + dx
				pY := p.y + dy
				grid[pY][pX] = true
				if pY+1 > maxBlockHeight {
					maxBlockHeight = pY + 1
				}
			}
		}

		return maxBlockHeight + heightAddedFromRepeats
	}

	p1 := func() {
		fmt.Printf("%d\n", getHeightAfterDroppingNBlocks(2022))
	}

	p2 := func() {
		fmt.Printf("%d\n", getHeightAfterDroppingNBlocks(1_000_000_000_000))
	}

	p1()
	p2()
}
