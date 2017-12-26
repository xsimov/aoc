package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type ip string

func (i *ip) supportsTLS() (possible bool) {
	var st int
	s := []byte(*i)
	for j := 0; j < len(s)-3; j++ {
		if s[j] == '[' {
			st += 1
		} else if s[j] == ']' {
			st -= 1
		}
		if isABBA(s[j : j+4]) {
			if st != 0 {
				return false
			} else {
				possible = true
			}
		}
	}
	if possible {
		fmt.Println("found one:", string(s))
	}
	return possible
}

func main() {
	var count int
	for _, ip := range getIPs() {
		if ip.supportsTLS() {
			count += 1
		}
	}

	fmt.Println(count)
}

func isABBA(s []byte) bool {
	return s[0] == s[3] && s[1] == s[2] && s[0] != s[1]
}

func getIPs() (ips []ip) {
	f, err := os.Open("../assets/day7.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("could not read contents: %v", err)
	}
	for _, ipStr := range strings.Split(string(content), "\n") {
		ips = append(ips, ip(ipStr))
	}
	return
}
