package main

import "fmt"

func main() {

	ar := []int{10, 20, 30, 40, 50}
	fr := []int{3, 4, 0, 2, 1}

	//fmt.Println(exp(ar))
	index(ar, fr)
}

func exp(ar []int) []int {

	dl := make([]int, len(ar)*2)

	for i, val := range ar {
		dl[i] = val
		dl[len(ar)+i] = val
	}
	return dl
}

func index(ar []int, fr []int) {

	for _, val := range fr {
		fmt.Print(ar[val], " ")
	}
	return
}

