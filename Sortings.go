package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	L := []int{5, 8, 4, 345, 8, 9, 2, 1, 3, 4, 6, 9, 4, 3, 3, 6, 7, 3, 4, 5}
	fmt.Println(L)

	quicksort(L)
}

func quicksort(L []int) {
	pivot := L[0]
	var left []int
	var right []int

	for _, el := range L {
		if el <= pivot {
			left = append(left, el)
		} else {
			right = append(right, el)
		}
	}

	fmt.Println("Pivot: ", pivot)
	fmt.Println("Left: ", left)
	fmt.Println("Right: ", right)
	fmt.Println()
}
