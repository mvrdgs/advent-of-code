package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	pwd, _ := os.Getwd()

	f, err := os.Open(pwd + "/day1/day1.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var sum int
	for scanner.Scan() {
		regex := regexp.MustCompile(`\d`)
		numbers := regex.FindAllString(scanner.Text(), -1)
		num, err := strconv.Atoi(fmt.Sprintf("%s%s", numbers[0], numbers[len(numbers)-1]))
		check(err)
		sum += num
	}

	fmt.Println(sum)
}
