package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	r := regexp.MustCompile(`<a class="audibleTile__artworkLink" href="(.*)">`)
	i := 0

	// Read lines from stdin
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		match := r.FindStringSubmatch(line)
		if len(match) == 0 {
			continue
		}
		i++
		fmt.Println(match[1])
	}
	fmt.Println("Total", i)
}
