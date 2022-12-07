package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type file struct {
	size int
	name string
}

type dir struct {
	name   string
	size   int
	files  map[string]*file
	dirs   map[string]*dir
	parent *dir
}

func (d *dir) calc_size() {
	files_size := 0
	dirs_size := 0
	for _, file := range d.files {
		files_size += file.size
	}
	for _, dir := range d.dirs {
		if dir.size == -1 {
			dir.calc_size()
		}
		dirs_size += dir.size
	}
	d.size = files_size + dirs_size
}

func (d *dir) sum_size() int {
	dirs_size := 0

	for _, dir := range d.dirs {
		if dir.size <= 100000 {
			dirs_size += dir.size
		}
		dirs_size += dir.sum_size()
	}

	return dirs_size
}

func (d *dir) print_dir_sizes() {
	if len(d.dirs) == 0 {
		fmt.Println(d.size)
	} else {
		for _, dir := range d.dirs {
			dir.print_dir_sizes()
		}
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
	var fileServer dir
	fileServer.name = "/"
	fileServer.size = -1
	fileServer.files = make(map[string]*file)
	fileServer.dirs = make(map[string]*dir)

	var curr_dir *dir

	for scanner.Scan() {
		/*
			$ cd dirname
			$ cd ..
			$ cd /
			$ ls
			dir dirname
			fileSize filename
		*/

		input := scanner.Text()

		is_change_to_home, _ := regexp.MatchString(`^\$ cd (\/)$`, input)
		is_change_back, _ := regexp.MatchString(`^\$ cd (..)$`, input)
		is_change_into, _ := regexp.MatchString(`^\$ cd (\w+)$`, input)
		is_list, _ := regexp.MatchString(`^\$ ls$`, input)
		is_dir, _ := regexp.MatchString(`^dir (\w+)$`, input)
		is_file, _ := regexp.MatchString(`^\d+ (\w+\.*\w+)$`, input)

		if is_change_to_home {
			curr_dir = &fileServer
			// fmt.Println("is_change_to_home")

		} else if is_change_back {
			curr_dir = curr_dir.parent
			// fmt.Println("is_change_back")

		} else if is_change_into {
			curr_dir = curr_dir.dirs[input[5:]]
			// fmt.Println("is_change_into", input[5:])

		} else if is_list {
			// fmt.Println("is_list")

		} else if is_dir {
			var new_dir dir
			new_dir.name = input[4:]
			new_dir.size = -1
			new_dir.files = make(map[string]*file)
			new_dir.dirs = make(map[string]*dir)
			new_dir.parent = curr_dir

			// fmt.Println("is_dir", new_dir.name)
			curr_dir.dirs[new_dir.name] = &new_dir

		} else if is_file {
			x := strings.Split(input, " ")
			filesize, _ := strconv.Atoi(x[0])
			filename := x[1]

			var new_file file
			new_file.size = filesize
			new_file.name = filename

			// fmt.Println("is_file", filesize, filename)
			curr_dir.files[new_file.name] = &new_file
		}
	}
	check(scanner.Err())

	fileServer.calc_size()

	fmt.Println(fileServer.sum_size())
	// fileServer.print_dir_sizes()

	// output: LJSVLTWQM
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
