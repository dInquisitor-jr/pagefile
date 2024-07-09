package main

// ReadAddress(address int) (byte, error): Reads a byte from the specified address.
// ReadPage(pageID int) ([]byte, error): Reads an entire page by its ID.

// gc24{}

func SearchString(mem *Memory) (string, error) {
	// fmt.Println("works?")

	// content, _ := mem.ReadPage(0)
	for i := 1; i < 10000; i++ {
		content1, _ := mem.ReadAddress(i)
		content2, _ := mem.ReadAddress(i + 1)
		if content1 == 'g' && content2 == 'c' {
			return "found" + string(i), nil
		}
	}
	// fmt.Println(string())
	return "", nil
}
