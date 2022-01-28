package main

import (
	"fmt"
	"log"
	"os"

	"github.com/salasberryfin/torrent-client-go/client"
	"github.com/salasberryfin/torrent-client-go/network"
	"github.com/salasberryfin/torrent-client-go/utils"
)

//const fileName = "archlinux.iso.torrent"
//const fileName = "ubuntu.iso.torrent"
//const fileName = "debian.iso.torrent"

const fileName = "test-ubuntu.torrent"

func main() {
	wd, err := os.Getwd()
	utils.Check(err, "Failed to get current working directory")

	torrent, err := client.New(wd, fileName)
	tracker, err := torrent.NewHTTPTracker()
	if err != nil {
		log.Fatal("Error generating HTTP tracker: ", err)
	}
	resp := tracker.Request(torrent.Data.Announce)
	peers := resp.GetPeers()
	fmt.Println("Response from tracker:", peers[0])
	network.InitHandshake(tracker.InfoHash, tracker.PeerID)
}
