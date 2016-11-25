package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Read lines from stdin
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}
