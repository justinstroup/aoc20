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

type void struct{}

var empty void

// type Set map[string]void

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func findLongestContiguousSubset(data []int, target int) []int {
	var sum = 0
	var indicesOfLongestSs = []int{0, 0}

	for i, j := 0, 0; i < len(data); i++ {
		if data[i] > target {
			j = i + 1
			sum = 0
			continue
		}
		sum += data[i]
		// printf("sum: %d\n", sum)
		// printf("i: %d\tj: %d\n", i, j)
		if sum == target {
			if i-j > indicesOfLongestSs[1]-indicesOfLongestSs[0] {
				indicesOfLongestSs[0], indicesOfLongestSs[1] = j, i
			}
		} else if sum > target {
			for {
				sum -= data[j]
				j++
				if sum == target {
					if i-j > indicesOfLongestSs[1]-indicesOfLongestSs[0] {
						indicesOfLongestSs[0], indicesOfLongestSs[1] = j, i
					}
				} else if sum < target || j == i-1 {
					break
				}
			}
		}
	}

	return data[indicesOfLongestSs[0]:indicesOfLongestSs[1]]
}

func main() {
	defer writer.Flush()

	var preamble []int
	var num int

	for {
		if _, err := scanf("%d\n", &num); err == io.EOF {
			break
		} else {
			preamble = append(preamble, num)
		}
	}

	var longestContiguousSubset = findLongestContiguousSubset(preamble, 1124361034)
	var min, max = longestContiguousSubset[0], longestContiguousSubset[1]
	for _, v := range longestContiguousSubset {
		min = Min(min, v)
		max = Max(max, v)
	}

	printf("answer: %d\n", min+max)
}
