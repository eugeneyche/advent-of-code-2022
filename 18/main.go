package main

import (
	"fmt"
	"os"
)

type ivec3 struct {
	x, y, z int
}

func (a ivec3) add(b ivec3) ivec3 {
	return ivec3{a.x + b.x, a.y + b.y, a.z + b.z}
}

func intMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func intMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	cubes := []ivec3{}
	var x, y, z int
	for true {
		if _, err := fmt.Fscanf(file, "%d,%d,%d\n", &x, &y, &z); err != nil {
			break
		}
		cubes = append(cubes, ivec3{x, y, z})
	}

	allDirections := []ivec3{{-1, 0, 0}, {1, 0, 0}, {0, -1, 0}, {0, 1, 0}, {0, 0, -1}, {0, 0, 1}}

	isOccupied := map[ivec3]bool{}
	for _, cube := range cubes {
		isOccupied[cube] = true
	}

	p1 := func() {
		surfaceArea := 0
		for _, cube := range cubes {
			for _, delta := range allDirections {
				pos := cube.add(delta)
				if _, ok := isOccupied[pos]; !ok {
					surfaceArea += 1
				}
			}
		}

		fmt.Printf("%v\n", surfaceArea)
	}

	p2 := func() {
		isOccupied := map[ivec3]bool{}
		for _, cube := range cubes {
			isOccupied[cube] = true
		}

		boundMin := cubes[0]
		boundMax := cubes[0]
		for _, cube := range cubes {
			boundMin.x = intMin(boundMin.x, cube.x)
			boundMin.y = intMin(boundMin.y, cube.y)
			boundMin.z = intMin(boundMin.z, cube.z)
			boundMax.x = intMax(boundMax.x, cube.x)
			boundMax.y = intMax(boundMax.y, cube.y)
			boundMax.z = intMax(boundMax.z, cube.z)
		}

		boundMin = boundMin.add(ivec3{-1, -1, -1})
		boundMax = boundMax.add(ivec3{1, 1, 1})

		hasGas := map[ivec3]bool{}

		canSpreadTo := func(pos ivec3) bool {
			_, occupied := isOccupied[pos]
			_, gasAlready := hasGas[pos]
			return !occupied && !gasAlready &&
				boundMin.x <= pos.x && pos.x <= boundMax.x &&
				boundMin.y <= pos.y && pos.y <= boundMax.y &&
				boundMin.z <= pos.z && pos.z <= boundMax.z
		}

		var floodWithGas func(ivec3)
		floodWithGas = func(pos ivec3) {
			hasGas[pos] = true
			for _, delta := range allDirections {
				nextPos := pos.add(delta)
				if canSpreadTo(nextPos) {
					floodWithGas(nextPos)
				}
			}
		}

		floodWithGas(boundMin)

		surfaceArea := 0
		for _, cube := range cubes {
			for _, delta := range allDirections {
				pos := cube.add(delta)
				if _, gas := hasGas[pos]; gas {
					surfaceArea += 1
				}
			}
		}

		fmt.Printf("%v\n", surfaceArea)
	}

	p1()
	p2()
}
