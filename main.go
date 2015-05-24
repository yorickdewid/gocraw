package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"net/http"
	"io/ioutil"
	"strings"
	)


func check(e error) {
	if e != nil {
    	panic(e)
	}
}

// Request webcontent from url
func Webrequest(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	return string(body)
}

// Write content to file
func SaveFile(File string, ctx string) {
	d1 := []byte(ctx)
	err := ioutil.WriteFile(File, d1, 0644)
	check(err)
}

// Substract name from URL
func Makefilename(URL string) string {
	usz := len(URL)

	if URL[usz-1] == '/' {
		URL = URL[0:usz-1]
	}

	protpos := strings.Index(URL, "//")
	URL = URL[protpos+2:len(URL)]

	return strings.Replace(URL, ".", "_", -1)
}

func main() {
	file, err := os.Open("targetlist.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	r, _ := regexp.Compile("^https?://(www.)?[a-zA-Z0-9.]{2,512}.[a-z]{2,10}/?$")
	for scanner.Scan() {
		line := scanner.Text()

		if r.MatchString(line) {
			fmt.Println("Valid: " + line)
			html := Webrequest(line)
			OutName := Makefilename(line) + ".txt"
			SaveFile(OutName, html)
		}
	}
}
