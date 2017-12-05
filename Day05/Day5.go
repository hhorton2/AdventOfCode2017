package main

import (
	"io/ioutil"
	"strconv"
	"strings"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	dat, err := ioutil.ReadFile("Day05/input.txt")
	check(err)
	datString := string(dat)
	program := convertStringToIntArray(datString)
	solvePartOne(program)
	program = convertStringToIntArray(datString)
	solvePartTwo(program)
	fmt.Printf("%v", time.Since(start))
}

func solvePartTwo(program []int) {
	position := 0
	count := 0
	for position < len(program) {
		count++
		newPosition := position + program[position]
		if program[position] >= 3 {
			program[position]--
		} else {
			program[position]++
		}
		position = newPosition
	}
	fmt.Println(count)
}

func solvePartOne(program []int) {
	position := 0
	count := 0
	for position < len(program) {
		count++
		newPosition := position + program[position]
		program[position]++
		position = newPosition
	}
	fmt.Println(count)
}

func convertStringToIntArray(input string) []int {
	var program []int
	for _, letter := range strings.Split(input, "\n") {
		number, _ := strconv.Atoi(letter)
		program = append(program, number)
	}
	return program
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
