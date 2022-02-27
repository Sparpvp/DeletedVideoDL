package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/dustin/go-humanize"
)

type WriteCounter struct {
	Total uint64
}

type Video struct {
	wb_id string
	v_id  string
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

func DownloadVideo(hVideo Video) error {

	fmt.Println("Starting download...")
	out, err := os.Create(hVideo.v_id + ".mp4")
	if err != nil {
		log.Fatalln(err)
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Print("\n\nDetected CTRL^C: closing handles...\nGoodbye!\n")
		out.Close()
		os.Exit(1)
	}()

	res, err := http.Get(hVideo.wb_id)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	counter := &WriteCounter{}
	if _, err = io.Copy(out, io.TeeReader(res.Body, counter)); err != nil {
		out.Close()
		return err
	}

	fmt.Print("\n")
	out.Close()

	return nil
}

func main() {
	var youtube_url string
	fmt.Println("Insert (deleted or not) youtube video to download: ")
	fmt.Scan(&youtube_url)

	id := strings.Split(youtube_url, "?v=")
	video_id := id[len(id)-1]
	fmt.Println("Video Id:", video_id)
	webarchive_id := "https://web.archive.org/web/2oe_/http://wayback-fakeurl.archive.org/yt/" + video_id

	hVideo := Video{
		v_id:  video_id,
		wb_id: webarchive_id,
	}

	err := DownloadVideo(hVideo)
	if err != nil {
		log.Fatalln("Couldn't download video")
	}
}
