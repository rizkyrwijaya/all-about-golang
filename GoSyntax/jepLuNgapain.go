package main

import "fmt"

func main() {
	EvenNum := [5]int{0, 2, 4, 6, 8}

	for i, value := range EvenNum {
		fmt.Println(value, i)
	}

	numSlice := []int{5, 4, 3, 2, 1}
	sliced := numSlice[0:]
	fmt.Println(sliced)

	slice2 := make([]int, 5, 10)

	copy(slice2, numSlice)

	fmt.Println(slice2)

	slice3 := append(numSlice, 3, 0, -1)

	fmt.Println(slice3)
}
