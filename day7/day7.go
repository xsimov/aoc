package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type ip string

func (i *ip) supportsSSL() bool {
	var st int
	s := []byte(*i)
	for j := 0; j < len(s)-2; j++ {
		if s[j] == '[' {
			st += 1
		} else if s[j] == ']' {
			st -= 1
		}
		if st == 0 && isABA(s[j:j+3]) {
			fmt.Println("ABA: ", string(s[j:j+3]), st == 0)
			if (*i).scanForBAB(s[j], s[j+1]) {
				return true
			}
		}
	}
	return false
}

func (i *ip) scanForBAB(a, b byte) bool {
	var st int
	s := []byte(*i)
	for j := 0; j < len(s)-3; j++ {
		if s[j] == '[' {
			st += 1
		} else if s[j] == ']' {
			st -= 1
		}
		if st > 0 && isBAB(s[j:j+3], a, b) {
			fmt.Println("BAB: ", string(s[j:j+3]), st == 1)
			return true
		}
	}
	return false
}

func main() {
	var count int
	for _, ip := range getIPs() {
		if ip.supportsSSL() {
			count += 1
		}
	}

	fmt.Println(count)
}

func isABBA(s []byte) bool {
	return s[0] == s[3] && s[1] == s[2] && s[0] != s[1]
}

func isABA(s []byte) bool {
	if s[0] == '[' || s[0] == ']' || s[1] == '[' || s[1] == ']' || s[2] == '[' || s[2] == ']' {
		return false
	}
	return s[0] == s[2] && s[0] != s[1]
}

func isBAB(s []byte, a, b byte) bool {
	return s[0] == b && s[1] == a && s[2] == b
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
