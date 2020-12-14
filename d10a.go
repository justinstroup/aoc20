package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func scanf(f string, a ...interface{}) (n int, err error) { return fmt.Fscanf(reader, f, a...) }
func printf(f string, a ...interface{})                   { fmt.Fprintf(writer, f, a...) }

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	defer writer.Flush()

	var adapters []int
	var n int
	for {
		if _, err := scanf("%d\n", &n); err == io.EOF {
			break
		}

		adapters = append(adapters, n)
	}

	sort.Ints(adapters)
	var oneJoltDiff, threeJoltDiff = 0, 1
	var prev = 0
	for _, adapter := range adapters {
		// printf("adapter: %d\n", adapter)
		diff := Abs(prev - adapter)
		if diff == 1 {
			oneJoltDiff++
		} else if diff == 3 {
			threeJoltDiff++
		}

		prev = adapter
	}

	printf("answer: %d\n", oneJoltDiff*threeJoltDiff)
}
