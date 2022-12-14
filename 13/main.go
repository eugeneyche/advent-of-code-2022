package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type valueType int

const (
	numberType valueType = iota
	listType
)

type value struct {
	which  valueType
	number int
	list   []value
}

type packetPair struct {
	left  value
	right value
}

func newNumber(numberVal int) value {
	return value{numberType, numberVal, []value{}}
}

func newList(subVals ...value) value {
	return value{listType, 0, []value(subVals)}
}

func parseLineToValue(line string) value {
	rootList := newList()
	listStack := []*value{&rootList}
	// Ignore first [ and last ], so we the root list is kept in list stack
	cursor := 1
	for cursor < len(line)-1 {
		currentList := listStack[len(listStack)-1]
		if line[cursor] == '[' {
			newListValue := newList()
			currentList.list = append(currentList.list, newListValue)
			listStack = append(listStack, &currentList.list[len(currentList.list)-1])
			cursor++
		} else if isDigit(line[cursor]) {
			numberVal := 0
			for isDigit(line[cursor]) {
				numberVal = 10*numberVal + int(line[cursor]-'0')
				cursor++
			}
			currentList.list = append(currentList.list, newNumber(numberVal))
		} else if line[cursor] == ']' {
			listStack = listStack[:len(listStack)-1]
			cursor++
		} else {
			cursor++
		}
	}
	return rootList
}

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	packetPairs := []packetPair{}
	for scanner.Scan() {
		left := parseLineToValue(scanner.Text())
		scanner.Scan()
		right := parseLineToValue(scanner.Text())
		packetPairs = append(packetPairs, packetPair{left, right})
		scanner.Scan()
	}

	p1 := func() {
		total := 0
		for index, pair := range packetPairs {
			if compare(&pair.left, &pair.right) >= 0 {
				total += index + 1
			}
		}
		fmt.Printf("%d\n", total)
	}

	p2 := func() {
		twoDivider := newList(newList(newNumber(2)))
		sixDivider := newList(newList(newNumber(6)))
		allValues := []*value{&twoDivider, &sixDivider}
		for i := 0; i < len(packetPairs); i++ {
			allValues = append(allValues, &packetPairs[i].left, &packetPairs[i].right)
		}
		sort.Sort(byCompare(allValues))
		var twoIndex, sixIndex int
		for i := 0; i < len(allValues); i++ {
			if allValues[i] == &twoDivider {
				twoIndex = i + 1
			}
			if allValues[i] == &sixDivider {
				sixIndex = i + 1
			}
		}
		fmt.Printf("%d\n", twoIndex*sixIndex)
	}

	_ = p1
	p2()
}

func isDigit(b byte) bool {
	numberVal := b - '0'
	return 0 <= numberVal && numberVal < 10
}

func compare(left, right *value) int {
	if left.which == numberType && right.which == numberType {
		return right.number - left.number
	}
	if left.which == listType && right.which == listType {
		for i := 0; i < len(left.list); i++ {
			if i >= len(right.list) {
				return -1
			}
			compValue := compare(&left.list[i], &right.list[i])
			if compValue != 0 {
				return compValue
			}
		}
		if len(left.list) < len(right.list) {
			return 1
		}
		return 0
	}
	if left.which == numberType {
		newLeft := newList(newNumber(left.number))
		return compare(&newLeft, right)
	} else {
		newRight := newList(newNumber(right.number))
		return compare(left, &newRight)
	}
}

type byCompare []*value

func (a byCompare) Len() int {
	return len(a)
}

func (a byCompare) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byCompare) Less(i, j int) bool {
	return compare(a[i], a[j]) > 0
}

func prettyPrint(val *value) {
	prettyPrintImpl(val)
	fmt.Println()
}

func prettyPrintImpl(val *value) {
	if val.which == numberType {
		fmt.Printf("%d", val.number)
	} else if val.which == listType {
		fmt.Printf("[")
		for i, subVal := range val.list {
			if i > 0 {
				fmt.Printf(",")
			}
			prettyPrint(&subVal)
		}
		fmt.Printf("]")
	}
}
