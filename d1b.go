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

func threeSum(A []int, sum int) (a int, b int, c int) {
	memo := make(map[int]int)

	for i := 0; i < len(A); i++ {
		diff := sum - A[i]
		if diff < 0 {
			continue
		}

		for j := i + 1; j < len(A); j++ {
			if v, doesExist := memo[A[j]]; doesExist {
				return A[i], v, A[j]
			} else {
				if A[j] <= diff {
					memo[diff-A[j]] = A[j]
				}
			}
		}
	}

	return
}

func main() {
	defer writer.Flush()

	sum := 2020
	var n int
	var A []int
	for {
		if _, err := scanf("%d\n", &n); err == io.EOF {
			break
		}

		A = append(A, n)
	}

	a, b, c := threeSum(A, sum)
	printf("%d\n", a*b*c)
}
