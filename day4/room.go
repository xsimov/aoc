package main

import (
	"fmt"
	"sort"
	"strings"
)

type Room struct {
	name, checksum string
	sectorId       int
}

func (r *Room) isReal() bool {
	letters := r.fiveMostCommonLettersInName()
	sort.Strings(letters)
	fmt.Println(strings.Join(letters, ""), r.checksum)
	return strings.Join(letters, "") == r.checksum
}

func (r *Room) fiveMostCommonLettersInName() []string {
	table := make(map[string]int)

	for _, c := range r.name {
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
