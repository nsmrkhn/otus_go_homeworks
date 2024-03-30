package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	h := "Hello, OTUS!"
	h = stringutil.Reverse(h)
	fmt.Println(h)
}
