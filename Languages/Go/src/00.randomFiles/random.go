package random

import (
	"fmt"
)

func randomMain() {
	fmt.Println("Hello, World!")
}

func quicksort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	pivot := arr[len(arr)/2]
	left := make([]int, 0, len(arr)/2)
	right := make([]int, 0, len(arr)/2)
	for _, val := range arr {
		if val < pivot {
			left = append(left, val)
		} else if val > pivot {
			right = append(right, val)
		}
	}
	quicksort(left)
	fmt.Println(pivot)
	quicksort(right)
}
