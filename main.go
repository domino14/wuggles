package main

import (
	"flag"
	"github.com/domino14/macondo/gaddag"
)

func main() {
	//gaddagFile := flag.
	flag.Parse()
	gaddagFilename := flag.Args()[0]
	board := flag.Args()[1]
	sg := gaddag.LoadGaddag(gaddagFilename)

	wuggler(sg, board)
}

// LIASERTAIDKEMAIR 265 words
