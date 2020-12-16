package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func scanf(f string, a ...interface{}) (n int, err error) { return fmt.Fscanf(reader, f, a...) }
func printf(f string, a ...interface{})                   { fmt.Fprintf(writer, f, a...) }

type Point struct {
	X int
	Y int
}

func BytesToString(data []byte) string {
	return string(data[:])
}

func deepCopyBytesArray(data [][]byte) [][]byte {
	var cpy [][]byte
	var temp []byte

	for i := 0; i < len(data); i++ {
		temp = make([]byte, len(data[i]))
		copy(temp, data[i])
		cpy = append(cpy, temp)
	}

	return cpy
}

func getNumOfAdjacentOccupiedSeats(grid [][]byte, p Point) int {
	var numOfAdjacentOccupiedSeats = 0

	// printf("%c: %d, %d\n", grid[p.Y][p.X], p.X, p.Y)
	// printf("\n")
	var left = p.X - 1
	var right = p.X + 1
	var up = p.Y - 1
	var down = p.Y + 1

	if left >= 0 {
		// Left
		if grid[p.Y][left] == '#' {
			numOfAdjacentOccupiedSeats++
		}
		// Left Up
		if up >= 0 && grid[up][left] == '#' {
			numOfAdjacentOccupiedSeats++
		}
		// Left Down
		if down < len(grid) && grid[down][left] == '#' {
			numOfAdjacentOccupiedSeats++
		}
	}

	if right < len(grid[p.Y]) {
		// Right
		if grid[p.Y][right] == '#' {
			numOfAdjacentOccupiedSeats++
		}
		// Right Up
		if up >= 0 && grid[up][right] == '#' {
			numOfAdjacentOccupiedSeats++
		}
		// Right Down
		if down < len(grid) && grid[down][right] == '#' {
			numOfAdjacentOccupiedSeats++
		}
	}

	// Up
	if up >= 0 && grid[up][p.X] == '#' {
		numOfAdjacentOccupiedSeats++
	}
	// Down
	if down < len(grid) && grid[down][p.X] == '#' {
		numOfAdjacentOccupiedSeats++
	}

	// printf("%d\n", numOfAdjacentOccupiedSeats)
	return numOfAdjacentOccupiedSeats
}

func main() {
	defer writer.Flush()

	var grid [][]byte
	var s string
	for {
		if _, err := scanf("%s\n", &s); err == io.EOF {
			break
		}

		grid = append(grid, []byte(s))
	}

	var numOfOccupiedSeats = 0
	var nextGrid = deepCopyBytesArray(grid)

	for didStateChange := true; didStateChange; {
		didStateChange = false
		numOfOccupiedSeats = 0
		for _, v := range grid {
			printf("%v\n", BytesToString(v))
		}
		printf("\n")

		var p = Point{0, 0}
		for i := 0; i < len(grid); i++ {
			p.Y = i
			for j := 0; j < len(grid[i]); j++ {
				p.X = j
				switch grid[i][j] {
				case 'L':
					if getNumOfAdjacentOccupiedSeats(grid, p) == 0 {
						nextGrid[i][j] = '#'
						didStateChange = true
						numOfOccupiedSeats++
					}
					break
				case '#':
					if getNumOfAdjacentOccupiedSeats(grid, p) >= 4 {
						nextGrid[i][j] = 'L'
						didStateChange = true
						break
					}

					numOfOccupiedSeats++
					break
				default:
					break
				}
			}
		}

		grid = deepCopyBytesArray(nextGrid)
	}

	printf("Number of occupied seats: %d\n", numOfOccupiedSeats)
}
