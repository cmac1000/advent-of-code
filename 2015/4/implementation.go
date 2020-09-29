package main

import (
	"crypto/md5"
	"fmt"
	"hash"
	"io"
	"strconv"
	"strings"
)

// Verifier checks a hash against some criteria
type Verifier func(h hash.Hash) bool

func hashToString(h hash.Hash) string {
	return fmt.Sprintf("%x", h.Sum(nil))
}

func mine(key string, verifier Verifier) int {
	candidate := 0
	for {
		var combinedString strings.Builder
		combinedString.WriteString(key)
		combinedString.WriteString(strconv.Itoa(candidate))
		hash := md5.New()
		io.WriteString(hash, combinedString.String())
		if verifier(hash) {
			return candidate
		}
		candidate = candidate + 1
	}
}

func partOne(key string) int {
	verifier := func(h hash.Hash) bool {
		for _, char := range []rune(hashToString(h))[0:5] {
			if char != '0' {
				return false
			}
		}
		return true
	}
	return mine(key, verifier)
}

func partTwo(key string) int {
	verifier := func(h hash.Hash) bool {
		for _, char := range []rune(hashToString(h))[0:6] {
			if char != '0' {
				return false
			}
		}
		return true
	}
	return mine(key, verifier)
}

func main() {
	fmt.Println(partTwo("yzbqklnj"))
}
