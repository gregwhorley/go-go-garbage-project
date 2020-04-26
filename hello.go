package main

import (
	"fmt"
	"github.com/gregwhorley/go-go-garbage-project/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Printf(morestrings.ReverseRunes("!oG olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
