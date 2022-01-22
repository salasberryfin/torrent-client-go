package main

import (
	"fmt"
	"os"

	"github.com/salasberryfin/torrent-client-go/client"
	"github.com/salasberryfin/torrent-client-go/utils"
)

//const fileName = "archlinux.iso.torrent"
//const fileName = "ubuntu.iso.torrent"
const fileName = "debian.iso.torrent"

//const fileName = "test-ubuntu.torrent"

func main() {
	wd, err := os.Getwd()
	utils.Check(err, "Failed to get current working directory")

	torrent, err := client.ParseTorrentFile(wd, fileName)
	fmt.Println("Torrent:", torrent.Tracker)
	//tracker, err := client.NewTracker(torrent, &wg)
	torrent.TrackerRequest()
	// trackerResponse, errTracker := torrentclient.TrackerRequest(torrent, &wg)
	// if errTracker != nil {
	// 	log.Fatal("Error when generating HTTP Tracker request: ", errTracker)
	// }
	// resp, errResp := torrentclient.ParseTrackerResponse(trackerResponse)
	// if errResp != nil {
	// 	log.Fatal("Failed to parse HTTP Tracker response: ", errResp)
	// }
	// log.Print("Tracker response: ", resp)
}
