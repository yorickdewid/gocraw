package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"net/http"
	"io/ioutil"
	"strings"
	"flag"
	"sync"
)


var p = fmt.Println

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

func HandleRequest(wg *sync.WaitGroup, req string) {
	defer wg.Done()
	p("Request: " + req)
	html := Webrequest(req)
	OutName := Makefilename(req) + ".txt"
	SaveFile(OutName, html)
}

func OpenConfig(File string) *os.File {
	file, err := os.Open(File)
	if err != nil {
		panic(err)
	}

	return file
}

func main() {

	if len(os.Args) < 2 {
		p("Missing input")
		os.Exit(1)
	}

	filename := flag.String("file", "filename", "Read targets from config file")

	flag.Parse()

	r, _ := regexp.Compile("^https?://(www.)?[a-zA-Z0-9.]{2,512}.[a-z]{2,10}/?$")

	if *filename != "filename" {
		var wg sync.WaitGroup
		conf := OpenConfig(*filename)

		defer conf.Close()

		scanner := bufio.NewScanner(conf)
		for scanner.Scan() {
			line := scanner.Text()

			if len(line) == 0 {
				continue
			}

			if line[0:1] == "#" {
				continue
			}

			if r.MatchString(line) {
				wg.Add(1)
				go HandleRequest(&wg, line)
			}
		}
		wg.Wait()
	} else {
		Url := os.Args[1:][0]
		if r.MatchString(Url) {
			p("Request: " + Url)
			html := Webrequest(Url)
			OutName := Makefilename(Url) + ".txt"
			SaveFile(OutName, html)
		} else {
			p("Not a valid URL")
			os.Exit(0)
		}
	}
}
