package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Coordinates struct {
	X int
	Y int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	var grid [][]byte
	i := 0
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
		i++
	}

	var numOfTrees = 0
	var pos = Coordinates{0, 0}
	var lenOfRow = len(grid[0])
	for i := 1; i < len(grid); i++ {
		pos.X += 3
		pos.Y += 1

		if grid[pos.Y][pos.X%lenOfRow] == '#' {
			numOfTrees++
		}
	}

	fmt.Println("Number of trees:", numOfTrees)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
