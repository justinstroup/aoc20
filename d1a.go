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

func main() {
	defer writer.Flush()

	sum := 2020
	memo := make(map[int]int)
	var n int
	for {
		if _, err := scanf("%d\n", &n); err == io.EOF {
			break
		}

		if v, doesExist := memo[n]; doesExist {
			printf("%d\n", v*n)
		} else {
			if n <= sum {
				memo[sum-n] = n
			}
		}
	}
}
