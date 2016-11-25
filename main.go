package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	urls := getUrls()
	for _, url := range urls {
		fmt.Println(url)
	}
	fmt.Println("Total", len(urls))
}

func getUrls() []string {
	urls := []string{}
	r := regexp.MustCompile(`<a class="audibleTile__artworkLink" href="(.*)">`)

	// Read lines from stdin
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		match := r.FindStringSubmatch(line)
		if len(match) == 0 {
			continue
		}
		urls = append(urls, match[1])
	}
	return urls
}
