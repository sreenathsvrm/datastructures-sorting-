package main

import "fmt"

type Node struct {
	next *Node
	data int
	prev *Node
}

type Doublelinkedlist struct {
	head *Node
	tail *Node
}

func (Dl *Doublelinkedlist) Add(data int) {
	NewNode := &Node{
		data: data,
	}
	if Dl.head == nil {
		Dl.head = NewNode
	} else {
		Dl.tail.next = NewNode
		NewNode.prev = Dl.tail
	}

	Dl.tail = NewNode
}

func (Dl *Doublelinkedlist) delete(data int) {

	if Dl.head == nil {
		return
	}

	currentNode := Dl.head

	for currentNode != nil {

		if currentNode.data == data {
			Dl.deleteHelper(currentNode)
			return
		}
		currentNode = currentNode.next
	}
}

func (Dl *Doublelinkedlist) deleteHelper(Node *Node) {

	if Node.prev != nil {
		Node.prev.next = Node.next
	} else {
		Dl.head = Node.next
	}

	if Node.next != nil {
		Node.next.prev = Node.prev
	} else {
		Dl.tail = Node.prev
	}
}
func (Dl *Doublelinkedlist) Print() {

	if Dl.head == nil {
		fmt.Println("List is empty")
		return
	}

	currentNode := Dl.head

	for currentNode != nil {
		fmt.Print(currentNode.data, " ")
		currentNode = currentNode.next

	}
	fmt.Println()
}

func main() {
	c := &Doublelinkedlist{}
	c.Add(1)
	c.Add(2)
	c.Add(3)
	c.Print()

	c.delete(1)
	c.Print()

	c.Add(6)
	c.Print()
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 