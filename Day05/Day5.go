package main

import (
	"io/ioutil"
	"strconv"
	"strings"
	"fmt"
	"time"
	"sync"
)

func main() {
	start := time.Now()
	dat, err := ioutil.ReadFile("Day05/input.txt")
	check(err)
	datString := string(dat)
	var wg sync.WaitGroup
	var program []int
	var program2 []int
	wg.Add(2)
	go convertStringToIntArray(datString, &program, &wg)
	go convertStringToIntArray(datString, &program2, &wg)
	wg.Wait()
	wg.Add(2)
	go solvePartOne(program, &wg)
	go solvePartTwo(program2, &wg)
	wg.Wait()
	fmt.Printf("%v", time.Since(start))
}

func solvePartTwo(program []int, wg *sync.WaitGroup) {
	defer wg.Done()
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

func solvePartOne(program []int, wg *sync.WaitGroup) {
	defer wg.Done()
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

func convertStringToIntArray(input string, program *[]int, wg *sync.WaitGroup){
	defer wg.Done()
	for _, letter := range strings.Split(input, "\n") {
		number, _ := strconv.Atoi(letter)
		*program = append(*program, number)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
