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

func reverse[T int | string](slice []T) []T {
	newInts := make([]T, len(slice))

	newIntsLen := len(slice) - 1
	for i := 0; i < len(slice); i++ {
		newInts[newIntsLen] = slice[i]

		newIntsLen--
	}

	return newInts
}
