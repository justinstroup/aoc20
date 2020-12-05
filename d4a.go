package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// type PassportField string

const (
	BirthYear      = "byr"
	IssueYear      = "iyr"
	ExpirationYear = "eyr"
	Height         = "hgt"
	HairColor      = "hcl"
	EyeColor       = "ecl"
	PassportID     = "pid"
	CountryID      = "cid"
)

const minOfRequiredPassportFields = 7

func isValidRequiredPassportField(field string) bool {
	switch field {
	case BirthYear, EyeColor, IssueYear, ExpirationYear, Height, HairColor,
		PassportID:
		return true
	default:
		return false
	}
}

func isValidOptionalPassportField(field string) bool {
	switch field {
	case CountryID:
		return true
	default:
		return false
	}
}

func isValidPassport(passport string) bool {
	var fields = strings.Fields(passport)
	var lenOfFields = len(fields)
	if lenOfFields < minOfRequiredPassportFields {
		return false
	}

	numOfValidFields := 0
	for i := 0; i < lenOfFields; i++ {
		field := strings.Split(fields[i], ":")[0]
		if isValidRequiredPassportField(field) {
			numOfValidFields++
		} else {
			if isValidOptionalPassportField(field) {
				continue
			}
			return false
		}
	}

	return numOfValidFields >= minOfRequiredPassportFields
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	var lines []string

	var ph string

	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			ph += text + " "
		} else {
			lines = append(lines, strings.TrimSpace(ph))
			ph = ""
		}
	}

	var count = 0
	for i := 0; i < len(lines); i++ {
		if isValidPassport(lines[i]) {
			count++
		}
	}

	fmt.Println("count: ", count)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
