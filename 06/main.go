package main

import (
	"fmt"
	"os"
)

func main() {
	signal, _ := os.ReadFile(os.Args[1])

	p1 := func() {
		for i := 0; i < len(signal)-4; i++ {
			seen := map[byte]bool{}
			isUnique := true

			for j := 0; j < 4; j++ {
				c := signal[i+j]
				if _, ok := seen[c]; ok {
					isUnique = false
					break
				}
				seen[c] = true
			}

			if isUnique {
				fmt.Printf("%d\n", i+4)
				break
			}
		}
	}

	p2 := func() {
		for i := 0; i < len(signal)-14; i++ {
			seen := map[byte]bool{}
			isUnique := true

			for j := 0; j < 14; j++ {
				c := signal[i+j]
				if _, ok := seen[c]; ok {
					isUnique = false
					break
				}
				seen[c] = true
			}

			if isUnique {
				fmt.Printf("%d\n", i+14)
				break
			}
		}
	}

	p1()
	p2()
}
