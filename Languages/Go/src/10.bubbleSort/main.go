package main

import (
	"fmt"
)

func main() {
	list, err := fmt.Scanln(&int2)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("%d\n %d\n", int2, list)

}

// func bubbleSort() {
//
// }

