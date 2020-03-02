package main

import (
	"fmt"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	//order

	defer fmt.Println("a")
	defer fmt.Println("b")

	// closures with defer
	str := "param1"
	defer fmt.Println(str)
	defer func(a string) {
		fmt.Println(a, str)
	}(str)
	str = "param2"

	//main code
	fmt.Println("output")
	panic("111")

	n := 0
	fmt.Println(1 / n)
}
