package main

import (
	"bufio"
	"fmt"
	"os"
)

func findEqualLetter(str string) int {
	for i := 0; i < len(str)/2; i++ {
		for j := len(str) / 2; j < len(str); j++ {
			if str[i] == str[j] {
				return int(str[i])
			}
		}
	}
	return 0
}

func main() {
	scanner := scan("./input.txt")
	/************************************************************

		Main logic is written below
		scanner.Scan() - scans a new line
		scanner.Text() - returns a new scanned line
		scanner.Err()  - returns an occured error while scanning


	************************************************************/
	sum := 0

	for scanner.Scan() {
		letter := findEqualLetter(scanner.Text())
		//	a - z =>  1 - 26 (97 - 122)
		//	A - Z => 27 - 52 (65 - 90)
		if letter >= 97 {
			sum += letter - 96
		} else if letter >= 65 {
			sum += letter - 38
		}
	}
	check(scanner.Err())

	fmt.Println(sum) // output: 7795
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
