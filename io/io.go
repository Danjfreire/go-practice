package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// creates a read stream of string data
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)

	// reads the stream 8 bytes at a time until reaches the end (EOF)
	for {
		n, err := r.Read(b)
		fmt.Printf("n=%v err=%v b=%v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
