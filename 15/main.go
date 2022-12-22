package main

import (
	"fmt"
	"os"
	"sort"
)

type ivec2 struct {
	x int
	y int
}

type sensor struct {
	pos           ivec2
	closestBeacon ivec2
	mDist         int
}

func iAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func mDist(a, b ivec2) int {
	return iAbs(a.x-b.x) + iAbs(a.y-b.y)
}

type ByPosX []sensor

func (a ByPosX) Len() int {
	return len(a)
}

func (a ByPosX) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByPosX) Less(i, j int) bool {
	return a[i].pos.x < a[j].pos.x
}

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	var sx, sy, bx, by int
	sensors := []sensor{}

	for true {
		if _, err := fmt.Fscanf(file, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by); err != nil {
			break
		}
		sPos := ivec2{sx, sy}
		bPos := ivec2{bx, by}
		mDist := mDist(sPos, bPos)
		sensors = append(sensors, sensor{sPos, bPos, mDist})
	}

	minSensorX := 0
	maxSensorX := 0
	maxMDist := 0
	for _, s := range sensors {
		if s.pos.x < minSensorX {
			minSensorX = s.pos.x
		}
		if s.pos.x > maxSensorX {
			maxSensorX = s.pos.x
		}
		if s.mDist > maxMDist {
			maxMDist = s.mDist
		}
	}

	p1 := func() {
		rowY := 20
		if os.Args[1] == "sample.txt" {
			rowY = 10
		}
		total := 0
		for x := minSensorX - maxMDist; x <= maxSensorX+maxMDist; x++ {
			inRange := false
			for _, s := range sensors {
				mDistToSensor := mDist(ivec2{x, rowY}, s.pos)
				if s.mDist >= mDistToSensor {
					inRange = true
					break
				}
			}
			if inRange {
				total++
			}
		}
		seen := map[ivec2]bool{}
		for _, s := range sensors {
			if s.closestBeacon.y != rowY {
				continue
			}
			if _, ok := seen[s.closestBeacon]; !ok {
				seen[s.closestBeacon] = true
				total -= 1
			}
		}

		fmt.Printf("%d\n", total)
	}

	p2 := func() {
		searchSize := 4_000_000
		if os.Args[1] == "sample.txt" {
			searchSize = 20
		}

		sort.Sort(ByPosX(sensors))

		var dBeaconPos ivec2
		for y := 0; y <= searchSize; y++ {
			x := 0
			for _, s := range sensors {
				distY := iAbs(y - s.pos.y)
				distX := s.mDist - distY
				if distX < 0 {
					continue
				}
				beginX := s.pos.x - distX
				endX := s.pos.x + distX
				if beginX <= x && x <= endX {
					x = endX + 1
				}
			}
			if x <= searchSize {
				dBeaconPos = ivec2{x, y}
				break
			}
		}

		fmt.Printf("%d\n", dBeaconPos.x*4_000_000+dBeaconPos.y)
	}

	p1()
	p2()
}
