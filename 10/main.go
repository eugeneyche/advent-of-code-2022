package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type operation int

const (
	noop operation = iota
	addx
)

type instruction struct {
	op  operation
	arg int
}

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	instrs := []instruction{}

	for scanner.Scan() {
		instrParts := strings.Split(scanner.Text(), " ")
		opStr := instrParts[0]
		instr := instruction{noop, 0}
		switch opStr {
		case "addx":
			addxArg, _ := strconv.Atoi(instrParts[1])
			instr = instruction{addx, addxArg}
		}
		instrs = append(instrs, instr)
	}

	p1 := func() {
		xReg := 1
		clk := 0
		total := 0

		sampleSignal := func() {
			if clk == 20 || clk == 60 || clk == 100 || clk == 140 || clk == 180 || clk == 220 {
				total += clk * xReg
			}
		}

		for _, instr := range instrs {
			switch instr.op {
			case noop:
				clk++
				sampleSignal()
			case addx:
				clk++
				sampleSignal()
				clk++
				sampleSignal()
				xReg += instr.arg
			}
		}

		fmt.Printf("%d\n", total)
	}

	p2 := func() {
		xReg := 1
		clk := 0
		crt := [6][40]rune{}

		draw := func() {
			xScan := (clk - 1) % 40
			yScan := (clk - 1) / 40

			xDelta := xReg - xScan
			if -1 <= xDelta && xDelta <= 1 {
				crt[yScan][xScan] = '#'
			} else {
				crt[yScan][xScan] = ' '
			}
		}

		for _, instr := range instrs {
			switch instr.op {
			case noop:
				clk++
				draw()
			case addx:
				clk++
				draw()
				clk++
				draw()
				xReg += instr.arg
			}
		}

		for y := 0; y < 6; y++ {
			fmt.Println(string(crt[y][:]))
		}
	}

	_ = p1
	p2()
}
