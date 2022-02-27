package downloader

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Sparpvp/DeletedVideoDL/src/parser"
)

func DownloadVideo(hVideo parser.Video) error {

	fmt.Println("Starting download...")
	out, err := os.Create(hVideo.Video_Id + ".mp4")
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

	counter := &parser.WriteCounter{}
	head, _ := http.Head(hVideo.Webarchive_Id)
	counter.Init(head.ContentLength)

	res, err := http.Get(hVideo.Webarchive_Id)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	if _, err = io.Copy(out, io.TeeReader(res.Body, counter)); err != nil {
		out.Close()
		return err
	}

	fmt.Print("\n")
	out.Close()

	return nil
}
