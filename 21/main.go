package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type monkeyJob struct {
	isValue bool
	value   int
	lhs     string
	rhs     string
	op      *operation
}

type operation struct {
	symbol        string
	do            func(int, int) int
	solveForLeft  func(int, int) int
	solveForRight func(int, int) int
}

var add = operation{
	"+",
	func(l, r int) int { return l + r },
	func(v, r int) int { return v - r },
	func(v, l int) int { return v - l },
}

var subtract = operation{
	"-",
	func(l, r int) int { return l - r },
	func(v, r int) int { return v + r },
	func(v, l int) int { return l - v },
}

var multiply = operation{
	"*",
	func(l, r int) int { return l * r },
	func(v, r int) int { return v / r },
	func(v, l int) int { return v / l },
}

var divide = operation{
	"/",
	func(l, r int) int { return l / r },
	func(v, r int) int { return v * r },
	func(v, l int) int { return l / v },
}

func invRhsDivide(v, l int) int {
	return l / v
}

func NewMonkeyJobValue(value int) monkeyJob {
	return monkeyJob{true, value, "", "", nil}
}

func NewMonkeyJobExpr(
	lhs,
	rhs string,
	op *operation,
) monkeyJob {
	return monkeyJob{false, 0, lhs, rhs, op}
}

func stripColon(s string) string {
	return s[:len(s)-1]
}

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	jobs := map[string]monkeyJob{}
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		monkey := stripColon(tokens[0])
		if len(tokens) == 2 {
			value, _ := strconv.Atoi(tokens[1])
			jobs[monkey] = NewMonkeyJobValue(value)
		} else {
			var op *operation
			switch tokens[2] {
			case "+":
				op = &add
			case "-":
				op = &subtract
			case "*":
				op = &multiply
			case "/":
				op = &divide
			}
			jobs[monkey] = NewMonkeyJobExpr(tokens[1], tokens[3], op)
		}
	}

	var eval func(monkey string) int
	eval = func(monkey string) int {
		job := jobs[monkey]
		if job.isValue {
			return job.value
		}
		return job.op.do(
			eval(job.lhs),
			eval(job.rhs),
		)
	}

	p1 := func() {
		fmt.Printf("%d\n", eval("root"))
	}

	p2 := func() {
		isHumnAncestor := map[string]bool{}
		var lookForHumns func(monkey string) bool
		lookForHumns = func(monkey string) bool {
			job := jobs[monkey]

			var isAncestor bool
			if job.isValue {
				isAncestor = monkey == "humn"
			} else {
				isAncestor = lookForHumns(job.lhs) || lookForHumns(job.rhs)
			}
			isHumnAncestor[monkey] = isAncestor
			return isAncestor
		}
		lookForHumns("root")

		var solveForHumn func(monkey string, targetValue int) int
		solveForHumn = func(monkey string, targetValue int) int {
			if monkey == "humn" {
				return targetValue
			}
			// Monkey must be an expr otherwise
			job := jobs[monkey]
			var lhs, rhs int
			if isHumnAncestor[job.lhs] {
				rhs = eval(job.rhs)
				lhs = job.op.solveForLeft(targetValue, rhs)
				return solveForHumn(job.lhs, lhs)
			} else {
				lhs = eval(job.lhs)
				rhs = job.op.solveForRight(targetValue, lhs)
				return solveForHumn(job.rhs, rhs)
			}
		}

		// I decided not to the do the scuffed thing:
		// override root's job to subtract and set target value to 0
		var startMonkey string
		var targetValue int
		root := jobs["root"]
		if isHumnAncestor[root.lhs] {
			startMonkey = root.lhs
			targetValue = eval(root.rhs)
		} else {
			startMonkey = root.rhs
			targetValue = eval(root.lhs)
		}
		fmt.Printf("%d\n", solveForHumn(startMonkey, targetValue))
	}

	p1()
	p2()
}
