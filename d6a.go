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

	set := make(map[byte]void)
	var A []map[byte]void
	var s string
	for {
		if n, err := scanf("%s\n", &s); err == io.EOF {
			A = append(A, set)
			break
		} else {
			if n != 0 {
				for i := 0; i < len(s); i++ {
					set[s[i]] = empty
				}
			} else {
				A = append(A, set)
				set = nil
				set = make(map[byte]void)
			}
		}
	}

	var count = 0
	for _, v := range A {
		count += len(v)
	}

	printf("count: %d\n", count)
}
