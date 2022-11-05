package client

import (
	"github.com/salasberryfin/torrent-client-go/torrent"
)

// Client is an instance of the torrent client that is necessary to download any file
type Client struct {
	Torrent     torrent.Torrent
	Tracker     torrent.HTTPTracker
	AnnounceURL string
	Peer        torrent.Peer
}
