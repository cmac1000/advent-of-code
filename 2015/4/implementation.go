package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func partOne(key string) int {
	candidate := 0
	for {
		var combinedString strings.Builder
		combinedString.WriteString(key)
		combinedString.WriteString(strconv.Itoa(candidate))
		hash := md5.New()
		io.WriteString(hash, combinedString.String())
		if string([]rune(fmt.Sprintf("%x", hash.Sum(nil)))[0:5]) == "00000" {
			return candidate
		}
		candidate = candidate + 1

	}
}

func main() {
	fmt.Println(partOne("yzbqklnj"))
}
