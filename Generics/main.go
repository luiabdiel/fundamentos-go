package main

import "fmt"

func main() {
	sliceInt := []int{5, 1, 2, 3}
	sliceString := []string{"a", "b", "c", "d"}
	newInts := reverse(sliceInt)
	newStrings := reverse(sliceString)

	fmt.Println(newInts)
	fmt.Println(newStrings)
}

type constrainCustom interface {
	int | string
}

func reverse[T constrainCustom](slice []T) []T {
	newSlices := make([]T, len(slice))

	newSlicesLen := len(slice) - 1
	for i := 0; i < len(slice); i++ {
		newSlices[newSlicesLen] = slice[i]

		newSlicesLen--
	}

	return newSlices
}
