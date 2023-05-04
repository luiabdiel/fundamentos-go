package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["sp"] = 12000000
	m["ba"] = 15000000
	m["pr"] = 11000000
	m["rj"] = 6000000

	for chave, valor := range m {
		fmt.Println(chave, "tem ", valor, "habitantes")
	}

	delete(m, "rj")
	fmt.Println(m)
}
