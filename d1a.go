package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	memo := make(map[int]int)

	scanner := bufio.NewScanner(os.Stdin)
	var count = 1
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		var num, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if memo[num] != 0 {
			fmt.Println("Answer:", memo[num]*num)
			fmt.Println("Count:", count)
		} else {
			var diff = 2020 - num
			if diff > -1 {
				memo[diff] = num
			}
		}

		count++
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
