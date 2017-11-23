package main

import (
	"sort"
	"strings"
)

type Room struct {
	decodedName, checksum, encodedName string
	sectorId                           int
}

func (r *Room) isReal() bool {
	letters := r.fiveMostCommonLettersInName()
	sort.Strings(letters)
	return strings.Join(letters, "") == r.checksum
}

func (r *Room) decodeName() {
	r.decodedName = string(decryptCaesar(r.encodedName, r.sectorId))
}

func (r *Room) fiveMostCommonLettersInName() []string {
	table := make(map[string]int)

	for _, c := range r.encodedName {
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

func decryptCaesar(s string, i int) []rune {
	c := shiftRune(s[0], i)
	if len(s) > 1 {
		return append([]rune{c}, decryptCaesar(s[1:], i)...)
	}
	return []rune{c}
}

func shiftRune(c byte, i int) rune {
	if c == '-' {
		return ' '
	}
	j := (int(c) - 97 + i) % 26
	if j < 0 {
		j += 26
	}
	j += 97
	return rune(j)
}
