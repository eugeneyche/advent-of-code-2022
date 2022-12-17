package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	ruckSacks := []string{}

	for scanner.Scan() {
		ruckSacks = append(ruckSacks, scanner.Text())
	}

	p1 := func() {
		total := 0
		for _, ruckSack := range ruckSacks {
			halfLen := len(ruckSack) / 2
			first := ruckSack[:halfLen]
			second := ruckSack[halfLen:]
			lettersInFirst := map[rune]bool{}
			for _, c := range first {
				lettersInFirst[c] = true
			}
			var overlappingLetter rune
			for _, c := range second {
				if _, ok := lettersInFirst[c]; ok {
					overlappingLetter = c
				}
			}
			total += getLetterVal(overlappingLetter)
		}
		fmt.Printf("%d\n", total)
	}

	p2 := func() {
		total := 0
		for i := 0; i < len(ruckSacks); i += 3 {
			isOverlapping := map[rune]bool{}
			first := ruckSacks[i]
			second := ruckSacks[i+1]
			third := ruckSacks[i+2]

			for _, c := range first {
				isOverlapping[c] = true
			}

			for _, ruckSack := range []string{second, third} {
				newIsOverlapping := map[rune]bool{}
				for _, c := range ruckSack {
					if _, ok := isOverlapping[c]; ok {
						newIsOverlapping[c] = true
					}
				}
				isOverlapping = newIsOverlapping
			}

			var overlappingLetter rune
			for letter := range isOverlapping {
				overlappingLetter = letter
			}
			total += getLetterVal(overlappingLetter)
		}
		fmt.Printf("%d\n", total)
	}

	p1()
	p2()
}

func getLetterVal(l rune) int {
	if l >= 'a' {
		return int(l-'a') + 1
	}
	return int(l-'A') + 27
}
