package main

import (
	"fmt"
	"math"
)

func main() {

	br := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	//fmt.Println(ist.arlink(br))
	//	d := arlink(br)

	ist := list{}

	ist.arlink(br)
	//ist.prin()
	//	ist.append(55)

	//ist.prin()
	//ist.prepend(7)
	//ist.prin()
	//fmt.Println()
	//	ist.remove(3)
	ist.prin()

	fmt.Println()

	fmt.Println(ist.third())

}

type Node struct {
	data int
	Next *Node
}

type list struct {
	head *Node
	tail *Node
}

func (l *list) arlink(br []int) {

	for _, val := range br {
		l.append(val)
	}

}

func (l *list) prin() {

	head := l.head

	if head == nil {
		return
	}

	for head != nil {
		fmt.Println(head.data, "")
		head = head.Next
	}

}

func (l *list) append(data int) {
	Newnode := &Node{data: data}

	if l.head == nil {
		l.head = Newnode
	} else {
		l.tail.Next = Newnode
	}
	l.tail = Newnode
}

func (l *list) prepend(data int) {
	temp := &Node{data: data}
	if l.head == nil {
		l.head = temp
	} else {
		temp.Next = l.head
		l.head = temp
	}

}

func (l *list) remove(data int) {
	if l.head == nil {
		return
	}
	if l.head.data == data {
		l.head = l.head.Next
		return
	}

	currentNode := l.head
	for currentNode != nil {
		if currentNode.Next != nil && currentNode.Next.data == data {
			currentNode.Next = currentNode.Next.Next
			return
		}
		currentNode = currentNode.Next

	}
}

func (l *list) third() int {
	fl, sl, tl := math.MinInt, math.MinInt, math.MinInt

	element := l.head

	for element != nil {
		if element.data > fl {
			tl = sl
			sl = fl
			fl = element.data
		} else if element.data > sl {
			tl = sl
			sl = element.data
		} else if element.data > tl {
			tl = element.data
		}
		element = element.Next
	}
	return tl
}
