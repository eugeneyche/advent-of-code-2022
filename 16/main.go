package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type valveInfo struct {
	id              string
	flowRate        int
	connectedValves []string
}

type bitset uint64

func (bs bitset) get(i int) bool {
	return (bs & (1 << i)) > 0
}

func (bs bitset) set(i int) bitset {
	return bs | (1 << i)
}

func intMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func intMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)

	valveById := map[string]valveInfo{}

	var id string
	var flowRate int
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Sscanf(line, "Valve %s has flow rate=%d; tunnels lead to valves ", &id, &flowRate)
		remainingIdx := intMax(
			strings.Index(line, "valves ")+7,
			strings.Index(line, "valve ")+6,
		)
		connectedValvesIds := strings.Split(line[remainingIdx:], ", ")

		valveById[id] = valveInfo{
			id,
			flowRate,
			connectedValvesIds,
		}
	}

	numValves := len(valveById)
	indexById := map[string]int{}
	flowByIndex := make([]int, numValves)
	nextIndex := 0
	for id, valve := range valveById {
		indexById[id] = nextIndex
		flowByIndex[nextIndex] = valve.flowRate
		nextIndex++
	}

	distBetween := [][]int{}
	for i := 0; i < numValves; i++ {
		distBetween = append(distBetween, make([]int, numValves))
		for j := 0; j < numValves; j++ {
			distBetween[i][j] = 100_000_000
		}
	}

	for _, valve := range valveById {
		idx, _ := indexById[valve.id]
		distBetween[idx][idx] = 0
		for _, nextId := range valve.connectedValves {
			nextIdx, _ := indexById[nextId]
			distBetween[idx][nextIdx] = 1
		}
	}

	for k := 0; k < numValves; k++ {
		for i := 0; i < numValves; i++ {
			for j := 0; j < numValves; j++ {
				distBetween[i][j] = intMin(
					distBetween[i][j],
					distBetween[i][k]+distBetween[k][j],
				)
			}
		}
	}

	p1 := func() {
		type dfsState struct {
			time     int
			idx      int
			isOpened bitset
		}

		startIdx, _ := indexById["AA"]

		memo := map[dfsState]int{}
		var dfs func(dfsState) int
		dfs = func(s dfsState) int {
			if result, ok := memo[s]; ok {
				return result
			}

			maxReleased := 0
			for nextIdx := 0; nextIdx < numValves; nextIdx++ {
				timeLeft := s.time - distBetween[s.idx][nextIdx] - 1
				if !s.isOpened.get(nextIdx) && timeLeft > 0 && flowByIndex[nextIdx] > 0 {
					released := timeLeft * flowByIndex[nextIdx]
					maxReleased = intMax(
						maxReleased,
						released+dfs(dfsState{
							timeLeft,
							nextIdx,
							s.isOpened.set(nextIdx),
						}),
					)
				}
			}

			memo[s] = maxReleased
			return maxReleased
		}

		fmt.Printf("%d\n", dfs(dfsState{30, startIdx, bitset(0)}))
	}

	p2 := func() {
		type dfsState struct {
			time       int
			idx        int
			isOpened   bitset
			isElephant bool
		}

		startIdx, _ := indexById["AA"]

		memo := map[dfsState]int{}
		var dfs func(dfsState) int
		dfs = func(s dfsState) int {
			if result, ok := memo[s]; ok {
				return result
			}

			maxReleased := 0
			for nextIdx := 0; nextIdx < numValves; nextIdx++ {
				timeLeft := s.time - distBetween[s.idx][nextIdx] - 1
				if timeLeft > 0 && flowByIndex[nextIdx] > 0 && !s.isOpened.get(nextIdx) {
					released := timeLeft * flowByIndex[nextIdx]
					maxReleased = intMax(
						maxReleased,
						released+dfs(dfsState{
							timeLeft,
							nextIdx,
							s.isOpened.set(nextIdx),
							s.isElephant,
						}),
					)
				}
			}

			if !s.isElephant {
				maxReleased = intMax(
					maxReleased,
					dfs(dfsState{
						26,
						startIdx,
						s.isOpened,
						true,
					}),
				)
			}

			memo[s] = maxReleased
			return maxReleased
		}

		fmt.Printf("%d\n", dfs(dfsState{26, startIdx, bitset(0), false}))
	}

	_ = p1
	p2()
}
