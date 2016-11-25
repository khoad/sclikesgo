package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

// Usage: cat ~/Desktop/sclikes_html.txt | go run main.go
func main() {
	urls := getBrowserUrls()
	for _, url := range urls {
		wurl := getWaveFormUrl(url)
		fmt.Println(wurl)
	}
	fmt.Println("Total", len(urls))
}

func getWaveFormUrl(browserUrl string) string {
	// input
	// browserUrl := "https://soundcloud.com/lana-del-rey/ultraviolence-disciples-remix-1"
	resp, err := http.Get(browserUrl)
	if err != nil {
		fmt.Println("Error getting", browserUrl)
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading", browserUrl)
	}
	respString := string(respBytes)

	r := regexp.MustCompile(`"waveform_url":"(http.*\.json)"`)

	match := r.FindStringSubmatch(respString)
	waveformUrl := match[1]

	// output
	// waveformUrl := "https://wis.sndcdn.com/iCvi12jhGTIQ_m.json"
	return waveformUrl
}

func getBrowserUrls() []string {
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
		urls = append(urls, "https://soundcloud.com"+match[1])
	}
	return urls
}
