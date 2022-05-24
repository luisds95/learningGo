package main

import (
	"fmt"
	"go_intro/first/morestrings"
)

func main() {
	str := "Hello :)"
	fmt.Println(str)
	fmt.Println(morestrings.ReverseRunes(str))
}
