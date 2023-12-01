package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func scenic_view(trees [][]int, x int, y int) int {
	top := 0
	right := 0
	bottom := 0
	left := 0

	// top
	for i := x - 1; i >= 0; i-- {
		top++
		if trees[i][y] >= trees[x][y] {
			break
		}
	}

	// right
	for i := y + 1; i < len(trees[x]); i++ {
		right++
		if trees[x][i] >= trees[x][y] {
			break
		}
	}

	// bottom
	for i := x + 1; i < len(trees); i++ {
		bottom++
		if trees[i][y] >= trees[x][y] {
			break
		}
	}

	// left
	for i := y - 1; i >= 0; i-- {
		left++
		if trees[x][i] >= trees[x][y] {
			break
		}
	}

	return top * right * bottom * left
}

func main() {
	scanner := scan("./input.txt")
	/************************************************************

	  Main logic is written below
	  scanner.Scan() - scans a new line
	  scanner.Text() - returns a new scanned line
	  scanner.Err()  - returns an occured error while scanning


	************************************************************/
	trees := make([][]int, 0, 1000)

	for scanner.Scan() {
		row := make([]int, 0, 1000)
		for _, x := range strings.Split(scanner.Text(), "") {
			d, _ := strconv.Atoi(x)
			row = append(row, d)
		}
		trees = append(trees, row)
	}
	check(scanner.Err())

	scenic_scores := make([]int, 0, 10000)

	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[i]); j++ {
			scenic_scores = append(scenic_scores, scenic_view(trees, i, j))
		}
	}

	sort.Ints(scenic_scores)
	fmt.Println(scenic_scores[len(scenic_scores)-1])
}

func scan(path string) *bufio.Scanner {
	file, err := os.Open(path)
	check(err)
	return bufio.NewScanner(file)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
		// panic(err)
	}
}
