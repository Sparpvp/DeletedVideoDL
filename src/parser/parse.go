package parser

import (
	"fmt"
	"strings"

	"github.com/dustin/go-humanize"
)

type WriteCounter struct {
	Total uint64
}

type Video struct {
	Webarchive_Id string
	Video_Id      string
}

func (wc *WriteCounter) Write(p []byte) (int, error) {

	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {

	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}
