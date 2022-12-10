package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// ////////////////////////////////////////////////
type Pixel struct {
	cell string
}

func (p *Pixel) on() {
	p.cell = "#"
}

func (p *Pixel) off() {
	p.cell = "."
}

// ////////////////////////////////////////////////
type Monitor struct {
	output  [][]*Pixel
	height  int
	width   int
	i_index int
	j_index int
	cycle   int
}

func (m *Monitor) init(height int, width int) {
	m.output = make([][]*Pixel, height)
	m.height = height
	m.width = width
	m.i_index = 0
	m.j_index = 0
	m.cycle = 0

	for i := 0; i < height; i++ {
		m.output[i] = make([]*Pixel, width)

		for j := 0; j < width; j++ {
			var px Pixel
			px.off()
			m.output[i][j] = &px
		}
	}
}

func (m *Monitor) write(s *Sprite) {
	m.output[m.i_index][m.j_index].cell = s.pixels[m.j_index].cell
}

func (m *Monitor) write_cycle(s *Sprite) {
	m.write(s)
	if m.width-1 == m.j_index {
		if m.height-1 == m.i_index {
			panic("monitor cycle write overflow")
		}
		m.j_index = 0
		m.i_index++

	} else {
		m.j_index++
	}
	m.cycle++
}

func (m *Monitor) print() {
	for i := 0; i < m.height; i++ {
		for j := 0; j < m.width; j++ {
			fmt.Print(m.output[i][j].cell)
		}
		fmt.Println()
	}
}

// ////////////////////////////////////////////////
type Sprite struct {
	registerX int
	pixels    []*Pixel
	width     int
}

func (s *Sprite) init(length int, width int) {
	s.registerX = 1
	s.pixels = make([]*Pixel, length)
	s.width = width

	for i := 0; i < length; i++ {
		var px Pixel
		if i < width {
			px.on()
		} else {
			px.off()
		}
		s.pixels[i] = &px
	}
}

func (s *Sprite) setPos(x int) {
	for i := 0; i < len(s.pixels); i++ {
		if i >= x-s.width/2 && i < x-s.width/2+s.width {
			s.pixels[i].on()
		} else {
			s.pixels[i].off()
		}
	}
}

func (s *Sprite) addX(x int) {
	s.registerX += x
	s.setPos(s.registerX)
}

func (s *Sprite) posVal(x int) string {
	return s.pixels[x].cell
}

// ////////////////////////////////////////////////
func main() {
	scanner := scan("./input.txt")
	/************************************************************

	  Main logic is written below
	  scanner.Scan() - scans a new line
	  scanner.Text() - returns a new scanned line
	  scanner.Err()  - returns an occured error while scanning


	************************************************************/
	const monitor_height = 7
	const monitor_width = 40
	const sprite_width = 3

	var monitor Monitor
	monitor.init(monitor_height, monitor_width)

	var sprite Sprite
	sprite.init(monitor_width, sprite_width)

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		is_pause := len(input) == 1
		value := 0
		if !is_pause {
			value, _ = strconv.Atoi(input[1])
		}

		if is_pause {
			monitor.write_cycle(&sprite)
		} else {
			monitor.write_cycle(&sprite)
			monitor.write_cycle(&sprite)
			sprite.addX(value)
		}
	}
	check(scanner.Err())

	// fmt.Println(len(monitor.output))
	monitor.print()
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
