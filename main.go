package main

import (
	"flag"
	"os"
	"path/filepath"
	"strconv"

	"github.com/domino14/macondo/gaddag"
)

var LexiconPath = os.Getenv("LEXICON_PATH")

// Note: the Q is treated as QU (there is no Q tile in Wuggles, but it is
// represented as Q in this CLI).
func main() {
	flag.Parse()
	lexiconName := flag.Args()[0]
	board := flag.Args()[1]
	round := flag.Args()[2]
	dawg, err := gaddag.LoadDawg(filepath.Join(LexiconPath, "dawg", lexiconName+".dawg"))
	if err != nil {
		panic(err)
	}
	rdNum, err := strconv.Atoi(round)
	if rdNum < 1 || rdNum > 4 || err != nil {
		panic("round number must be between 1 and 4 inclusive")
	}
	runeBoard, err := convertBoard(board, rdNum)
	if err != nil {
		panic(err)
	}
	wuggler(dawg, runeBoard, rdNum)
}

// LIASERTAIDKEMAIR 274 words using TWL14
