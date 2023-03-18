package main

import (
	"fmt"
	"strings"
)

func init() {
	//fmt.Println("hiii")
}

func isPalindrome(s string) bool {

	var charArray = []rune{}

	for _, char := range s {

		if int(char) >= 97 && int(char) <= 122 || int(char) >= 65 && int(char) <= 90 {

			charArray = append(charArray, char)
		}
	}

	s = string(charArray)

	s = strings.ToLower(s)
	fmt.Println(s)

	for i := 0; i <= len(s)/2; i++ {

		if s[i] != s[len(s)-(i+1)] {
			return false
		}
	}

	return true

}

func main() {
	//	fmt.Println("hello")

	//	print(arraytolinked([]int{1, 23, 5, 4, 2}))

	//	 fmt.Println(isPalindrome("A man, a plan, a canal: panama"))

	array := []int{1, 2, 4, 5, 6, 7}

	//	 s := string(array)

	//	 s = strings.ToLower(s)
	//	 fmt.Println("Array", array)

	//array = insert(array, 2, 10)
	fmt.Println(insert(array, 2, 10))

	//	 fmt.Println("\nAfter insertion\n", array)

	//	 array = deleteByIndex(array, 2)

	//	fmt.Println("\nAfter deletion\n", array)

	//reverse(array)

	//	fmt.Println("\nAfter reverse\n", array)

	//	 n := &node{data: 1}

	//	 n.next = &node{data: 0}
	//	 n.next.next = &node{data: 1}

	//	 val := binaryToDecimal(n, 0)

	//	 fmt.Println(val)

	//fmt.Println(binarySearch(array, 1))

	//	 mer := merge([]int{1, 2, 4}, []int{6, 9, 7, 8})

	// fmt.Println(mer)
}

// given value will be inserted into next of given index
// considering array all index have values
// so need to create a new array to insert
func insert(array []int, index int, value int) []int {

	newArray := make([]int, len(array)+1)

	newIdx := 0
	oldIdx := 0

	for newIdx < len(newArray) {

		if newIdx == index {
			newArray[newIdx] = value
			newIdx++
			continue
		}

		newArray[newIdx] = array[oldIdx]
		oldIdx++
		newIdx++
	}

	return newArray
}

// array have a fixed size so creating a new array
// and copy all value except delete value
func deleteByIndex(array []int, index int) []int {

	newArray := make([]int, len(array)-1)

	newIdx := 0

	for i, val := range array {

		if i == index {
			continue
		}
		newArray[newIdx] = val
		newIdx++
	}

	return newArray
}

func reverse(array []int) []int {

	rightIdx := len(array) - 1

	for leftIdx := 0; leftIdx <= len(array)/2; leftIdx++ {

		array[leftIdx], array[rightIdx] = array[rightIdx], array[leftIdx]

		rightIdx--

	}

	return array

}

func binarySearch(array []int, val int) bool {

	var (
		start = 0
		end   = len(array) - 1
	)

	for start <= end {

		mid := start + ((end - start) / 2)

		if array[mid] == val {
			return true
		}

		if val < array[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return false
}

func merge(array1 []int, array2 []int) []int {

	var (
		l1, l2     = len(array1), len(array2)

		fIdx, sIdx = 0, l2 - 1

		newArray   = make([]int, l1+l2)
	)

	for i := 0; i < len(newArray); i++ {

		if fIdx < l1 {

			newArray[i] = array1[fIdx]

			fIdx++

		}

		if sIdx >= 0 {
			newArray[len(newArray)-(i+1)] = array2[sIdx]

			sIdx--
		}
	}

	return newArray
}

type node struct {
	data int
	next *node
}

func arraytolinked(ar []int) *node {

	var head, tail *node

	for _, val := range ar {

		newNode := &node{data: val}

		if head == nil {
			head = newNode
		} else {
			tail.next = newNode
		}
		tail = newNode
	}

	return head

}

func print(head *node) {

	if head == nil {

		return

	}

	for head != nil {

		fmt.Print(head.data, " ")

		head = head.next

	}
}
