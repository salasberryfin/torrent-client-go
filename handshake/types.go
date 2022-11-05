package handshake

import "github.com/salasberryfin/torrent-client-go/torrent"

// Handshake contains the information used for the initial TCP connection
type Handshake struct {
	Pstr     string
	Reserved []byte
	InfoHash []byte
	PeerID   []byte
	Peer     torrent.Peer
}
