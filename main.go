package main

import (
	"bufio"
	"fmt"
	"os"
	)

func main() {
	file, err := os.Open("targetlist.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
