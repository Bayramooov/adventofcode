package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isFullyOverlap(x []int) bool {
	return x[0] <= x[2] && x[1] >= x[3] || x[0] >= x[2] && x[1] <= x[3]
}

func split(str string) []int {
	r := make([]int, 4)

	x := strings.Split(str, ",")
	p1, p2 := x[0], x[1]

	x = strings.Split(p1, "-")
	r[0], _ = strconv.Atoi(x[0])
	r[1], _ = strconv.Atoi(x[1])

	x = strings.Split(p2, "-")
	r[2], _ = strconv.Atoi(x[0])
	r[3], _ = strconv.Atoi(x[1])

	return r
}

func main() {
	scanner := scan("./input.txt")
	/************************************************************

	  Main logic is written below
	  scanner.Scan() - scans a new line
	  scanner.Text() - returns a new scanned line
	  scanner.Err()  - returns an occured error while scanning


	************************************************************/
	counter := 0

	for scanner.Scan() {
		if isFullyOverlap(split(scanner.Text())) {
			counter++
		}
	}
	check(scanner.Err())

	fmt.Println(counter) // output: 657
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
