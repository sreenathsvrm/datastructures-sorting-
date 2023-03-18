package mian

import "fmt"

type minHeap struct {
	minHeap []int // to hold slice values if it given or new created
}

// func to heapify an array
func (min *minHeap) Heapify(array []int) {

	// appen it old array and heapify
	min.minHeap = append(min.minHeap, array...)

	//find last parent
	parentIdx := min.parent(len(array) - 1)

	// loop throgh all parent and heapify the array
	for parentIdx >= 0 {
		min.shiftDown(parentIdx)
		parentIdx-- // move to next parent
	}
}

// func to inser a new element in a heap
func (min *minHeap) Insert(data int) {

	// append the data to the last of heap and shift up
	min.minHeap = append(min.minHeap, data)

	//shift up the last element
	min.shiftUp(len(min.minHeap) - 1)
}

// func to remove a value from heap
func (min *minHeap) Remove() {

	//find the last value from the heap and swap with first value
	// then delete the last value that is the old first value
	min.minHeap[0], min.minHeap[len(min.minHeap)-1] = min.minHeap[len(min.minHeap)-1], min.minHeap[0]

	//delete the last value using slice out the last value
	min.minHeap = min.minHeap[:len(min.minHeap)-1] //slice from start to last index

	//then shif down the top value to its corrext positin
	min.shiftDown(0)

}

// func to get the top value from minHeap(min value)
func (min *minHeap) Peek() int {

	return min.minHeap[0]
}

// func to display all elements in a heap
func (min *minHeap) DisplayAllvalues() {

	for _, value := range min.minHeap {

		fmt.Print(value, " ")
	}
	fmt.Println()
}

// **** min heap helper functions ****

// func to shift down an element compare to its child to bottom if needed

func (min *minHeap) shiftDown(parentIdx int) {

	leftChildIdx := min.leftChild(parentIdx) // get left child index
	endIdx := len(min.minHeap) - 1           // this is the last index

	// check until the last child not availble

	for leftChildIdx <= endIdx {

		rightChildIdx := min.rightChild(parentIdx) //first get the right index
		indexToShift := leftChildIdx               // assume index to shif as left index

		//check right child index is valid then compare value to the left index
		if rightChildIdx <= endIdx && min.minHeap[rightChildIdx] < min.minHeap[leftChildIdx] {
			// if right index value is less than left index value the change indexTo shift as rightIndex
			indexToShift = rightChildIdx
		}

		//check the indexToShift with parent need to shif or not
		if min.minHeap[indexToShift] < min.minHeap[parentIdx] {
			//swap the values
			min.minHeap[indexToShift], min.minHeap[parentIdx] = min.minHeap[parentIdx], min.minHeap[indexToShift]
			parentIdx = indexToShift                // the shifted index as parent
			leftChildIdx = min.leftChild(parentIdx) // find the next child
		} else {
			return // if parent value smaller thean child value then directly return
		}
	}
}

// func to shift up an elemen comparare to its parent

func (min *minHeap) shiftUp(childIdx int) {

	parentIdx := min.parent(childIdx)

	// compare bottom to the last parent if value less
	for parentIdx >= 0 && min.minHeap[childIdx] < min.minHeap[parentIdx] {

		//swap child value and parent value
		min.minHeap[childIdx], min.minHeap[parentIdx] = min.minHeap[parentIdx], min.minHeap[childIdx]

		childIdx = parentIdx // assign child ass parent

		parentIdx = min.parent(childIdx) // get new child's parent
	}
}

// func to get index of parent by child index
func (min *minHeap) parent(ChildIndex int) int {

	return ((ChildIndex - 1) / 2) // its parent will be in its index -1 / 2 position
}

// func to get index of left child by parent index
func (min *minHeap) leftChild(prentIndex int) int {

	return ((prentIndex * 2) + 1)
}

// func to get index of right child by pretn inde
func (min *minHeap) rightChild(prentIndex int) int {

	return ((prentIndex * 2) + 2)
}
