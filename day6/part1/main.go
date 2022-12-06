package main

import (
	"bufio"
	"fmt"
	"os"
)

func isUnique4Chars(str string) bool {
	return str[0] != str[1] &&
		str[0] != str[2] &&
		str[0] != str[3] &&
		str[1] != str[2] &&
		str[1] != str[3] &&
		str[2] != str[3]
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
			if isUnique4Chars(input[i : i+4]) {
				fmt.Println(i + 4)
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
