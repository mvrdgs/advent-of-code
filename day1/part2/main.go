package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

var (
	spelledNumbers = map[string]string{
		"one":       "1",
		"two":       "2",
		"three":     "3",
		"four":      "4",
		"five":      "5",
		"six":       "6",
		"seven":     "7",
		"eight":     "8",
		"nine":      "9",
		"ten":       "10",
		"eleven":    "11",
		"twelve":    "12",
		"thirteen":  "13",
		"fourteen":  "14",
		"fifteen":   "15",
		"sixteen":   "16",
		"seventeen": "17",
		"eighteen":  "18",
		"nineteen":  "19",
		"twenty":    "20",
	}
	spelledKeys = reflect.ValueOf(spelledNumbers).MapKeys()
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getNumbers(s string) []string {
	var result []string
	var buffer string

	for _, v := range s {
		if unicode.IsDigit(v) {
			result = append(result, string(v))
			buffer = ""
		} else {
			buffer += string(v)
			for v := range spelledKeys {
				v := spelledKeys[v].String()
				if strings.Contains(buffer, v) {
					result = append(result, spelledNumbers[v])
					buffer = buffer[len(buffer)-1:]
				}
			}
		}
	}

	return result
}

func main() {
	pwd, _ := os.Getwd()

	f, err := os.Open(pwd + "/day1/day1.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var sum int
	for scanner.Scan() {
		text := scanner.Text()
		numbers := getNumbers(text)
		reflect.TypeOf(numbers)
		num, err := strconv.Atoi(fmt.Sprintf("%s%s", numbers[0], numbers[len(numbers)-1]))
		check(err)
		sum += num
	}

	fmt.Println(sum)
}
