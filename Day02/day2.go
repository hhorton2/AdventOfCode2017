package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("Day02/input.txt")
	check(err)
	datString := string(dat)
	rows := strings.Split(datString, "\n")
	checkSum := getChecksum(rows)
	divisibleCheckSum := getDivisibleChecksum(rows)
	fmt.Println(checkSum)
	fmt.Println(divisibleCheckSum)
}
func getChecksum(rows []string) int {
	checksum := 0
	for _, row := range rows {
		min := math.MaxInt32
		max := math.MinInt32
		for _, column := range strings.Split(row, "	") {
			currentNum, _ := strconv.Atoi(column)
			if currentNum < min {
				min = currentNum
			} else if currentNum > max {
				max = currentNum
			}
		}
		checksum += max - min
	}
	return checksum
}

func getDivisibleChecksum(rows []string) int {
	checksum := 0
	for _, row := range rows {
		for i, column := range strings.Split(row, "	") {
			currentNum, _ := strconv.Atoi(column)
			for j, compareColumn := range strings.Split(row, "	") {
				if i != j {
					compareNum, _ := strconv.Atoi(compareColumn)
					if currentNum%compareNum == 0 {
						checksum += currentNum / compareNum
					}
				}
			}
		}
	}
	return checksum
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
