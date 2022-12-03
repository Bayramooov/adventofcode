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

	counter := 0
	badges := make([]string, 3)
	sum := 0

	cnt := 0

	for scanner.Scan() {
		badges[counter] = scanner.Text()
		counter++
		if counter > 2 {
			cnt++
			letter := findEqualLetter(badges)
			/*
				a - z => 1 - 26 ----- 97 - 122
				A - Z => 27 - 52 ----- 65 - 90
			*/
			fmt.Println(badges)
			if letter >= 97 {
				sum += letter - 96
			} else if letter >= 65 {
				sum += letter - 38
			}
			counter = 0
			badges = make([]string, 3)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
	fmt.Println(cnt)
}

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
