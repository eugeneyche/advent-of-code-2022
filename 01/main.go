package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	rawInput, _ := os.ReadFile(os.Args[1])
	input := string(rawInput)

	linesPerElf := strings.Split(input, "\n\n")
	caloriesPerElf := []int{}
	for _, inputs := range linesPerElf {
		totalCalories := 0
		for _, caloriesRaw := range strings.Split(inputs, "\n") {
			calories, _ := strconv.Atoi(caloriesRaw)
			totalCalories += calories
		}
		caloriesPerElf = append(caloriesPerElf, totalCalories)
	}
	sort.Ints(caloriesPerElf)

	p1 := func() {
		fmt.Printf("%d\n", caloriesPerElf[len(caloriesPerElf)-1])
	}

	p2 := func() {
		totalCalories := 0
		for _, calories := range caloriesPerElf[len(caloriesPerElf)-3:] {
			totalCalories += calories
		}
		fmt.Printf("%d\n", totalCalories)
	}

	p1()
	p2()
}
