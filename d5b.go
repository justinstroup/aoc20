package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

func findSeat(sortedListOfSeats []int) int {
	for i := 1; i < len(sortedListOfSeats); i++ {
		if sortedListOfSeats[i]-sortedListOfSeats[i-1] == 2 {
			fmt.Println("%v - %v", sortedListOfSeats[i-1], sortedListOfSeats[i])
			return sortedListOfSeats[i-1] + 1
		}
	}
	return -1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	var lines []string
	var seatIDs []int

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := 0; i < len(lines); i++ {
		row := findRow(lines[i][0:7])
		col := findCol(lines[i][7:10])

		seatID := (row * 8) + col
		seatIDs = append(seatIDs, seatID)
	}

	sort.Ints(seatIDs)
	answer := findSeat(seatIDs)
	if answer == -1 {
		fmt.Println("Seat ID not found")
	} else {
		fmt.Println("answer: ", answer)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
