package main

import (
	"fmt"
	"strconv"
)

// ReadAddress(address int) (byte, error): Reads a byte from the specified address.
// ReadPage(pageID int) ([]byte, error): Reads an entire page by its ID.

// gc24{}

func SearchString(mem *Memory) (string, error) {
	// fmt.Println("works?")
	// offset 23937
	// content1, _ := mem.ReadAddress(23937)
	// content2, _ := mem.ReadAddress(23938)
	// content3, _ := mem.ReadAddress(23939)

	// fmt.Println(content1, content2, content3)
	// len 1024
	// content, _ := mem.ReadPage(0)
	for i := 1; i < 10000; i++ {
		content, err := mem.ReadPage(0)
		if err != nil {
			fmt.Println(err)
			return "", err
		}

		// fmt.Println(len(content), "page length")

		for j := 1; j < len(content)-4; j++ {
			if content[j] == 'g' && content[j+1] == 'c' && content[j+2] == '2' && content[j+3] == '4' {
				return "found " + strconv.Itoa(i), nil
			}
		}

		// content1, _ := mem.ReadAddress(i)
		// content2, _ := mem.ReadAddress(i + 1)
		// content3, _ := mem.ReadAddress(i + 2)
		// content4, _ := mem.ReadAddress(i + 3)
		// if content1 == 'g' && content2 == 'c' && content3 == '2' && content4 == '4' {
		// 	return "found " + strconv.Itoa(i), nil
		// }
	}
	// fmt.Println(string())
	return "not found", nil
}
