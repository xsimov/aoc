package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

const input string = "ffykfhsq"

var (
	itemsFound      int
	pass            = make([]rune, 8)
	positionsFilled = make([]bool, 8)
)

func main() {
	start := time.Now()
	defer func() { fmt.Printf("It took: %v\n", time.Since(start)) }()

	for i := 0; itemsFound < 8; i++ {
		hashInput(i)
	}
	fmt.Println(string(pass))
}

func hashInput(i int) {
	h := md5.New()
	io.WriteString(h, input+strconv.Itoa(i))
	src := h.Sum(nil)
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	if strings.HasPrefix(string(dst), "00000") {
		if position, valid := isANumberBetween0And7(dst[5]); valid {
			fmt.Println("position is: ", position)
			if !positionsFilled[position] {
				pass[position] = rune(dst[6])
				positionsFilled[position] = true
				itemsFound++
				fmt.Println("found one: ", dst, string(pass), i)
			}
		}
	}
}

func isANumberBetween0And7(b byte) (int, bool) {
	num, err := strconv.Atoi(string(rune(b)))
	if err == nil && num >= 0 && num < 8 {
		return num, true
	}
	return 0, false
}
