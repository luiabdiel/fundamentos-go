package main

import (
	"fmt"
	"reflect"
)

func main() {
	var numeroInt8 int8 = 127
	var numeroInt int

	numeroInt = int(numeroInt8)

	fmt.Println(numeroInt8)
	fmt.Println(reflect.TypeOf(numeroInt8))
	fmt.Println(numeroInt)
	fmt.Println(reflect.TypeOf(numeroInt))

	var numeroFloat32 = 127.9
	var numeroInt64 int64

	numeroInt64 = int64(numeroFloat32)

	fmt.Println(numeroFloat32)
	fmt.Println(reflect.TypeOf(numeroFloat32))
	fmt.Println(numeroInt64)
	fmt.Println(reflect.TypeOf(numeroInt64))
}
