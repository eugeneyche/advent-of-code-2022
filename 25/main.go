package main

import (
	"bufio"
	"fmt"
	"os"
)

type snafu string

var snafuDigitToInt = map[byte]int{
	'=': -2,
	'-': -1,
	'0': 0,
	'1': 1,
	'2': 2,
}

func snafuToInt(s snafu) int {
	total := 0
	for i := 0; i < len(s); i++ {
		d := snafuDigitToInt[s[i]]
		total = total*5 + d
	}
	return total
}

var intToSnafuDigit = map[int]byte{
	-2: '=',
	-1: '-',
	0:  '0',
	1:  '1',
	2:  '2',
}

func intToSnafu(v int) snafu {
	digits := []int{}
	for i := 0; v > 0; i++ {
		d := v % 5
		if d >= 3 {
			d -= 5
			v += 5
		}
		v /= 5
		digits = append(digits, d)
	}

	strBuilder := make([]byte, len(digits))
	for i := 0; i < len(digits); i++ {
		strBuilder[len(digits)-i-1] = intToSnafuDigit[digits[i]]
	}

	return snafu(strBuilder)
}

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	fuelReqs := []snafu{}

	for scanner.Scan() {
		fuelReqs = append(fuelReqs, snafu(scanner.Text()))
	}

	p1 := func() {
		total := 0
		for _, s := range fuelReqs {
			total += snafuToInt(s)
		}
		fmt.Printf("%s\n", intToSnafu(total))
	}

	p1()
}
