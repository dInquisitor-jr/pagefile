package main

import (
	"fmt"
)

// ReadAddress(address int) (byte, error): Reads a byte from the specified address.
// ReadPage(pageID int) ([]byte, error): Reads an entire page by its ID.

func SearchString(mem *Memory) (string, error) {
	// fmt.Println("works?")
	content, _ := mem.ReadPage(0)
	fmt.Println(string(content))
	return "something", nil
}
