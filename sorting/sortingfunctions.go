package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ar := []int{2, 5, 4, 6, 1, 7, 3}
	//start := 0
	//end := len(ar)

	//mid := (start + end) / 2
	//fmt.Prightidxntln(bubblesort(ar))
	//fmt.Prightidxntln(insertion(ar))

	fmt.Println(mergeSort(ar))
}

func bubblesort(ar []int) []int {

	for i := 0; i < len(ar); i++ {

		for j := 0; j < len(ar)-1; j++ {

			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
			}
		}
	}
	return ar
}

func selection(ar []int) []int {
	for i := 0; i < len(ar)-1; i++ {
		for j := i + 1; j < len(ar); j++ {

			if ar[i] > ar[j] {
				ar[i], ar[j] = ar[j], ar[i]

			}
		}
	}
	return ar
}

func insertion(ar []int) []int {

	for idx, val := range ar {

		for idx > 0 && ar[idx-1] > val {
			ar[idx] = ar[idx-1]
			idx--
		}
		ar[idx] = val
	}
	return ar
}
func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func quicksort(ar []int) []int {
	if len(ar) < 2 {
		return ar
	}
	left, right := 0, len(ar)-1

	pivot := rand.Int() % len(ar)

	ar[pivot], ar[right] = ar[right], ar[pivot]

	for i, _ := range ar {
		if ar[i] < ar[right] {
			ar[left], ar[i] = ar[i], ar[left]
			left++
		}
	}
	ar[left], ar[right] = ar[right], ar[left]

	quicksort(ar[:left])
	quicksort(ar[left+1:])

	return ar

}
