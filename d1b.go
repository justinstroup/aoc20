package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func threeSum(lines []string) int {
	memo := make(map[int]int)

	for i := 0; i < len(lines); i++ {
		var curr, err = strconv.Atoi(lines[i])
		if err != nil {
			log.Fatal(err)
		}
		var a = 2020 - curr

		for j := i + 1; j < len(lines); j++ {
			if a < 0 {
				break
			}

			var num, err = strconv.Atoi(lines[j])
			if err != nil {
				log.Fatal(err)
			}

			if memo[num] != 0 {
				return memo[num] * num * curr
			} else {
				var diff = a - num
				if diff >= 0 {
					memo[diff] = num
				}
			}
		}
	}

	return -1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var answer = threeSum(lines)
	fmt.Println("Answer:", answer)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
