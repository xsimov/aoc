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

var pass []rune

func main() {
	start := time.Now()
	defer func() { fmt.Printf("It took: %v\n", time.Since(start)) }()

	done := make(chan bool)
	res := make(chan rune)

	var i int
	for i = 0; i < 10; i++ {
		go hashInput(i, res, done)
	}

	for len(pass) < 8 {
		select {
		case <-done:
		case r := <-res:
			fmt.Println("found one:", r)
			pass = append(pass, r)
		}

		i++
		go hashInput(i, res, done)
	}

	fmt.Println(string(pass))
}

func hashInput(i int, res chan rune, done chan bool) {
	h := md5.New()
	io.WriteString(h, input+strconv.Itoa(i))
	src := h.Sum(nil)
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	if strings.HasPrefix(string(dst), "00000") {
		res <- rune(dst[5])
		return
	}
	done <- true
}
