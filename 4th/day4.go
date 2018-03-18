package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var sectorSum int

func main() {
	lines, err := readFileByLines()
	if err != nil {
		log.Fatalf("could not import lines from file: %v", err)
	}
	for _, line := range lines {
		room := parseRoom(line)
		if room.isReal() {
			room.decodeName()
			if matched, _ := regexp.MatchString("(?i)(north)|(pole)", room.decodedName); matched {
				fmt.Printf("Found the room: it is called %q and it is located in sector %v\n", room.decodedName, room.sectorId)
			}
		}
	}
}

var encNameRE = regexp.MustCompile(`([a-z]+-)+`)
var sectorIdRE = regexp.MustCompile(`\d+`)
var checksumRE = regexp.MustCompile(`\[([a-z]+)\]`)

func parseRoom(s string) (r Room) {
	r.encodedName = encNameRE.FindAllString(s, -1)[0]

	var cksum []string
	c := checksumRE.FindStringSubmatch(s)[1]
	for _, r := range c {
		cksum = append(cksum, string(r))
	}
	sort.Strings(cksum)
	r.checksum = strings.Join(cksum, "")

	r.sectorId, _ = strconv.Atoi(sectorIdRE.FindAllString(s, -1)[0])
	return r
}

func readFileByLines() ([]string, error) {
	tfile, err := ioutil.ReadFile("../assets/day4.txt")
	if err != nil {
		return nil, fmt.Errorf("could not read file %v:", tfile)
	}
	lines := strings.Split(string(tfile), "\n")
	return lines[:(len(lines) - 1)], nil
}
