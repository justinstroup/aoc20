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

func main() {
	defer writer.Flush()

	lenOfPreamble := 25
	var preamble []int
	var num int
	var isValid = false
	i := 0
	for {
		if _, err := scanf("%d\n", &num); err == io.EOF {
			break
		} else {
			preamble = append(preamble, num)
		}

		if i >= lenOfPreamble+1 {
			memo := make(map[int]int)
			target := preamble[i-1]
			// printf("len of preamble: %d\ti: %d\n", i-1, i)
			for j := i - 1 - lenOfPreamble; j < i; j++ {
				// printf("target: %d\tj: %d\t%d\n", target, preamble[j], j)
				if _, doesExist := memo[preamble[j]]; doesExist {
					isValid = true
					break
				} else {
					memo[target-preamble[j]] = preamble[j]
				}
			}

			if !isValid {
				printf("%d is not valid.\n", target)
				break
			}
			isValid = false
			memo = nil
		}

		i++
	}
}
