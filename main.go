package main

import (
	"github.com/salasberryfin/torrent-client-go/client"
)

const fileName = "debian.iso.torrent"

func main() {
	client.Download(fileName)
}
