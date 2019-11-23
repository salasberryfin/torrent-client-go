package main

import (
    "os"
    "log"
    "sync"

    "github.com/salasberryfin/go-torrent-client/torrentclient"
)

const FILE_NAME = "Ubuntu_519_archive.torrent"

func main() {
    // wait for goroutines to finish
    var wg sync.WaitGroup
    wg.Add(1)
    wd, err := os.Getwd()
    if err != nil {
        panic(err)
    }

    torrent, err := torrentclient.Parse(wd, FILE_NAME)
    log.Print("Announce: ", torrent.Data.AnnounceList)
    trackerResponse, errTracker := torrentclient.TrackerRequest(torrent, &wg)
    if errTracker != nil {
        log.Fatal("Error when generating HTTP Tracker request: ", errTracker)
    }
    log.Print(trackerResponse)

    //wg.Wait()
}


