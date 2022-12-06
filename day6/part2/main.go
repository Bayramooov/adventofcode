package main

import (
	"bufio"
	"fmt"
	"os"
)

func isUniqueChars(str string) bool {
	for i := 0; i < len(str)-1; i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				return false
			}
		}
	}
	return true
}

func main() {
	scanner := scan("./input.txt")
	/************************************************************

	  Main logic is written below
	  scanner.Scan() - scans a new line
	  scanner.Text() - returns a new scanned line
	  scanner.Err()  - returns an occured error while scanning


	************************************************************/
	for scanner.Scan() {
		input := scanner.Text()

		for i := 0; i < len(input); i++ {
			if isUniqueChars(input[i : i+14]) {
				fmt.Println(i + 14)
				break
			}
		}

	}
	check(scanner.Err())

	// output: LJSVLTWQM
}

func scan(path string) *bufio.Scanner {
	file, err := os.Open(path)
	check(err)
	return bufio.NewScanner(file)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
