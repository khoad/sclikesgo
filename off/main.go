package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	urllib "net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {

	// Read lines from stdin
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		scurl := s.Text()
		if scurl == "" {
			continue
		}

		rex := regexp.MustCompile(`https://soundcloud.com/(.*)`)
		match := rex.FindStringSubmatch(scurl)
		songName := strings.Replace(match[1], "/", "_", 1)
		fullSongName := "/Users/knguyen/Desktop/sc/" + songName + ".mp3"

		if _, err := os.Stat(fullSongName); err == nil {
			fmt.Println("Song exists; skipping", fullSongName)
			continue
		}

		// scurl := "https://soundcloud.com/pattycrashed/patty-crash-kronic-pictures"
		offurl := off(scurl)

		// offurl := "https://cf-media.sndcdn.com/p0kVgs7gZukF.128.mp3?Policy=eyJTdGF0ZW1lbnQiOlt7IlJlc291cmNlIjoiKjovL2NmLW1lZGlhLnNuZGNkbi5jb20vcDBrVmdzN2dadWtGLjEyOC5tcDMiLCJDb25kaXRpb24iOnsiRGF0ZUxlc3NUaGFuIjp7IkFXUzpFcG9jaFRpbWUiOjE0ODAwODk3NDZ9fX1dfQ__&Signature=ohFuv~4FpKDVPV-vv7q-hsBFGDvLfu08A1Noa1Xr7NsI8BgbQsT6eCglwxw6UKEAwj8FPn4lbLFAKUNyhsEq-1cy8HWiYq0lPur207haDJPhavz83xHtqYyuezyZ4k5uZej8fI7ORND-ZnzcDAviXOfKrYQICCBZMMbcfjg5cZRtiJn2x3E4fOcFmKRsd~q70lc4R4j9hBEU8~9TgAQhaigmGYKmv2PAYvtQna76rZoHgStDMtwzyD3BuTRPlnKHS4q9d85ZL5AaGIT-oBixxg43S~ymTSYvkgqu~FMcaSLZD~Dzu-Ev8RSMZpFj~zR3orRUW-BY4NejiXbJ1-nWZA__&Key-Pair-Id=APKAJAGZ7VMH2PFPW6UQ"
		if offurl == "" {
			fmt.Println("Failed", scurl)
			ioutil.WriteFile(fullSongName+".failed", []byte(""), 0644)
		} else {
			songContent := download(offurl)
			ioutil.WriteFile(fullSongName, songContent, 0644)
		}
	}
}

func off(url string) string {
	sleepsec := 5
	fmt.Println("Sleeping", sleepsec, "seconds")
	time.Sleep(5 * time.Second)

	fmt.Println("Offing...", url)
	resp, err := http.PostForm("http://offliberty.com/off03.php",
		urllib.Values{"track": {url}})
	if err != nil {
		fmt.Println("ERR getting", url, err)
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Status not OK", resp.Status)
		return ""
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERR reading", url, err)
		return ""
	}

	respString := string(respBytes)
	// fmt.Println("Off resp", respString)

	rex := regexp.MustCompile(`<A HREF="(.*)" class="download" rel="noreferrer"`)
	match := rex.FindStringSubmatch(respString)

	if len(match) == 0 {
		fmt.Println("Regex not matching respString", respString)
		return ""
	}
	return match[1]
}

func download(url string) []byte {
	fmt.Println("Downloading...", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("ERR getting", url)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Status not OK")
		return nil
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERR reading", url)
		return nil
	}

	return respBytes
}