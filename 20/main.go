package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type linkedListNode struct {
	next  *linkedListNode
	prev  *linkedListNode
	value int
}

type linkedList struct {
	head *linkedListNode
	len  int
}

func (ll *linkedList) insertAfter(where *linkedListNode, what *linkedListNode) {
	ll.len++
	where.next, what.next = what, where.next
	what.prev, what.next.prev = where, what
}

func (ll *linkedList) append(value int) *linkedListNode {
	if ll.head == nil {
		ll.head = &linkedListNode{nil, nil, value}
		ll.head.next = ll.head
		ll.head.prev = ll.head
		ll.len = 1
		return ll.head
	}
	node := &linkedListNode{nil, nil, value}
	ll.insertAfter(ll.head.prev, node)
	return node
}

func (ll *linkedList) remove(node *linkedListNode) {
	ll.len--
	if ll.head == node && node != node.next {
		ll.head = node.next
		if ll.head == node {
			ll.head = nil
		}
	}
	node.next.prev, node.prev.next = node.prev, node.next
	node.next = nil
	node.prev = nil
}

func (ll *linkedList) step(from *linkedListNode, n int) *linkedListNode {
	cursor := from
	if n > 0 {
		nMod := n % ll.len
		for i := 0; i < nMod; i++ {
			cursor = cursor.next
		}
	} else if n < 0 {
		nMod := -n % ll.len
		for i := 0; i < nMod; i++ {
			cursor = cursor.prev
		}
	}
	return cursor
}

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	p1List := linkedList{}
	p1Nodes := []*linkedListNode{}
	p2List := linkedList{}
	p2Nodes := []*linkedListNode{}

	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		p1Nodes = append(p1Nodes, p1List.append(value))
		p2Nodes = append(p2Nodes, p2List.append(value))
	}

	p1 := func() {
		var cursor *linkedListNode

		for _, node := range p1Nodes {
			cursor = node.prev
			p1List.remove(node)
			cursor = p1List.step(cursor, node.value)
			p1List.insertAfter(cursor, node)
		}

		for _, node := range p1Nodes {
			if node.value == 0 {
				cursor = node
			}
		}

		total := p1List.step(cursor, 1000).value
		total += p1List.step(cursor, 2000).value
		total += p1List.step(cursor, 3000).value
		fmt.Printf("%d\n", total)
	}

	p2 := func() {
		var cursor *linkedListNode
		dkey := 811589153

		for i := 0; i < 10; i++ {
			for _, node := range p2Nodes {
				cursor = node.prev
				p2List.remove(node)
				cursor = p2List.step(cursor, node.value*dkey)
				p2List.insertAfter(cursor, node)
			}
		}

		for _, node := range p2Nodes {
			if node.value == 0 {
				cursor = node
			}
		}

		total := p2List.step(cursor, 1000).value
		total += p2List.step(cursor, 2000).value
		total += p2List.step(cursor, 3000).value
		fmt.Printf("%d\n", total*dkey)
	}

	p1()
	p2()
}
