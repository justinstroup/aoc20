package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const numOfRows = 128
const numOfCols = 8

func findRow(seat string) int {
	// FBFBBFFRLR
	row := 0
	low, high := 0, 127
	for i := 0; i < len(seat); i++ {
		if seat[i] == 'F' {
			high = (high + low) / 2
			row = low
		} else {
			low = ((high + low) / 2) + 1
			row = high
		}
	}
	return row
}

func findCol(seat string) int {
	col := 0
	low, high := 0, 7
	for i := 0; i < len(seat); i++ {
		if seat[i] == 'L' {
			high = ((high + low) / 2)
			col = low
		} else {
			low = ((high + low) / 2) + 1
			col = high
		}
	}
	return col
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	var lines []string
	var seatIDs []int

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	max := 0
	for i := 0; i < len(lines); i++ {
		row := findRow(lines[i][0:7])
		col := findCol(lines[i][7:10])

		seatID := (row * 8) + col
		if seatID > max {
			max = seatID
		}
		seatIDs = append(seatIDs, findRow(lines[i][0:7]))
	}

	fmt.Println("answer: ", max)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
