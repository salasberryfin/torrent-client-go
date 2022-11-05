package main

import (
	"fmt"
	"log"
	"os"

	"github.com/salasberryfin/torrent-client-go/client"
)

//const fileName = "archlinux.iso.torrent"
//const fileName = "ubuntu.iso.torrent"
const fileName = "debian.iso.torrent"

//const fileName = "test-ubuntu.torrent"

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get current working directory: ", err)
	}

	client, err := client.New(wd, fileName)
	if err != nil {
		log.Fatal("Failed to generate a new torrent client from file.")
	}
	fmt.Printf("The peer IP is %v\n", client.Peer.IP)
	fmt.Printf("The peer port is %v\n", client.Peer.Port)

	client.Connect()

	//tracker, err := torrent.NewHTTPTracker()
	//if err != nil {
	//	log.Fatal("Error generating HTTP tracker: ", err)
	//}
	//resp := tracker.Request(torrent.Data.Announce)
	//fmt.Printf("Expected to send request to tracker every %dms\n", resp.Interval)
	//peers := resp.GetPeers()
	//fmt.Println("Response from tracker:", peers[0])
	//network.NewConnection(tracker.InfoHash, tracker.PeerID)
	//network.Connect(peers[0].IP.String(), strconv.Itoa(peers[0].Port), msg)
	//r := network.RequestPayload{
	//	Index:  0,
	//	Begin:  0,
	//	Length: 0,
	//}
}
