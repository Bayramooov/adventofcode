package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func init_vertex(name string) *Vertex {
	var v Vertex
	v.name = name
	v.discovered = false
	v.edges = make([]*Edge, 0)
	return &v
}

type Vertex struct {
	name           string
	discovered     bool
	total_cost     int
	edges          []*Edge
	preferred_edge *Edge
}

func (v *Vertex) create_neighbour(edge_cost int, vertex_name string, is_duplex bool) {
	new_vertex := init_vertex(vertex_name)

	e := Edge{next_vertex: new_vertex, cost: edge_cost}
	v.edges = append(v.edges, &e)

	if is_duplex {
		new_vertex.edges = append(new_vertex.edges, &Edge{cost: e.cost, next_vertex: v})
	}
}

func (v *Vertex) link(vertex *Vertex, cost int, is_duplex bool) {
	if math.Abs(float64(int(v.name[0])-int(vertex.name[0]))) > 1 {
		return
	}
	e := Edge{cost: cost, next_vertex: vertex}
	v.edges = append(v.edges, &e)
	if is_duplex {
		e := Edge{cost: cost, next_vertex: v}
		vertex.edges = append(vertex.edges, &e)
	}
}

type Edge struct {
	cost        int
	next_vertex *Vertex
}

func main() {
	scanner := scan("./input.txt")
	/************************************************************

	  Main logic is written below
	  scanner.Scan() - scans a new line
	  scanner.Text() - returns a new scanned line
	  scanner.Err()  - returns an occured error while scanning


	************************************************************/
	grid := make([][]string, 0)
	start := make([]int, 2)
	end := make([]int, 2)

	for scanner.Scan() {
		grid = append(grid, strings.Split(scanner.Text(), ""))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "S" {
				start[0] = i
				start[1] = j

			} else if grid[i][j] == "E" {
				end[0] = i
				end[1] = j
			}
		}
	}
	check(scanner.Err())

	fmt.Println(grid)
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
