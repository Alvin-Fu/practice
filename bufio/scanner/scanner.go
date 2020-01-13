package scanner

import (
	"bufio"
	"log"
	"strings"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

var scanTests = []string{
	"",
	"a",
	"¼",
	"☹",
	"\x81",   // UTF-8 error
	"\uFFFD", // correctly encoded RuneError
	"abcdefgh",
	"abc def\n\t\tgh    ",
	"abc¼☹\x81\uFFFD日本語\x82abc",
}

func ScannerSplit() {
	read := strings.NewReader(scanTests[len(scanTests)-1])
	scanner := bufio.NewScanner(read)
	scanner.Split(bufio.ScanBytes)
}
