package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func is_visible(trees [][]int, x int, y int) bool {
	hidden := false

	// top
	for i := 0; i < x; i++ {
		if trees[i][y] >= trees[x][y] {
			hidden = true
			break
		}
	}

	if !hidden {
		return true
	}
	hidden = false

	// right
	for i := y + 1; i < len(trees[x]); i++ {
		if trees[x][i] >= trees[x][y] {
			hidden = true
			break
		}
	}

	if !hidden {
		return true
	}
	hidden = false

	// bottom
	for i := x + 1; i < len(trees); i++ {
		if trees[i][y] >= trees[x][y] {
			hidden = true
			break
		}
	}

	if !hidden {
		return true
	}
	hidden = false

	// left
	for i := 0; i < y; i++ {
		if trees[x][i] >= trees[x][y] {
			hidden = true
			break
		}
	}

	if !hidden {
		return true
	}

	return false
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
	visible_trees := 0

	for scanner.Scan() {
		row := make([]int, 0, 1000)
		for _, x := range strings.Split(scanner.Text(), "") {
			d, _ := strconv.Atoi(x)
			row = append(row, d)
		}
		trees = append(trees, row)
	}
	check(scanner.Err())

	visible_trees += (len(trees)+len(trees[0]))*2 - 4

	for i := 1; i < len(trees)-1; i++ {
		for j := 1; j < len(trees[i])-1; j++ {
			if is_visible(trees, i, j) {
				fmt.Println(i, j)
				visible_trees++
			}
		}
	}

	fmt.Println(visible_trees)
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
