package main

import "fmt"

func main() {
	x := 5
	y := &x

	*y = 10

	fmt.Println(x, *y) // 10
	fmt.Println(&x, y) // 10

	//y está apontando para x, então se y agora vale 10, significa que ele está apontando para 10, no caso, seria x;
}
