package main

import (
	"fmt"
	"math/rand"
	"time"
)

type intSlice []int

func newIntSlice(size int) intSlice {
	sampleSlice := intSlice{}

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for j := 0; j <= size; j++ {
		sampleSlice = append(sampleSlice, r.Intn(size-1))
	}

	return sampleSlice
}

func (is intSlice) checkEvenOrOdd() {
	for i := range is {
		if is[i]%2 == 0 {
			fmt.Println(is[i], " is even")
		} else {
			fmt.Println(is[i], " is odd")
		}
	}
}
