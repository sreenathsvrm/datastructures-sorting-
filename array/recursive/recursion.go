	package main

import "fmt"

func main() {

	// fmt.Println(factorial(5))

	//fmt.Println(sum(1))

	fmt.Println(fibinoci(7))

}

func factorial(n int) (res int) {
	if n == 1 || n == 0 {
		return 1
	}

	return n * factorial(n-1)

}
func sum(n int) (product int) {

	if n == 100 {
		return n
	}

	return n + sum(n+1)

}

func fibinoci(n int) (ar []int) {
	//var ar []int
	ar = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 2; i <= n; i++ {
		ar[i] = ar[i-1] + ar[i-2]
	}
	return ar

}

