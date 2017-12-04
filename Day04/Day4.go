package main

import (
	"io/ioutil"
	"strings"
	"fmt"
	"sync"
	"sort"
)

func main() {
	dat, err := ioutil.ReadFile("Day04/input.txt")
	check(err)
	datString := string(dat)
	passphraseStream := make(chan string, 10)
	go getPassphrases(datString, passphraseStream)
	passphraseStreamA := make(chan string, 10)
	passphraseStreamB := make(chan string, 10)
	go duplicateStream(passphraseStream, passphraseStreamA, passphraseStreamB)
	resultStreamPartOne := make(chan bool, 10)
	resultStreamPartTwo := make(chan bool, 10)
	go verifyPassphrasePartOne(passphraseStreamA, resultStreamPartOne)
	go verifyPassphrasePartTwo(passphraseStreamB, resultStreamPartTwo)
	var wg sync.WaitGroup
	wg.Add(2)
	go countValidPassphrases("one", resultStreamPartOne, &wg)
	go countValidPassphrases("two", resultStreamPartTwo, &wg)
	wg.Wait()
}

func countValidPassphrases(day string, resultStream <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	for result := range resultStream {
		if result {
			count++
		}
	}
	fmt.Printf("Day %v: %v\n", day, count)
}
func verifyPassphrasePartOne(passphraseStream <-chan string, resultStream chan<- bool) {
	defer close(resultStream)
	for passphrase := range passphraseStream {
		phraseMap := make(map[string]bool)
		isValid := true
	verifyWord:
		for _, word := range strings.Split(passphrase, " ") {
			if _, present := phraseMap[word]; !present {
				phraseMap[word] = true
			} else {
				isValid = false
				break verifyWord
			}
		}
		resultStream <- isValid
	}
}

func verifyPassphrasePartTwo(passphraseStream <-chan string, resultStream chan<- bool) {
	defer close(resultStream)
	for passphrase := range passphraseStream {
		phraseMap := make(map[string]bool)
		isValid := true
	verifyWord:
		for _, word := range strings.Split(passphrase, " ") {
			s := strings.Split(word, "")
			sort.Strings(s)
			word = strings.Join(s, "")
			if _, present := phraseMap[word]; !present {
				phraseMap[word] = true
			} else {
				isValid = false
				break verifyWord
			}
		}
		resultStream <- isValid
	}
}

func getPassphrases(input string, passphraseStream chan<- string) {
	defer close(passphraseStream)
	for _, passphrase := range strings.Split(input, "\n") {
		passphraseStream <- passphrase
	}
}

func duplicateStream(passPhrases <-chan string, forkA chan<- string, forkB chan<- string) {
	defer close(forkA)
	defer close(forkB)
	for position := range passPhrases {
		forkA <- position
		forkB <- position
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
