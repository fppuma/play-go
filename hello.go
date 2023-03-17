package main

import (
	"example/user/hello/morestrings"
	"fmt"
)

func main() {
	fmt.Println("Hello, world")
	newMessage := morestrings.ReverseRunes("!oG, olleH")
	fmt.Println(newMessage)
}
