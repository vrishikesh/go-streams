package main

import (
	"fmt"
	"io"
	"os"

	r "github.com/vrishikesh/go-streams/reader"
)

func main() {
	// reader := r.NewStringReader("Clear is better than clever")
	reader := r.NewAlphaReader("Clear is better than clever")
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Println(string(p[:n])) //should handle any remainding bytes.
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(p[:n]))
	}
}
