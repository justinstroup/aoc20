package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isValid(pass string, min int, max int, c byte) bool {
	if min == 0 || max == 0 {
		return false
	}

	return (pass[min-1] != c) != (pass[max-1] != c)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var count = 0
	for i := 0; i < len(lines); i++ {
		var values = strings.Fields(lines[i])
		var numOfOccurences = strings.Split(values[0], "-")
		min, err := strconv.Atoi(numOfOccurences[0])
		max, err := strconv.Atoi(numOfOccurences[1])
		if err != nil {
			log.Fatal(err)
		}

		var c = values[1]
		c = c[:len(c)-1]
		var pass = values[2]
		if isValid(pass, min, max, c[0]) {
			count++
		}
	}

	fmt.Println("count: ", count)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
