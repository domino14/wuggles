package main

import (
	"path/filepath"
	"testing"

	"github.com/domino14/macondo/gaddag"
)

func BenchmarkFind(b *testing.B) {
	tst := "LIASERTAIDKEMAIR"
	dawg, err := gaddag.LoadDawg(
		filepath.Join(LexiconPath, "dawg", "America.dawg"))
	if err != nil {
		b.Error(err)
	}
	runeBoard, err := convertBoard(tst, 1)
	if err != nil {
		b.Error(err)
	}
	for i := 0; i < b.N; i++ {
		wuggler(dawg, runeBoard, 1)
	}
}
