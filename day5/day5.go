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

	// ch := make(chan rune)

	for i := 0; len(pass) < 8; i++ {
		hashInput(i)
	}
	// for len(pass) < 8 {
	// 	select {
	// 	case r := <-ch:
	// 		pass = append(pass, r)
	// 	}
	// }

	fmt.Println(string(pass))
}

func hashInput(i int) {
	h := md5.New()
	io.WriteString(h, input+strconv.Itoa(i))
	src := h.Sum(nil)
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	if strings.HasPrefix(string(dst), "00000") {
		pass = append(pass, rune(dst[5]))
		fmt.Println("found one: ", dst, string(pass), i)
	}
}
