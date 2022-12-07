package main

import (
	"bufio"
	"fmt"
	"os"
)

type instruction struct {
	num       int
	fromStack int
	toStack   int
}

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	stacks := [][]rune{}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Scan all stacks
	for scanner.Scan() {
		line := scanner.Text()
		if line[1] == '1' {
			break
		}

		for i := 1; i < len(line); i += 4 {
			crate := []rune(line)[i]
			stackIdx := i / 4

			// Ensure we have enough stacks horizontally
			if stackIdx >= len(stacks) {
				stacks = append(stacks, []rune{})
			}

			if crate == ' ' {
				continue
			}

			stacks[stackIdx] = append(stacks[stackIdx], crate)
		}
	}

	for _, stack := range stacks {
		reverseList(stack)
	}

	// Skip new line
	scanner.Scan()

	instructions := []instruction{}

	// Scan all instructions
	var num, fromStack, toStack int
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Sscanf(line, "move %d from %d to %d", &num, &fromStack, &toStack)
		instructions = append(instructions, instruction{
			num,
			// Decrement to make 0-indexed
			fromStack - 1,
			toStack - 1,
		})
	}

	p1 := func() {
		for _, instr := range instructions {
			for i := 0; i < instr.num; i++ {
				fromStack := stacks[instr.fromStack]
				crate := fromStack[len(fromStack)-1]
				stacks[instr.fromStack] = fromStack[:len(fromStack)-1]
				toStack := stacks[instr.toStack]
				stacks[instr.toStack] = append(toStack, crate)
			}
		}

		topCrates := []rune{}
		for _, stack := range stacks {
			if len(stack) == 0 {
				continue
			}

			topCrates = append(topCrates, stack[len(stack)-1])
		}
		fmt.Println(string(topCrates))
	}

	p2 := func() {
		for _, instr := range instructions {
			fromStack := stacks[instr.fromStack]
			crates := fromStack[len(fromStack)-instr.num:]
			stacks[instr.fromStack] = fromStack[:len(fromStack)-instr.num]
			toStack := stacks[instr.toStack]
			stacks[instr.toStack] = append(toStack, crates...)
		}

		topCrates := []rune{}
		for _, stack := range stacks {
			if len(stack) == 0 {
				continue
			}

			topCrates = append(topCrates, stack[len(stack)-1])
		}
		fmt.Println(string(topCrates))
	}

	_ = p1
	p2()
}

func reverseList[T any](l []T) {
	for i := 0; i < len(l)/2; i += 1 {
		opIdx := len(l) - i - 1
		l[i], l[opIdx] = l[opIdx], l[i]
	}
}
