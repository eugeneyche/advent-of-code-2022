package main

import (
	"fmt"
	"os"
)

type Span struct {
	begin int
	end   int
}

type TestCase struct {
	aSpan Span
	bSpan Span
}

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	testCases := []TestCase{}

	var aBegin, aEnd, bBegin, bEnd int

	nItems, _ := fmt.Fscanf(file, "%d-%d,%d-%d", &aBegin, &aEnd, &bBegin, &bEnd)
	for nItems == 4 {
		testCases = append(testCases, TestCase{Span{aBegin, aEnd}, Span{bBegin, bEnd}})
		nItems, _ = fmt.Fscanf(file, "%d-%d,%d-%d", &aBegin, &aEnd, &bBegin, &bEnd)
	}

	p1 := func() {
		total := 0
		for _, testCase := range testCases {
			aSpan := testCase.aSpan
			bSpan := testCase.bSpan
			if fullyContains(aSpan, bSpan) || fullyContains(bSpan, aSpan) {
				total += 1
			}
		}
		fmt.Printf("%d\n", total)
	}

	p2 := func() {
		total := 0
		for _, testCase := range testCases {
			aSpan := testCase.aSpan
			bSpan := testCase.bSpan
			if overlaps(aSpan, bSpan) {
				total += 1
			}
		}
		fmt.Printf("%d\n", total)
	}

	p1()
	p2()
}

func fullyContains(a Span, b Span) bool {
	return a.begin <= b.begin && a.end >= b.end
}

func overlaps(a Span, b Span) bool {
	return a.begin <= b.end && a.end >= b.begin
}
