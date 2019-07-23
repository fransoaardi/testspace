package main

import (
	"bufio"
	"fmt"
	"os"
)

func usingNewReader() {
	file, err := os.Open("./file_read/sample_read.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}
	reader := bufio.NewReader(file)

	a, err := reader.Peek(1000) // peek length to read
	fmt.Println(string(a))
	fmt.Println("\n\n END bufio.NewReader\n")
}

func usingNewScanner() {
	file, err := os.Open("./file_read/sample_read.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	fmt.Println("\n END bufio.NewScanner")
}

func main() {
	usingNewReader()
	usingNewScanner()
}
