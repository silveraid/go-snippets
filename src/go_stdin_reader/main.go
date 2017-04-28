package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		fmt.Println("no pipe")
	} else {
		reader := bufio.NewReader(os.Stdin)
		for {
			input, err := reader.ReadString('\n')
			if err != nil && err == io.EOF {
				break
			}
			fmt.Printf("go: %s", input)
		}
	}
}
