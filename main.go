package main

import (
    "os"
    "log"

    "github.com/salasberryfin/go-torrent-client/torrentclient"
)

const FILE_NAME = "big-buck-bunny.torrent"

func main() {
    wd, err := os.Getwd()
    if err != nil {
        panic(err)
    }

    torrent, err := torrentclient.Parse(wd, FILE_NAME)
    httpTracker := torrentclient.GenerateTracker(torrent)
    log.Print(httpTracker)
    //torrent.Hash = infoHash
}


