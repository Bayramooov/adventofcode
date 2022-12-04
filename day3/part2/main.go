package main

import (
	"bufio"
	"fmt"
	"os"
)

func findEqualLetter(str []string) int {
	for i := 0; i < len(str[0]); i++ {
		for j := 0; j < len(str[1]); j++ {
			for k := 0; k < len(str[2]); k++ {
				if str[0][i] == str[1][j] && str[1][j] == str[2][k] {
					return int(str[0][i])
				}
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
	i := 0
	badges := make([]string, 3)
	sum := 0

	for scanner.Scan() {
		badges[i] = scanner.Text()
		i++
		if i > 2 {
			letter := findEqualLetter(badges)
			//	a - z =>  1 - 26 (97 - 122)
			//	A - Z => 27 - 52 (65 - 90)
			if letter >= 97 {
				sum += letter - 96
			} else if letter >= 65 {
				sum += letter - 38
			}
			i = 0
			badges = make([]string, 3)
		}
	}
	check(scanner.Err())

	fmt.Println(sum) // output: 2703
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
