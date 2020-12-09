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

	set := make(map[byte]int)
	var numOfPeople = 0
	var count = 0
	var s string
	for {
		if n, err := scanf("%s\n", &s); err == io.EOF {
			for _, v := range set {
				if v == numOfPeople {
					count++
				}
			}
			break
		} else {
			if n != 0 {
				numOfPeople++
				for i := 0; i < len(s); i++ {
					if v, doesExist := set[s[i]]; doesExist {
						set[s[i]] = v + 1
					} else {
						set[s[i]] = 1
					}
				}
			} else {
				for _, v := range set {
					if v == numOfPeople {
						count++
					}
				}
				numOfPeople = 0
				set = nil
				set = make(map[byte]int)
			}
		}
	}

	printf("count: %d\n", count)
}
