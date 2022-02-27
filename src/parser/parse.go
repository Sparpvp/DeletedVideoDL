package parser

import (
	"fmt"
	"strings"

	"github.com/dustin/go-humanize"
)

type WriteCounter struct {
	Total         uint64
	ContentLength int64
}

type Video struct {
	Webarchive_Id string
	Video_Id      string
}

func (wc *WriteCounter) Init(ContentLength int64) {
	wc.ContentLength = ContentLength
}

func (wc *WriteCounter) Write(p []byte) (int, error) {

	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {

	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("\rDownloading... %s/%v complete", humanize.Bytes(wc.Total), humanize.Bytes(uint64(wc.ContentLength)))
}
