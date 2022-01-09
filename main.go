package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/salasberryfin/torrent-client-go/client"
	"github.com/salasberryfin/torrent-client-go/utils"
)

//const fileName = "archlinux.iso.torrent"
const fileName = "ubuntu.iso.torrent"

func main() {
	// wait for goroutines to finish
	var wg sync.WaitGroup
	wg.Add(1)
	wd, err := os.Getwd()
	utils.Check(err, "Failed to get current working directory")

	torrent, err := client.ParseTorrentFile(wd, fileName)
	tracker, err := client.NewTracker(torrent, &wg)
	fmt.Println("Using port ", tracker.Port)
	// trackerResponse, errTracker := torrentclient.TrackerRequest(torrent, &wg)
	// if errTracker != nil {
	// 	log.Fatal("Error when generating HTTP Tracker request: ", errTracker)
	// }
	// resp, errResp := torrentclient.ParseTrackerResponse(trackerResponse)
	// if errResp != nil {
	// 	log.Fatal("Failed to parse HTTP Tracker response: ", errResp)
	// }
	// log.Print("Tracker response: ", resp)

	wg.Wait()
}
