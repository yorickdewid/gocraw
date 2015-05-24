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

func HandleRequest(req string) {
	fmt.Println("Request: " + req)
	html := Webrequest(req)
	OutName := Makefilename(req) + ".txt"
	SaveFile(OutName, html)
}

func main() {
	file, err := os.Open("gocraw.conf")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	r, _ := regexp.Compile("^https?://(www.)?[a-zA-Z0-9.]{2,512}.[a-z]{2,10}/?$")
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		if line[0:1] == "#" {
			continue
		}

		if r.MatchString(line) {
			go HandleRequest(line)
		}
	}

	var input string
	fmt.Scanln(&input)
	fmt.Println("Done")
}
