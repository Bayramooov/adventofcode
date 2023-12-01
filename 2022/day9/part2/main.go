package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type rope struct {
	knots []*knot
}

func (r *rope) start(len_knot int, i_index int, j_index int) {
	knots := make([]*knot, len_knot)
	knots[0] = &knot{i: i_index, j: j_index}
	for i := 1; i < len_knot; i++ {
		knot := knot{i: i_index, j: j_index, prev_knot: knots[i-1]}
		knots[i] = &knot
		knots[i-1].next_knot = &knot
	}
	r.knots = knots
}

type knot struct {
	i         int   // -y
	j         int   // x
	prev_knot *knot // near to head
	next_knot *knot // near to tail
}

func (k *knot) chase() {
	if k.prev_knot != nil {
		i_diff := k.prev_knot.i - k.i
		j_diff := k.prev_knot.j - k.j
		switch {
		case i_diff == -2 && j_diff == 0:
			k.i--
		case i_diff == 0 && j_diff == 2:
			k.j++
		case i_diff == 2 && j_diff == 0:
			k.i++
		case i_diff == 0 && j_diff == -2:
			k.j--

		case i_diff == -2 && j_diff == 1 || i_diff == -2 && j_diff == 2 || i_diff == -1 && j_diff == 2:
			k.i--
			k.j++
		case i_diff == 1 && j_diff == 2 || i_diff == 2 && j_diff == 2 || i_diff == 2 && j_diff == 1:
			k.i++
			k.j++
		case i_diff == 2 && j_diff == -1 || i_diff == 2 && j_diff == -2 || i_diff == 1 && j_diff == -2:
			k.i++
			k.j--
		case i_diff == -1 && j_diff == -2 || i_diff == -2 && j_diff == -2 || i_diff == -2 && j_diff == -1:
			k.i--
			k.j--
		}
	}

	if k.next_knot != nil {
		k.next_knot.chase()
	}
}

func main() {
	scanner := scan("./input.txt")
	file, err := os.Create("./data2.txt")
	check(err)
	/************************************************************

	  Main logic is written below
	  scanner.Scan() - scans a new line
	  scanner.Text() - returns a new scanned line
	  scanner.Err()  - returns an occured error while scanning

	************************************************************/
	grid := make([][]string, 1000)
	for i, _ := range grid {
		grid[i] = make([]string, 1000)
		for j, _ := range grid[i] {
			grid[i][j] = "."
		}
	}

	var r rope
	// r.start(10, 15, 11)
	r.start(10, 500, 500)
	head := r.knots[0]
	tail := r.knots[len(r.knots)-1]

	grid[tail.i][tail.j] = "#"

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		move := input[0]
		step, _ := strconv.Atoi(input[1])

		for i := 0; i < step; i++ {
			switch move {
			case "U":
				head.i--
			case "R":
				head.j++
			case "D":
				head.i++
			case "L":
				head.j--
			}
			head.chase()
			grid[head.i][head.j] = "1"
			// grid[tail.i][tail.j] = "#"
		}
	}
	check(scanner.Err())

	sum := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// fmt.Print(grid[i][j])
			file.WriteString(grid[i][j])

			if grid[i][j] == "#" {
				sum++
			}
		}
		// fmt.Println()
		file.WriteString("\n")
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
	}
}
