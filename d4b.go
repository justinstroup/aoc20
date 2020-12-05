package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

var eyeColors []string = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

const minOfRequiredPassportFields = 7

func stoi(s string) int {
	if num, err := strconv.Atoi(s); err == nil {
		return num
	}
	return -1
}

func contains(a []string, e string) bool {
	for _, s := range a {
		if s == e {
			return true
		}
	}

	return false
}

func isValidRequiredPassportField(name string, value string) bool {
	switch name {
	case BirthYear:
		num := stoi(value)
		return 1920 <= num && num <= 2002
	case IssueYear:
		num := stoi(value)
		return 2010 <= num && num <= 2020
	case ExpirationYear:
		num := stoi(value)
		return 2020 <= num && num <= 2030
	case Height:
		if strings.Contains(value, "cm") {
			num := stoi(strings.Split(value, "cm")[0])
			return 150 <= num && num <= 193
		} else if strings.Contains(value, "in") {
			num := stoi(strings.Split(value, "in")[0])
			return 59 <= num && num <= 76
		}
		return false
	case HairColor:
		hcl, _ := regexp.MatchString(`#[0-9,a-f]{6}`, value)
		return hcl && len(value) == 7
	case EyeColor:
		return contains(eyeColors, value)
	case PassportID:
		pid, _ := regexp.MatchString(`[0-9]{9}`, value)
		return pid && len(value) == 9
	default:
		return false
	}
}

func isValidOptionalPassportField(name string, value string) bool {
	switch name {
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
		field := strings.Split(fields[i], ":")
		name := field[0]
		value := field[1]
		if isValidRequiredPassportField(name, value) {
			numOfValidFields++
		} else {
			if isValidOptionalPassportField(name, value) {
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

	// TODO Fix bug that skips last text item
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			ph += text + " "
		} else {
			lines = append(lines, strings.TrimSpace(ph))
			ph = ""
		}
	}
	lines = append(lines, strings.TrimSpace(ph))

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
