package main

import "fmt"

func main() {

}

type maxHeap struct {
	maxHeap []int
}

//func to Heapify an array

func (max *maxHeap) Heapify(array []int) {

	//max.maxHeap = array // assign the array as maxHeap

	// insted of assign append the new array to heap then we won't loss the old value
	max.maxHeap = append(max.maxHeap, array...)

	//find last parent and move throgh all parents
	parentIdx := max.parentIdx(len(max.maxHeap) - 1)

	for parentIdx >= 0 {

		max.shiftDown(parentIdx)
		parentIdx--
	}
}

// func to insert a new element to the max heap
func (max *maxHeap) Insert(data int) {

	// first appen the new value to the end of heap
	max.maxHeap = append(max.maxHeap, data)

	//then shift up the new element to its heap position
	max.shiftUp(len(max.maxHeap) - 1)
}

// func to remove the max value from the max heap
func (max *maxHeap) Remove() {

	// swap last value and top value ( in heap remove value from top only)
	max.maxHeap[0], max.maxHeap[len(max.maxHeap)-1] = max.maxHeap[len(max.maxHeap)-1], max.maxHeap[0]

	// delete the last element that we swappen using slice it out
	max.maxHeap = max.maxHeap[:len(max.maxHeap)-1]

	// then shift down the top element that we swapped to its heapify posision
	max.shiftDown(0)
}

// func to get the max value from the heap that is top value

func (max *maxHeap) Peek() int {

	return max.maxHeap[0]
}

// func to dispalay all values in the max heap
func (max *maxHeap) DisplayAllvalues() {

	for _, val := range max.maxHeap {
		fmt.Print(val, " ")
	}
	fmt.Println()
}

// *** helper funcion for max heap ***

// to get parent index using chid index
func (max *maxHeap) parentIdx(childIdx int) int {
	return (childIdx - 1) / 2
}

// func to get left child's index using parent index
func (max *maxHeap) leftChildIdx(parentIdx int) int {
	return (parentIdx * 2) + 1
}

// func to get right child's index using parent index
func (max *maxHeap) rightChildIdx(parentIdx int) int {
	return (parentIdx * 2) + 2
}

// func to shift down a elment in max heap
func (max *maxHeap) shiftDown(parentIdx int) {

	// first find left child
	leftChildIdx := max.leftChildIdx(parentIdx)

	endIdx := len(max.maxHeap) - 1 // last index of heap

	// loop until left child not valid or left or right value not to shit down
	for leftChildIdx <= endIdx {

		//assume left child as need to shift
		idxToShift := leftChildIdx

		//find right child and check valid then check it is the index to shift
		rightIdx := max.rightChildIdx(parentIdx)

		if rightIdx <= endIdx && max.maxHeap[rightIdx] > max.maxHeap[leftChildIdx] {
			idxToShift = rightIdx
		}

		// if parent value is greater than shift index vlaue then return
		if max.maxHeap[parentIdx] > max.maxHeap[idxToShift] {
			return
		}
		// otherwise swap values and find new parent and child
		max.maxHeap[idxToShift], max.maxHeap[parentIdx] = max.maxHeap[parentIdx], max.maxHeap[idxToShift]
		// current child as parent and find new left child
		parentIdx = idxToShift
		leftChildIdx = max.leftChildIdx(parentIdx)
	}
}

// func to shift up an element
func (max *maxHeap) shiftUp(childIdx int) {

	//find the parent of child
	parentIdx := max.parentIdx(childIdx)

	//shift until reach top or heap satisfy

	for parentIdx >= 0 && max.maxHeap[parentIdx] < max.maxHeap[childIdx] {

		//swap the values
		max.maxHeap[parentIdx], max.maxHeap[childIdx] = max.maxHeap[childIdx], max.maxHeap[parentIdx]

		// assign current parent as child and find new parent
		childIdx = parentIdx
		parentIdx = max.parentIdx(childIdx)
	}
}

