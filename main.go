package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	)

func main() {
	file, err := os.Open("targetlist.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	r, _ := regexp.Compile("https?://(www.)?[a-zA-Z0-9.]{2,512}.[a-z]{2,10}")

	for scanner.Scan() {
		line := scanner.Text()

		if r.MatchString(line) {
			fmt.Println("Valid: " + line)
		}
	}
}
