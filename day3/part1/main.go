package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		letter := findEqualLetter(scanner.Text())
		/*
			a - z => 1 - 26 ----- 97 - 122
			A - Z => 27 - 52 ----- 65 - 90
		*/
		if letter >= 97 {
			sum += letter - 96
		} else if letter >= 65 {
			sum += letter - 38
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}

func findEqualLetter(str string) int {
	for i := 0; i < len(str)/2; i++ {
		for j := len(str) / 2; j < len(str); j++ {
			if str[i] == str[j] {
				return int(str[i])
			}
		}
	}
	return 0
}
