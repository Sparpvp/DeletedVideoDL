package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/Sparpvp/DeletedVideoDL/src/downloader"
	"github.com/Sparpvp/DeletedVideoDL/src/parser"
)

func main() {
	var youtube_url string
	fmt.Println("Insert (deleted or not) youtube video to download: ")
	fmt.Scan(&youtube_url)

	id := strings.Split(youtube_url, "?v=")
	video_id := id[len(id)-1]
	fmt.Println("Video Id:", video_id)
	webarchive_id := "https://web.archive.org/web/2oe_/http://wayback-fakeurl.archive.org/yt/" + video_id

	hVideo := parser.Video{
		Webarchive_Id: webarchive_id,
		Video_Id:      video_id,
	}

	err := downloader.DownloadVideo(hVideo)
	if err != nil {
		log.Fatalln("Couldn't download video")
	}
}
