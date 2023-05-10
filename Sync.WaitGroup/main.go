package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go callDatabase(&wg)
	go callAPI(&wg)
	go processInternal(&wg)

	wg.Wait()
}

func callDatabase(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)

	fmt.Println("Finalizado callDatabase")
	wg.Done()
}

func callAPI(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)

	fmt.Println("Finalizado callAPI")
	wg.Done()
}

func processInternal(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)

	fmt.Println("Finalizado processInternal")
	wg.Done()
}
