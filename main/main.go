package main

import (
	"fmt"

	"github.com/hallazzang/read"
)

func main() {
	s, err := read.Int()
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
