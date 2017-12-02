package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	dat, err := ioutil.ReadFile("Day02/input.txt")
	check(err)
	datString := string(dat)
	fmt.Println(datString)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
