package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b := make([]byte, 8)

	for {
		n, err := file.Read(b)
		if err == io.EOF {
			break
		}
		fmt.Printf("Byte slice: %q\n", b[:n])
	}
}
