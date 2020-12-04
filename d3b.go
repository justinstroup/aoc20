package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Coordinates struct {
	X int "default:0"
	Y int `default:0`
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	var grid [][]byte
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}

	const numOfSlopes = 5
	var numOfTrees [numOfSlopes]int
	var slopes = [numOfSlopes][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	var positions [numOfSlopes]Coordinates
	var lenOfRow = len(grid[0])
	// var lenOfCol = len(grid)
	for i := 1; i < len(grid); i++ {
		for j := 0; j < numOfSlopes; j++ {
			positions[j].X = positions[j].X + slopes[j][0]
			positions[j].Y = positions[j].Y + slopes[j][1]
			if positions[j].Y > len(grid) {
				break
			} else if grid[positions[j].Y][positions[j].X%lenOfRow] == '#' {
				numOfTrees[j]++
			}
		}
	}

	var count = numOfTrees[0]
	for i := 1; i < len(numOfTrees); i++ {
		count = count * numOfTrees[i]
	}
	fmt.Println("Number of trees:", count)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
