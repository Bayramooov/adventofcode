package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseReverseSplit(containers string) []string {
	result := make([]string, len(containers), 100)
	chrIndex := 0
	for i := len(containers) - 1; i >= 0; i-- {
		result[i] = string(containers[chrIndex])
		chrIndex++
	}
	return result
}

func parseIntSplit(moves string) []int {
	movesArr := strings.Split(moves, "-")
	result := make([]int, len(movesArr))
	for i := 0; i < len(movesArr); i++ {
		result[i], _ = strconv.Atoi(movesArr[i])
	}
	return result
}

func main() {
	initScan := scan("./init.txt")
	scanner := scan("./input.txt")
	/************************************************************

	  Main logic is written below
	  scanner.Scan() - scans a new line
	  scanner.Text() - returns a new scanned line
	  scanner.Err()  - returns an occured error while scanning


	************************************************************/
	containers := make([][]string, 9)
	i := 0
	for initScan.Scan() {
		containers[i] = parseReverseSplit(initScan.Text())
		i++
	}
	check(initScan.Err())

	for scanner.Scan() {
		moves := parseIntSplit(scanner.Text())
		from := moves[1] - 1
		to := moves[2] - 1
		containers[to] = append(containers[to], containers[from][len(containers[from])-moves[0]:]...)
		containers[from] = containers[from][:len(containers[from])-moves[0]]
	}
	check(scanner.Err())

	for _, val := range containers {
		fmt.Print(val[len(val)-1])
	}
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
