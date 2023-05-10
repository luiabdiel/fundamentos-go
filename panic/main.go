package main

import "os"

func main() {
	_, err := os.Open("z:/settings.txt")

	if err != nil {
		panic(err)
	}
}
