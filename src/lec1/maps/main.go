package main

import (
	"fmt"
)

func main() {
	stringsMap := map[string]int{}
	stringsMap["a"] = 1
	stringsMap["b"] = 2
	fmt.Printf("%#v\n", stringsMap)
	fmt.Printf("%#v\n", stringsMap["a"])
	fmt.Printf("%#v\n", stringsMap["d"])
}
