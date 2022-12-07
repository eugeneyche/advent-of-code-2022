package main

import (
	"fmt"
	"os"
)

type round struct {
	opVal int
	myVal int
}

type Hand int
type Outcome int

const (
	Rock Hand = iota
	Paper
	Scissors
)

const (
	Lose Outcome = iota
	Tie
	Win
)

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	rounds := []round{}

	var a, b int
	nScanned, _ := fmt.Fscanf(file, "%c %c\n", &a, &b)
	for nScanned == 2 {
		rounds = append(rounds, round{(a - 'A'), (b - 'X')})
		nScanned, _ = fmt.Fscanf(file, "%c %c\n", &a, &b)
	}

	p1 := func() {
		total := 0
		for _, round := range rounds {
			// Hand value
			myHand := Hand(round.myVal)
			opHand := Hand(round.opVal)
			roundScore := int(myHand) + 1
			// outcome = (myHand + 1 - opHand) % 3
			// An additional +3 is necessary to ensure modular subtraction is positive.
			roundOutcome := Outcome((myHand + 4 - opHand) % 3)
			if roundOutcome == Win {
				roundScore += 6
			} else if roundOutcome == Tie {
				roundScore += 3
			}
			total += roundScore
		}
		fmt.Printf("%d\n", total)
	}

	p2 := func() {
		total := 0
		for _, round := range rounds {
			opHand := Hand(round.opVal)
			roundOutcome := Outcome(round.myVal)
			// (opHand + outcome - 1) % 3 = myHand
			// An additional +3 is necessary to ensure modular subtraction is positive.
			myHand := Hand((int(opHand) + int(roundOutcome) + 2) % 3)
			roundScore := int(myHand) + 1
			if roundOutcome == Win {
				roundScore += 6
			} else if roundOutcome == Tie {
				roundScore += 3
			}
			total += roundScore
		}
		fmt.Printf("%d\n", total)
	}

	p1()
	p2()
}
