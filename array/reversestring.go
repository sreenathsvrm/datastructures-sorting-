package main

import "fmt"

func main() {

	ar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(bS(ar, 4))

}

func bS(ar []int, val int) bool {

	var (
		start = 0
		end   = len(ar) - 1
	)
	for start <= end {
		mid := (start + (end - start)) / 2
		if val == ar[mid] {
			return true
		}
		if val < ar[mid] {
			start = mid - 1
		} else {
			start = mid + 1
		}
	}
	return false
}
