package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	L := []int{5, 8, 4, 345, 8, 9, 2, 1, 3, 4, 6, 9, 4, 3, 3, 6, 7, 3, 4, 5}
	fmt.Println("Unsorted List: ",L)
	//fmt.Println("Sorted List", quicksort(L))
	fmt.Println("Sorted List", mergesort(L))
}

func quicksort(L []int) []int{
	if len(L) == 1{return []int{L[0]}}
	if len(L) == 0{return []int{}}

	pivot := L[0] //naive approach -> random pivot selection should be used instead
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

func mergesort(L []int) []int{

	length := len(L)
	if length == 1{
		return L
	}
	left := L[:length/2]
	right := L[(length/2):]
	left = mergesort(left)
	right = mergesort(right)

	var rval []int
	for {
		if len(left) > 0 && len(right) > 0{
			l_first := left[0]
			r_first := right[0]
			if l_first < r_first{
				rval = append(rval, l_first)
				left = left[1:]
			} else {
				rval = append(rval, r_first)
				right = right[1:]
			}
		} else if len(left) > 0 {
			rval = append(rval, left...)
			left = left[:0]
		} else if len(right) > 0 {
			rval = append(rval, right...)
			right = right[:0]
		} else {
			break
		}
	}
	return rval
}
