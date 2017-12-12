package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type frequencyTable map[rune]int

func (f frequencyTable) mostFrequentLetter() rune {
	var (
		currentMax  int
		finalLetter rune
	)
	for letter, times := range f {
		if times > currentMax {
			currentMax = times
			finalLetter = letter
		}
	}
	return finalLetter
}

var freqTables = make([]frequencyTable, 8)

func main() {
	receptions := getReceptions()
	for _, reception := range receptions {
		fmt.Println(reception)
		for pos, letter := range reception {
			if freqTables[pos] == nil {
				freqTables[pos] = make(frequencyTable)
			}
			freqTables[pos][rune(letter)] += 1
		}
	}

	message := make([]rune, 8)

	for i, freqTable := range freqTables {
		message[i] = freqTable.mostFrequentLetter()
	}
	fmt.Println(string(message))
}

func getReceptions() []string {
	f, err := os.Open("../assets/day6.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("could not read contents: %v", err)
	}
	return strings.Split(string(content), "\n")
}
