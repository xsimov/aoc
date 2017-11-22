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

type room struct {
	name, checksum string
	sectorId       int
}

func (r *room) isReal() bool {
	letters := r.fiveMostCommonLettersInName()
	sort.Strings(letters)
	fmt.Println(strings.Join(letters, ""), r.checksum)
	return strings.Join(letters, "") == r.checksum
}

func (r *room) fiveMostCommonLettersInName() []string {
	s := r.name
	table := make(map[string]int)

	for _, c := range s {
		table[string(c)] += 1
	}

	var ss []struct {
		k string
		v int
	}
	for k, v := range table {
		if k == "-" {
			continue
		}
		ss = append(ss, struct {
			k string
			v int
		}{k: k, v: v})
	}
	sort.Slice(ss, func(i, j int) bool {
		if ss[i].v == ss[j].v {
			return ss[i].k < ss[j].k
		}
		return ss[i].v > ss[j].v
	})

	var orderedStr []string

	for _, dto := range ss[:5] {
		orderedStr = append(orderedStr, dto.k)
	}

	return orderedStr
}

func main() {
	lines, err := readFileByLines()
	if err != nil {
		log.Fatalf("could not import lines from file: %v", err)
	}
	for _, line := range lines {
		room := parseRoom(line)
		if room.isReal() {
			sectorSum += room.sectorId
		}
	}
	fmt.Println(sectorSum)
}

var encNameRE = regexp.MustCompile(`([a-z]+-)+`)
var sectorIdRE = regexp.MustCompile(`\d+`)
var checksumRE = regexp.MustCompile(`\[([a-z]+)\]`)

func parseRoom(s string) (r room) {
	r.name = encNameRE.FindAllString(s, -1)[0]

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
	tfile, err := ioutil.ReadFile("assets/day4.txt")
	if err != nil {
		return nil, fmt.Errorf("could not read file %v:", tfile)
	}
	lines := strings.Split(string(tfile), "\n")
	allButLast := lines[:(len(lines) - 1)]
	return allButLast, nil
}
