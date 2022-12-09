package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	i int
	j int
}

func yetvol(h *point, t *point) {
	i_diff := h.i - t.i
	j_diff := h.j - t.j

	switch {
	case i_diff == -2 && j_diff == 0:
		t.i--
	case i_diff == 0 && j_diff == 2:
		t.j++
	case i_diff == 2 && j_diff == 0:
		t.i++
	case i_diff == 0 && j_diff == -2:
		t.j--

	case i_diff == -2 && j_diff == 1 || i_diff == -1 && j_diff == 2:
		t.i--
		t.j++
	case i_diff == 1 && j_diff == 2 || i_diff == 2 && j_diff == 1:
		t.i++
		t.j++
	case i_diff == 2 && j_diff == -1 || i_diff == 1 && j_diff == -2:
		t.i++
		t.j--
	case i_diff == -1 && j_diff == -2 || i_diff == -2 && j_diff == -1:
		t.i--
		t.j--
	}
}

func main() {
	scanner := scan("./input.txt")
	/************************************************************

	  Main logic is written below
	  scanner.Scan() - scans a new line
	  scanner.Text() - returns a new scanned line
	  scanner.Err()  - returns an occured error while scanning


	************************************************************/
	grid := make([][]int, 10000)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, 10000)
	}

	sum := 0

	s := point{i: 5000, j: 5000}

	h := point{i: s.i, j: s.j}
	t := point{i: s.i, j: s.j}

	grid[s.i][s.j] = 1

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		move := input[0]
		step, _ := strconv.Atoi(input[1])

		for i := 0; i < step; i++ {
			switch move {
			case "U":
				h.i--
			case "R":
				h.j++
			case "D":
				h.i++
			case "L":
				h.j--
			}
			yetvol(&h, &t)
			// grid[h.i][h.j] = 1
			grid[t.i][t.j] = 1
		}
	}
	check(scanner.Err())

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			sum += grid[i][j]
			// fmt.Print(grid[i][j])
		}
		// fmt.Println()
	}

	fmt.Println(sum)
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
