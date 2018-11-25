package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	L := []int{5, 8, 4, 345, 8, 9, 2, 1, 3, 4, 6, 9, 4, 3, 3, 6, 7, 3, 4, 5}
	fmt.Println("Unsorted List: ",L)
	fmt.Println("Sorted List", quicksort(L))
}

func quicksort(L []int) []int{
	if len(L) == 1{return []int{L[0]}}
	if len(L) == 0{return []int{}}

	pivot := L[0]
	var left []int
	var right []int

	for _, el := range L[1:] {
		if el <= pivot {
			left = append(left, el)
		} else {
			right = append(right, el)
		}
	}

	joined := append(quicksort(left), pivot)
	joined = append(joined, quicksort(right)...)
	return joined
}
