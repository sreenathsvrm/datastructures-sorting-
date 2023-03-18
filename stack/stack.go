package main

import "fmt"

func main() {

	ar:=[]int{1,23,4,5,6,7}
	

}

type node struct {
	data int
	next *node
}

type stack struct {
	head *node
}

func (s *stack) Push(data int) {

	newNode := &node{data: data}

	//push head to the next of newNode's next if head is nil or not
	newNode.next = s.head
	//change newNode as head
	s.head = newNode
}

// funciton to push multiple values to stack from user value
func (s *stack) MultiplePush() {
	var limit, data int

	fmt.Print("Enter how much values you need to enter: ")
	fmt.Scan(&limit)

	for i := 1; i <= limit; i++ {
		fmt.Print("Enter value: ")
		fmt.Scan(&data)
		s.Push(data)
	}
}

// funciton to delete a value from stack
// in stack head value been deleted
func (s *stack) Pop() int {
	if checkStackEmpty(s.head, true) {
		return -1
	}
	//move head to its next
	data := s.head.data
	s.head = s.head.next
	return data
}

func (s *stack) Peek() int {

	if s.head == nil {
		return -1
	}

	return s.head.data
}

// function to display all values in tha stackhhhhhhhhg
func (s *stack) DisplayOrder() {

	if checkStackEmpty(s.head, true) {
		return
	}

	currentNode := s.head

	for currentNode != nil {
		fmt.Print(currentNode.data, " ")
		currentNode = currentNode.next
	}
	fmt.Println()
}

// function to check the stack is empty
func checkStackEmpty(node *node, show bool) bool {

	if node == nil && show {
		fmt.Println("The stack is empty")
	}
	return node == nil
}

func (s *stack) IsStackEmpty() bool {
	return s.head == nil
}

type stackz struct {
	stack1 stack
	stack2 stack
}

func (s *stackz) undo() {
	val := s.stack1.Pop()
	s.stack2.Push(val)
}

func (s *stackz) redo() {
	value := s.stack2.Pop()
	s.stack1.Push(value)

}

func (s*stack) middlelement() *node {

	
	currentNode:=s.head

	for currentNode!=nil{
		currentNode=currentNode.next.next
		mid:=currentNode.next

     if mid.next==nil{
	   return mid
	 }

	 if 
	}

}