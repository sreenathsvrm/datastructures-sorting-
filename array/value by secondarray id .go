package main

import "fmt"

func main() {

	ar := []int{1, 2, 3}
	fmt.Println(check(ar))

}

func check(ar []int) []int {

	newarray := make([]int, len(ar)*2)

	for i, val := range ar {

		newarray[i] = val
		newarray[len(ar)+i] = val
	}
	return newarray
}
