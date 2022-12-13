package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type operation int

const (
	add operation = iota
	multiply
)

const old = -1

type inspectFn struct {
	lhs int
	rhs int
	op  operation
}

type monkey struct {
	items          []int
	inspect        inspectFn
	testDivisible  int
	ifTrueThrowTo  int
	ifFalseThrowTo int
}

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	monkeys := []monkey{}

	for scanner.Scan() {
		scanner.Scan()
		itemStrs := strings.Split(scanner.Text()[len("  Starting items: "):], ", ")
		items := []int{}
		for _, itemStr := range itemStrs {
			item, _ := strconv.Atoi(itemStr)
			items = append(items, item)
		}

		scanner.Scan()
		var lhsStr, rhsStr, opStr string
		fmt.Sscanf(scanner.Text(), "  Operation: new = %s %s %s\n", &lhsStr, &opStr, &rhsStr)
		var lhs, rhs int
		var op operation
		if lhsStr == "old" {
			lhs = old
		} else {
			lhs, _ = strconv.Atoi(lhsStr)
		}
		if rhsStr == "old" {
			rhs = old
		} else {
			rhs, _ = strconv.Atoi(rhsStr)
		}
		if opStr == "*" {
			op = multiply
		} else {
			op = add
		}
		inspect := inspectFn{lhs, rhs, op}

		scanner.Scan()
		var testDivisible int
		fmt.Sscanf(scanner.Text(), "  Test: divisible by %d", &testDivisible)

		scanner.Scan()
		var ifTrueThrowTo int
		fmt.Sscanf(scanner.Text(), "    If true: throw to monkey %d", &ifTrueThrowTo)

		scanner.Scan()
		var ifFalseThrowTo int
		fmt.Sscanf(scanner.Text(), "    If false: throw to monkey %d", &ifFalseThrowTo)

		newMonkey := monkey{
			items,
			inspect,
			testDivisible,
			ifTrueThrowTo,
			ifFalseThrowTo,
		}

		monkeys = append(monkeys, newMonkey)

		scanner.Scan()
	}

	p1 := func() {
		inspectCount := []int{}
		for i := 0; i < len(monkeys); i++ {
			inspectCount = append(inspectCount, 0)
		}

		for i := 0; i < 20; i++ {
			for j := 0; j < len(monkeys); j++ {
				currentMonkey := &monkeys[j]
				inspectCount[j] += len(currentMonkey.items)
				for _, item := range currentMonkey.items {
					lhs := currentMonkey.inspect.lhs
					rhs := currentMonkey.inspect.rhs
					if lhs == old {
						lhs = item
					}
					if rhs == old {
						rhs = item
					}
					new := lhs + rhs
					if currentMonkey.inspect.op == multiply {
						new = lhs * rhs
					}
					new /= 3
					thrownToMonkey := &monkeys[currentMonkey.ifTrueThrowTo]
					if new%currentMonkey.testDivisible != 0 {
						thrownToMonkey = &monkeys[currentMonkey.ifFalseThrowTo]
					}
					thrownToMonkey.items = append(thrownToMonkey.items, new)
				}
				currentMonkey.items = []int{}
			}
		}
		sort.Ints(inspectCount)
		topInspect := inspectCount[len(monkeys)-1]
		secondInspect := inspectCount[len(monkeys)-2]
		monkeyBusiness := topInspect * secondInspect
		fmt.Printf("%d\n", monkeyBusiness)
	}

	p2 := func() {
		inspectCount := []int{}
		mod := 1
		for i := 0; i < len(monkeys); i++ {
			inspectCount = append(inspectCount, 0)
			mod *= monkeys[i].testDivisible
		}

		for i := 0; i < 10_000; i++ {
			for j := 0; j < len(monkeys); j++ {
				currentMonkey := &monkeys[j]
				inspectCount[j] += len(currentMonkey.items)
				for _, item := range currentMonkey.items {
					lhs := currentMonkey.inspect.lhs
					rhs := currentMonkey.inspect.rhs
					if lhs == old {
						lhs = item
					}
					if rhs == old {
						rhs = item
					}
					new := lhs + rhs
					if currentMonkey.inspect.op == multiply {
						new = lhs * rhs
					}
					new %= mod
					thrownToMonkey := &monkeys[currentMonkey.ifTrueThrowTo]
					if new%currentMonkey.testDivisible != 0 {
						thrownToMonkey = &monkeys[currentMonkey.ifFalseThrowTo]
					}
					thrownToMonkey.items = append(thrownToMonkey.items, new)
				}
				currentMonkey.items = []int{}
			}
		}
		sort.Ints(inspectCount)
		topInspect := inspectCount[len(monkeys)-1]
		secondInspect := inspectCount[len(monkeys)-2]
		monkeyBusiness := topInspect * secondInspect
		fmt.Printf("%d\n", monkeyBusiness)
	}

	_ = p1
	p2()
}
