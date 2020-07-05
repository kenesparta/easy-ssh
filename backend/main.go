package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func check(e error) {
	if nil != e {
		log.Println(e)
	}
}

func readFile(fileName string) {
	dat, err := ioutil.ReadFile(fileName)
	check(err)
	fmt.Print(string(dat))
}

func main() {
	readFile("config")
}
