package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type CircularList struct {
	head *Node
	tail *Node
}

func (c *CircularList) Insert(data int) {
	newNode := &Node{data: data}

	if c.head == nil {
		c.head = newNode
	} else {
		c.tail.next = newNode
	}
	c.tail = newNode
	c.tail.next = c.head

}

func (c *CircularList) Print() {
	current := c.head

	for current.next != c.head {
		fmt.Println(current.data)
		current = current.next
	}
	fmt.Println(current.data)
}
func main() {
	c := &CircularList{}
	c.Insert(1)
	c.Insert(2)
	c.Insert(3)
	c.Print()
}
