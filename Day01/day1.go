package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	dat, err := ioutil.ReadFile("Day01/input.txt")
	check(err)
	datString := string(dat)
	numbers := getNumbers(datString)
	partOneSum := getPartOneSum(numbers)
	partTwoSum := getPartTwoSum(numbers)
	fmt.Println(partOneSum)
	fmt.Println(partTwoSum)
}
func getPartOneSum(numbers []int) int {
	sum := 0
	currentCheck := 0
	for _, number := range numbers {
		if number == currentCheck {
			sum += number
		} else {
			currentCheck = number
		}
	}
	if currentCheck == numbers[0] {
		sum += numbers[0]
	}
	return sum
}

func getPartTwoSum(numbers []int) int {
	sum := 0
	halfLength := len(numbers) / 2
	for i := 0; i < halfLength; i++ {
		if numbers[i] == numbers[i+halfLength] {
			sum += numbers[i] * 2
		}
	}
	return sum
}

func getNumbers(input string) []int {
	var numbers []int
	for i := 1; i < len(input); i++ {
		number, _ := strconv.Atoi(input[i-1 : i])
		numbers = append(numbers, number)
	}
	return numbers
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
