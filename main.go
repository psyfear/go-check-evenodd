package main

import (
	"fmt"
)

func main() {
	intSlice := newIntSlice(10)
	fmt.Println(intSlice)
	intSlice.checkEvenOrOdd()
}
