package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	dat, err := ioutil.ReadFile("Day06/input.txt")
	check(err)
	initialStateString := strings.Split(string(dat), "	")
	initialState := make([]int, len(initialStateString))
	maxIndex := 0
	max := math.MinInt32
	var persistString string
	for i, bank := range initialStateString {
		initialState[i], _ = strconv.Atoi(bank)
		if initialState[i] > max {
			max = initialState[i]
			maxIndex = i
		}
		persistString += " " + bank
	}
	stateMap := make(map[string]int)
	stateMap[persistString] = 0
	stepCount := 1
	redistributeCount := initialState[maxIndex]
findDuplicateState:
	for {
		persistString = ""
		redistributeCount = initialState[maxIndex]
		initialState[maxIndex] = 0
		startingIndex := maxIndex + 1
		max = 0
		maxIndex = 0
		if startingIndex == len(initialState) {
			startingIndex = 0
		}
		for i := startingIndex; redistributeCount != 0; {
			initialState[i]++
			i++
			redistributeCount--
			if i == len(initialState) {
				i = 0
			}
		}
		for i, value := range initialState {
			persistString += " " + strconv.Itoa(value)
			if initialState[i] > max || (initialState[i] == max && i < maxIndex) {
				max = initialState[i]
				maxIndex = i
			}
		}
		if _, present := stateMap[persistString]; present {
			stateMap[persistString] = stepCount - stateMap[persistString]
			break findDuplicateState
		}
		stateMap[persistString] = stepCount
		stepCount++
	}
	fmt.Printf("%v\n%v\n%v", stepCount, stateMap[persistString], time.Since(start))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
