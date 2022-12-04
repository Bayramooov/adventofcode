package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	counter := 0

	for scanner.Scan() {
		if isFullyOverlap(split(scanner.Text())) {
			fmt.Println(scanner.Text())
			counter++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(counter)
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

func isFullyOverlap(x []int) bool {
	return x[0] <= x[2] && x[1] >= x[3] || x[0] >= x[2] && x[1] <= x[3]
}
