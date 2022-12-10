package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

registerX := 1
cycle := 0
strength := 0

func calc_strength(20)

func main() {
	scanner := scan("./input.txt")
	/************************************************************

	  Main logic is written below
	  scanner.Scan() - scans a new line
	  scanner.Text() - returns a new scanned line
	  scanner.Err()  - returns an occured error while scanning


	************************************************************/
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")

		is_pause := len(input) == 1
		value := 0
		if !is_pause {
			value, _ = strconv.Atoi(input[1])
		}

		if is_pause {
			cycle++
			continue
		}


	}
	check(scanner.Err())
}

func scan(path string) *bufio.Scanner {
	file, err := os.Open(path)
	check(err)
	return bufio.NewScanner(file)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
