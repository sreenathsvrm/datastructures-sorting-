package main

import (
	"fmt"
	"math"
)

func main() {

	dr := []int{5, 3, 6, 2, 7, 2, 7}
	list1 := &linkedlist1{}

	list1.add(1)
	list1.add(2)
	list1.add(3)
	list1.add(2)
	list1.add(1)

	list1.priintds()

	//list1.delete(5)
	//list1.priintds()

	//list1.priintds()

	//fmt.Println(list1.thirdl())

	//fmt.Println(list1.pal())
	//list1.sld()
	list1.al(dr)
	fmt.Println()

	list1.priintds()

}

type Node struct {
	data int
	next *Node
}

type linkedlist1 struct {
	head *Node
	tail *Node
}

func (l *linkedlist1) al(dr []int) {

	for _, val := range dr {
		l.add(val)
	}

}

func (l *linkedlist1) add(data int) {

	newnode := &Node{data: data}

	if l.head == nil {
		l.head = newnode
	} else {
		l.tail.next = newnode
	}
	l.tail = newnode
}

func (l *linkedlist1) del(val int) {

	if l.head == nil {
		return
	}

	if l.head.data == val {
		l.head = l.head.next
		return
	}

	currentNode := l.head

	for currentNode != nil {

		if currentNode.next != nil && currentNode.next.data == val {
			currentNode.next = currentNode.next.next
			return
		}
		currentNode = currentNode.next
	}
}

func (l *linkedlist1) prepend(val int) {

	temp := &Node{data: val}

	if l.head != nil {
		l.head = temp
	} else {
		temp.next = l.head
		l.head = temp
	}
}

func (l *linkedlist1) priintds() {
	if l.head == nil {
		return
	}

	currentNode := l.head

	for currentNode != nil {
		fmt.Println(currentNode.data, "")
		currentNode = currentNode.next
	}
	fmt.Println()
}

func (l *linkedlist1) delete(val int) {

	if l.head == nil {
		return
	}

	if l.head.data == val {
		l.head = l.head.next
		return
	}
	head := l.head

	for head != nil {
		if head.next != nil && head.next.data == val {
			head.next = head.next.next

			// next node nil means its the tail
			if head.next == nil {
				l.tail = head
			}
			return
		}
		head = head.next
	}

}

func (l *linkedlist1) thirdl() int {
	var tl, sl, fl = math.MinInt, math.MinInt, math.MinInt

	if l.head == nil {
		return -1
	}

	currentNode := l.head

	for currentNode != nil {
		if currentNode.data > fl {
			tl = sl
			sl = fl
			fl = currentNode.data
		} else if currentNode.data > sl {
			tl = sl
			sl = currentNode.data
		} else if currentNode.data > tl {
			tl = currentNode.data
		}
		currentNode = currentNode.next

	}

	return tl

}

func (l *linkedlist1) pal() bool {

	gr := []int{}

	for l.head != nil {

		gr = append(gr, l.head.data)
		l.head = l.head.next

	}

	for i := 0; i < (len(gr)-1)/2; i++ {

		j := len(gr) - 1 - i

		if gr[i] != gr[j] {
			return false
		}
	}
	return true
}

func (l *linkedlist1) sld() {

	currentNode := l.head

	for currentNode != nil {

		if currentNode.next.next.next == nil {
			currentNode.next = currentNode.next.next
			return
		}

		currentNode = currentNode.next
	}

}
