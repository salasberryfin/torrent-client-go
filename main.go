package main

import (
    "os"
    "log"

    "github.com/salasberryfin/go-torrent-client/torrentclient"
)

const FILE_NAME = "cosmos-laundromat.torrent"
//const FILE_NAME = "ubuntu.torrent"
//const FILE_NAME = "big-buck-bunny.torrent"

func main() {
    wd, err := os.Getwd()
    if err != nil {
        panic(err)
    }

    torrent, err := torrentclient.Parse(wd, FILE_NAME)
    log.Print("Announce: ", torrent.Data.AnnounceList)
    httpTracker, errTracker := torrentclient.GenerateTracker(torrent)
    if errTracker != nil {
        log.Print("Error when generating HTTP Tracker: ", errTracker)
    }
    log.Print(httpTracker)
    //torrent.Hash = infoHash
}


